package main

import (
	"basicbots/delay"
	"basicbots/eval"
	"basicbots/tokenizer"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

func ResetRobots() error {
	for i := 0; i < numberOfRobots; i++ {
		Robots[i].Damage = 0
		Robots[i].Status = ALIVE
		Robots[i].Heading = 0.0 // float64(rand.Intn(359)) // Should be 0-359
		Robots[i].HeadingWanted = Robots[i].Heading
		Robots[i].Speed = 0.0
		Robots[i].SpeedWanted = Robots[i].Speed
		Robots[i].SpeedHold = 0.0
		Robots[i].Distance = 0.0
		Robots[i].Reload = 0
		Robots[i].Scan = 0
		Robots[i].Width = 2
		Robots[i].Cannon = 0

		// Place the robots on the battle field.
		switch i {
		case 0: // Upper Left
			//Robots[i].X = 1.0
			//Robots[i].Y = 1.0
			Robots[i].X = float64(rand.Intn(100)) + 100.0
			Robots[i].Y = float64(rand.Intn(100)) + 100.0
		case 1: // Lower Right
			//Robots[i].X = 500.0
			//Robots[i].Y = 500.0
			Robots[i].X = float64(rand.Intn(100)) + 800.0
			Robots[i].Y = float64(rand.Intn(100)) + 800.0
		case 2: // Upper Right
			//Robots[i].X = 800.0
			//Robots[i].Y = 200.0
			Robots[i].X = float64(rand.Intn(100)) + 800.0
			Robots[i].Y = float64(rand.Intn(100)) + 100.0
		case 3: // Lower Left
			//Robots[i].X = 200.0
			//Robots[i].Y = 800.0
			Robots[i].X = float64(rand.Intn(100)) + 100.0
			Robots[i].Y = float64(rand.Intn(100)) + 800.0
		}

		Robots[i].XOrigin = Robots[i].X
		Robots[i].YOrigin = Robots[i].Y

		for m := 0; m < MAXMISSILES; m++ {
			Missiles[i][m].Status = AVAILABLE
			Missiles[i][m].Reload = 0
		}

	}
	return nil
}

func InitRobots() error {
	var err error

	err = ResetRobots()
	if err != nil {
		return err
	}

	if len(evaluator) != 0 {
		evaluator = evaluator[:0]
	}
	if len(token) != 0 {
		token = token[:0]
	}

	for i := 0; i < numberOfRobots; i++ {

		/*
			robotFileName := filepath.Base(robotFileNameWithPath)
			robotFileNameNoExt := robotFileName[:len(robotFileName)-len(filepath.Ext(robotFileName))]
			robotOutput := "logs/" + robotFileNameNoExt + ".out"
			robotDebug1 := "logs/" + robotFileNameNoExt + ".d1"
			robotDebug2 := "logs/" + robotFileNameNoExt + ".d2"
		*/

		robotFileNameWithPath := flag.Args()[i]
		if len(Robots[i].Program) == 0 {
			Robots[i].Program, err = ioutil.ReadFile(robotFileNameWithPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not load '%s'\n", flag.Args()[i])
				os.Exit(1)
			}
			Robots[i].Name = filepath.Base(flag.Args()[i])
		}

		// Tokenize
		tt := tokenizer.New(string(Robots[i].Program))
		token = append(token, tt)

		ee, err := eval.New(token[i])
		if err != nil {
			return err
		}

		/*
			if debug {
				fmt.Printf("Robot:%d:%s InitRobot X:%5.2f Y:%5.2f\n", i, Robots[i].Name, Robots[i].X, Robots[i].Y)
			}
		*/
		if trace {
			if i == 1 {
				ee.SetTrace(true)
			}
		}
		ee.RegisterBuiltin("LOCX", 0, FunctionLocX)
		ee.RegisterBuiltin("LOCY", 0, FunctionLocY)
		ee.RegisterBuiltin("SPEED", 0, FunctionSpeed)
		ee.RegisterBuiltin("DAMAGE", 0, FunctionDamage)
		ee.RegisterBuiltin("DRIVE", 2, FunctionDrive)
		ee.RegisterBuiltin("SCAN", 2, FunctionScan)
		ee.RegisterBuiltin("CANNON", 2, FunctionCannon)
		evaluator = append(evaluator, ee)

	}

	return nil
}

func RunRobots() error {

	var alive int

	if battledisplay {
		scr.Show()
	}

	if battledisplay {
		go eventloop()
	}

	cycles = 0

	for {

		if battledisplay {
			if cycles%30 == 0 {
				plotbattlefield()
				scr.Show()
			}
		}

		select {
		case etype = <-event:
			// fmt.Println("received message", msg)
		default:
		}

		if etype == 99 {
			break
		}

		for current = 0; current < numberOfRobots; current++ {

			if Robots[current].Damage >= 100 {
				Robots[current].Status = DEAD
				// fmt.Fprintf(os.Stderr, "Robot:%d DEAD s:%d  d:%d\n", current, Robots[current].Status, Robots[current].Damage)
				//alive--
				//Robots[current].MU.Unlock()
				continue
			}

			if Robots[current].Reload > 0 {
				Robots[current].Reload--
			}

			err := evaluator[current].RunStep()
			if err != nil {
				//evaluator[current].ProgramEnd = true
				//fmt.Fprintf(os.Stderr, "Robot:%d DEAD evaluator[current].RunStep() err \n", current)
				//alive--
				etype = 98
				em := fmt.Sprintf("Error running program:\n\t%s\n", err.Error())
				return errors.New(em)
			}

			if evaluator[current].ProgramEnd {
				Robots[current].Damage = 100
				Robots[current].Status = DEAD
				// fmt.Fprintf(os.Stderr, "Robot:%d DEAD evaluator[current].ProgramEnd\n", current)
				//alive-- // if program ends remove it from the cound
				continue // dont end battlebots if a single program ends.
			}

		}

		alive = 0
		for nn := 0; nn < numberOfRobots; nn++ {
			if Robots[nn].Status == ALIVE {
				alive++
			}
		}

		// if there is only 1 or 0 robots left alive - break out of the loop
		if alive == 0 || alive == 1 {
			break
		}

		cycles++

		// end of simulation ?
		if cycles == maxCycles {
			if debug {
				fmt.Fprintln(os.Stderr, "Cycle limit has been reached.")
			}
			break
		}

		// Time to update movement
		if cycles%MOVECLICKS == 0 {
			moverobot()
			movemissile()
			//fmt.Fprintf(os.Stderr, "c: %d\n", cycles)
		}

		// Pause for needed time to slow down the battledisplay
		if battledisplay {
			delay.Delay(cycledelay)
		}
		//fmt.Fprintf(os.Stderr, "cycles: %d\n", cycles)
	}
	//fmt.Fprintf(os.Stderr, "c: %d\n", cycles)
	//fmt.Printf("Alive: %d\n", alive)
	//for nn := 0; nn < numberOfRobots; nn++ {
	//	fmt.Printf("%d ", Robots[nn].Status)
	//}
	//fmt.Println()

	//	fmt.Fprintf(os.Stderr, "out of cpu loop %d - alive=%d\n", cycles, alive)
	if alive == 0 {
		for nn := 0; nn < numberOfRobots; nn++ {
			Robots[nn].Tie++
			// fmt.Fprintf(os.Stderr, "nn:%d %d\n", nn, Robots[nn].Status)
		}
	}

	if alive != 0 {
		if alive == 1 {
			for nn := 0; nn < numberOfRobots; nn++ {
				if Robots[nn].Status == ALIVE {
					Robots[nn].Winner++
				} else {
					Robots[nn].Lose++
				}
			}
		}

		if alive > 1 {
			for nn := 0; nn < numberOfRobots; nn++ {
				if Robots[nn].Status == ALIVE {
					Robots[nn].Tie++
				}
				if Robots[nn].Status == DEAD {
					Robots[nn].Lose++
				}
			}
		}
	}

	return nil
}

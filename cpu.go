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

// ResetRobots : Reset all robots to default values. Used to reset robots between matches.
func ResetRobots() error {

	// Cheesy way to scamble a array
	// Used for randomizing the starting location of robots.
	pp := make([]int, 4)
	for i := 0; i < 4; i++ {
		pp[i] = i
	}
	// 25 rounds of swaps
	for i := 0; i < 25; i++ {
		var s1, s2 int
		s1 = rand.Intn(4)
		for {
			s2 = rand.Intn(4)
			if s2 != s1 {
				break
			}
		}
		swap1 := pp[s1]
		swap2 := pp[s2]
		pp[s2] = swap1
		pp[s1] = swap2
	}

	// Reset / Init all robots.
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
		switch pp[i] {
		case 0: // Upper Left
			Robots[i].X = float64(rand.Intn(100)) + 100.0
			Robots[i].Y = float64(rand.Intn(100)) + 100.0
		case 1: // Lower Right
			Robots[i].X = float64(rand.Intn(100)) + 800.0
			Robots[i].Y = float64(rand.Intn(100)) + 800.0
		case 2: // Upper Right
			Robots[i].X = float64(rand.Intn(100)) + 800.0
			Robots[i].Y = float64(rand.Intn(100)) + 100.0
		case 3: // Lower Left
			Robots[i].X = float64(rand.Intn(100)) + 100.0
			Robots[i].Y = float64(rand.Intn(100)) + 800.0
		}

		Robots[i].XOrigin = Robots[i].X
		Robots[i].YOrigin = Robots[i].Y

		// Set all missiles to default.
		for m := 0; m < MAXMISSILES; m++ {
			Missiles[i][m].Status = AVAILABLE
			Missiles[i][m].Reload = 0
		}

	}
	return nil
}

// InitRobots : Initialize the robots. Used to load the program, reset eval and token.
func InitRobots() error {
	var err error

	err = ResetRobots()
	if err != nil {
		return err
	}

	// Clear the previous slice if any
	if len(evaluator) != 0 {
		evaluator = evaluator[:0]
	}

	// Clear the previous slice if any
	if len(token) != 0 {
		token = token[:0]
	}

	// Loop and load each robots source and initialize eval,token and customer basic functions.
	for i := 0; i < numberOfRobots; i++ {

		/*
			robotFileName := filepath.Base(robotFileNameWithPath)
			robotFileNameNoExt := robotFileName[:len(robotFileName)-len(filepath.Ext(robotFileName))]
			robotOutput := "logs/" + robotFileNameNoExt + ".out"
			robotDebug1 := "logs/" + robotFileNameNoExt + ".d1"
			robotDebug2 := "logs/" + robotFileNameNoExt + ".d2"
		*/
		var robotFileNameWithPath string
		if bench {
			robotFileNameWithPath = "testbots/bench2.bas"
		} else {
			robotFileNameWithPath = flag.Args()[i]
		}
		if len(Robots[i].Program) == 0 {
			if bench {
				Robots[i].Program = []byte(benchbot)
			} else {
				Robots[i].Program, err = ioutil.ReadFile(robotFileNameWithPath)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not load '%s'\n", flag.Args()[i])
				os.Exit(1)
			}
			//Robots[i].Name = filepath.Base(flag.Args()[i])
			Robots[i].Name = filepath.Base(robotFileNameWithPath)

		}

		// Create tokens for the robots source. Tokenize
		tt := tokenizer.New(string(Robots[i].Program))
		token = append(token, tt)

		// Create new eval for the robot.
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

		// Add the customer funcitons of basicbots to the eval
		ee.RegisterBuiltin("LOCX", 0, FunctionLocX)
		ee.RegisterBuiltin("LOCY", 0, FunctionLocY)
		ee.RegisterBuiltin("SPEED", 0, FunctionSpeed)
		ee.RegisterBuiltin("DAMAGE", 0, FunctionDamage)
		ee.RegisterBuiltin("DRIVE", 2, FunctionDrive)
		ee.RegisterBuiltin("SCAN", 2, FunctionScan)
		ee.RegisterBuiltin("CANNON", 2, FunctionCannon)
		ee.RegisterBuiltin("IN", 0, FunctionIn)
		ee.RegisterBuiltin("OUT", 1, FunctionOut)
		ee.RegisterBuiltin("STRC$", 2, FunctionSTRC) // STRC <number>, <count>
		ee.RegisterBuiltin("TEAM", 0, FunctionTeam)
		evaluator = append(evaluator, ee)

	}

	return nil
}

// RunRobots : Main loop for executing the code for the robots, triggers movement.
func RunRobots() error {

	// var alive int

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
			checkAlive(current)
			if Robots[current].Status == DEAD {
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
				etype = BASICERROR
				em := fmt.Sprintf("Error running program:\n\t%s\n", err.Error())
				return errors.New(em)
			}

			if evaluator[current].ProgramEnd {
				Robots[current].Damage = 100
				Robots[current].Status = DEAD
				continue // dont end battlebots if a single program ends.
			}

		}

		if endCondition() {
			break
		}

		cycles++

		// end of simulation ?
		if cycles == maxCycles {
			for nn := 0; nn < numberOfRobots; nn++ {
				if Robots[nn].Status == ALIVE {
					Robots[nn].Tie++
					Robots[nn].Points++
				}
			}
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
	}

	//totalPoints = numberOfRobots

	endGame()

	//fmt.Println(teams, a)

	//	fmt.Fprintf(os.Stderr, "out of cpu loop %d - alive=%d\n", cycles, alive)
	//TeamsWinner()

	//alive = countAlive()

	/*
		if alive == 0 {
			for nn := 0; nn < numberOfRobots; nn++ {
				Robots[nn].Tie++
				// fmt.Fprintf(os.Stderr, "nn:%d %d\n", nn, Robots[nn].Status)
			}
		}
	*/
	/*
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
					if Robots[nn].Status == ALIVE {1
						Robots[nn].Tie++
					}
					if Robots[nn].Status == DEAD {
						Robots[nn].Lose++
					}
				}
			}
		}
	*/
	return nil
}

/*
func countAlive() int {
	a := 0
	for nn := 0; nn < numberOfRobots; nn++ {
		if Robots[nn].Status == ALIVE {
			a++
		}
	}
	return a
}
*/

func robotStatusToint() int {
	var a int

	if Robots[0].Status == ALIVE {
		a = a | 1
	}
	if Robots[1].Status == ALIVE {
		a = a | 2
	}
	if Robots[2].Status == ALIVE {
		a = a | 4
	}
	if Robots[3].Status == ALIVE {
		a = a | 8
	}
	return a
}

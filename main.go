package main

import (
	"flag"
	"fmt"

	"os"
	"strings"
	"time"
)

// main : Really?
func main() {

	var err error

	var versionflag bool

	// Parse flags to expose the arguments
	flag.IntVar(&maxCycles, "c", 1500000, "Maximum number of cycles. Default is 1500000.")
	flag.BoolVar(&debug, "d", false, "Enable debug.")
	flag.BoolVar(&trace, "trace", false, "Enale program trace.")
	flag.BoolVar(&battledisplay, "b", false, "Enale battle display.")
	flag.IntVar(&matchcount, "m", 1, "Number of matches to simulate.")
	flag.BoolVar(&versionflag, "v", false, "Display version and credits.")
	flag.BoolVar(&teams, "t", false, "Enable teams.")
	flag.BoolVar(&bench, "bench", false, "Do benchmarking.")
	//flag.BoolVar(&timingTest, "time", false, "Turn on timing tests.")
	flag.Parse()

	if !bench {
		// if -v is set and there are no robots on the command line, call version() to output the version data.
		// I dont know why I am checking for robots. I really shouldnt.
		if versionflag || len(flag.Args()) == 0 {
			version()
			os.Exit(0)
		}
	}
	cycledelay = int64(time.Microsecond * 100)

	// Seed the random number generator. Doesn't need to be crypto strong.
	//rand.Seed(time.Now().UnixNano())

	Robots = make([]Robot, MAXROBOTS)
	numberOfRobots = len(flag.Args())
	if numberOfRobots > MAXROBOTS { // Check number of robots on the command line
		fmt.Fprintln(os.Stderr, "To many robots. Max:", MAXROBOTS)
		os.Exit(3)
	}

	/*
		if timingTest {
			battledisplay = true
			matchcount = 1
		}
	*/

	if bench {
		numberOfRobots = 4
		battledisplay = false
		matchcount = 500
		teams = false
	}

	// if teams is set, number of robots must be four
	if numberOfRobots != 4 && teams {
		fmt.Fprintf(os.Stderr, "Teams flag is set and the number of robots does not equal four.\n")
		os.Exit(3)
	}

	// If the battledisplay flag is set, make sure the matchcount = 1. Overriding -m
	if battledisplay {
		matchcount = 1
		initDisplay() // Create and Initialize the tcell module.
	}

	startTime := time.Now() // benching
	for match := 0; match < matchcount; match++ {
		err = InitRobots()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		exiterror = RunRobots()
		if etype != 0 {
			break
		}
	}

	// benching based on SGNIPS from https://crobots.deepthought.it/home.php?page=sgnips
	endDuration := time.Since(startTime).Seconds()
	if bench {
		ww := (1500000 * 500) / endDuration
		fmt.Printf("Bench: %d\n", int(ww))
		os.Exit(0)
	}

	// if battledisplay flaf is set finialize tcell
	if battledisplay {
		// if the exit type is not a exit key then delay 10 seconds. Time to read the error or the score before restoring the screem
		if etype != ESCKEY {
			time.Sleep(10 * time.Second)
		}
		scr.Fini()
	}

	if exiterror != nil {
		fmt.Fprintf(os.Stderr, "%v\n", exiterror) // Some kind of error, most likely error in basic program
		os.Exit(1)                                // Exit without showing winner.
	}

	// Output the w,t,l,p for all robots
	for i := 0; i < numberOfRobots; i++ {
		space := strings.Repeat(" ", 20-len(Robots[i].Name))
		//points := (Robots[i].Winner * 3) + Robots[i].Tie
		fmt.Printf("%s%s w:%05d t:%05d l:%05d p:%05d\n", Robots[i].Name, space, Robots[i].Winner, Robots[i].Tie, Robots[i].Lose, Robots[i].Points)
	}
	if teams {
		t := "Team1"
		space := strings.Repeat(" ", 20-len(t))
		fmt.Printf("%s%s w:%05d t:%05d l:%05d p:%05d\n", t, space,
			Robots[0].Winner|Robots[1].Winner,
			Robots[0].Tie|Robots[1].Tie,
			Robots[0].Lose|Robots[1].Lose,
			Robots[0].Points|Robots[1].Points)

		t = "Team2"
		space = strings.Repeat(" ", 20-len(t))
		fmt.Printf("%s%s w:%05d t:%05d l:%05d p:%05d\n", t, space,
			Robots[2].Winner|Robots[3].Winner,
			Robots[2].Tie|Robots[3].Tie,
			Robots[2].Lose|Robots[3].Lose,
			Robots[2].Points|Robots[3].Points)
	}

}

// version : Print to stdout a version message along with copyright and other required messages.
func version() {
	fmt.Printf("basicbots %s\n\n", VERSION)
	fmt.Printf("basicbots is created and copyrighted by William Jones\nand is licensed under GNU GPL v2\n\n")
	fmt.Printf("Modules used are:\n")
	fmt.Printf("gobasic by Steve Kemp and used by permission.\n")
	fmt.Printf("\tgobasic is licensed unther GPL v2 and can be \n\tfound at https://github.com/skx/gobasic\n\n")
	fmt.Printf("tcell by Garrett D'Amore and is licensend under Apache License 2.0\n\n")
	os.Exit(0)
}

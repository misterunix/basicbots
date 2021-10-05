package main

import (
	"flag"
	"fmt"
	"math/rand"
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

	flag.Parse()

	// if -v is set and there are no robots on the command line, call version() to output the version data.
	// I dont know why I am checking for robots. I really shouldnt.
	if versionflag || len(flag.Args()) == 0 {
		version()
		os.Exit(0)
	}

	cycledelay = int64(time.Microsecond * 100)

	// Seed the random number generator. Doesn't need to be crypto strong.
	rand.Seed(time.Now().UnixNano())

	Robots = make([]Robot, MAXROBOTS)
	numberOfRobots = len(flag.Args())
	if numberOfRobots > MAXROBOTS { // Check number of robots on the command line
		fmt.Fprintln(os.Stderr, "To many robots. Max:", MAXROBOTS)
		os.Exit(3)
	}

	// if teams is set, number of robots must be four
	if numberOfRobots != 4 && teams {
		fmt.Fprintf(os.Stderr, "Teams flag is set and the number of robots does not equal four.\n")
		os.Exit(3)
	}

	// If the battledisplay flag is set, make sure the matchcount = 1. OVerriding -m
	if battledisplay {
		matchcount = 1
		initDisplay() // Create and Initialize the tcell module.
	}

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

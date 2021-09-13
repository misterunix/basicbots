package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	var err error

	var versionflag bool

	// Parse flags to expose the arguments
	flag.IntVar(&maxCycles, "c", 1500000, "Maximum number of cycles. Default is 1500000.")
	flag.BoolVar(&debug, "d", false, "Enable debug.")
	flag.BoolVar(&trace, "t", false, "Enale program trace.")
	flag.BoolVar(&battledisplay, "b", false, "Enale battle display.")
	flag.IntVar(&matchcount, "m", 1, "Number of matches to simulate.")
	flag.BoolVar(&versionflag, "v", false, "Display version and credits.")

	flag.Parse()

	if versionflag || len(flag.Args()) == 0 {
		version()
		os.Exit(0)
	}

	cycledelay = int64(time.Microsecond * 100)

	// Seed the random number generator. Doesn't need to be crypto strong.
	rand.Seed(time.Now().UnixNano())

	Robots = make([]Robot, MAXROBOTS)
	numberOfRobots = len(flag.Args())
	if numberOfRobots > MAXROBOTS {
		fmt.Fprintln(os.Stderr, "To many robots. Max:", MAXROBOTS)
		os.Exit(3)
	}

	if battledisplay {
		matchcount = 1
		initDisplay()
	}

	for match := 0; match < matchcount; match++ {

		//fmt.Fprintf(os.Stderr, "match start: %d\n", match)

		err = InitRobots()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		exiterror = RunRobots()
		if etype != 0 {
			break
		}

		//fmt.Fprintf(os.Stderr, "match end: %d\n", match)

	}

	if battledisplay {
		if etype != 99 {
			time.Sleep(10 * time.Second)
		}
		scr.Fini()
	}

	if exiterror != nil {
		fmt.Fprintf(os.Stderr, "%v\n", exiterror)
		os.Exit(1) // Exit without showing winner.
	}

	for i := 0; i < numberOfRobots; i++ {
		space := strings.Repeat(" ", 20-len(Robots[i].Name))
		points := (Robots[i].Winner * 3) + Robots[i].Tie
		fmt.Printf("%s%s w:%05d t:%05d l:%05d p:%05d\n", Robots[i].Name, space, Robots[i].Winner, Robots[i].Tie, Robots[i].Lose, points)
		//fmt.Printf("%+v\n", Robots[i])
	}

}

func version() {
	fmt.Printf("basicbots %s\n\n", VERSION)
	fmt.Printf("basicbots is created and copyrighted by William Jones\nand is licensed under GNU GPL v2\n\n")
	fmt.Printf("Modules used are:\n")
	fmt.Printf("gobasic by Steve Kemp and used by permission.\n")
	fmt.Printf("\tgobasic is licensed unther GPL v2 and can be \n\tfound at https://github.com/skx/gobasic\n\n")
	fmt.Printf("tcell by Garrett D'Amore and is licensend under Apache License 2.0\n\n")
	os.Exit(0)
}

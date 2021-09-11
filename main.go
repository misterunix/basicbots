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

	// Parse flags to expose the arguments
	flag.IntVar(&maxCycles, "c", 1500000, "Maximum number of cycles. Default is 1500000.")
	flag.BoolVar(&debug, "d", false, "Enable debug.")
	flag.BoolVar(&trace, "t", false, "Enale program trace.")
	flag.BoolVar(&battledisplay, "b", false, "Enale battle display.")
	flag.IntVar(&matchcount, "m", 1, "Number of matches to simulate.")

	flag.Parse()

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

		err = RunRobots()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		//fmt.Fprintf(os.Stderr, "match end: %d\n", match)

	}

	if battledisplay {
		scr.Fini()
	}

	for i := 0; i < numberOfRobots; i++ {
		space := strings.Repeat(" ", 20-len(Robots[i].Name))
		fmt.Printf("%s%s w:%05d t:%05d l:%05d\n", Robots[i].Name, space, Robots[i].Winner, Robots[i].Tie, Robots[i].Lose)
		//fmt.Printf("%+v\n", Robots[i])
	}

}

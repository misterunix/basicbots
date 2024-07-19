package main

import (
	"log"
	"os"
	"strings"
)

type Challenge struct {
	RobotsName []string `json:"robotname"` // List of robots
	Wins       []int    `json:"wins"`      // List of wins
	Ties       []int    `json:"ties"`      // List of ties
	Losses     []int    `json:"losses"`    // List of losses
}

type Robot struct {
	Filename string    `json:"filename"` // Filename of the robot
	Count    int       `json:"count"`    // Number of times this robot has competed
	Points   int       `json:"points"`   // Points scored by this robot
	Win      int       `json:"win"`      // Number of wins
	Tie      int       `json:"tie"`      // Number of ties
	Loss     int       `json:"loss"`     // Number of losses
	Battles  Challenge `json:"battles"`  // List of battles this robot has competed agaist other robots
}

func main() {

	robots := make([]Robot, 0)

	file, err := os.ReadDir("../robots/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range file {
		if f.IsDir() {
			continue
		}
		fn := f.Name()
		fn = strings.ToLower(fn)

		if strings.HasSuffix(fn, ".bas") {
			yt := Robot{}

			robots = append(robots, yt)
		}
	}

	/*
		nbots := len(robots)
		if nbots <= 1 {
			fmt.Println("Not enough robots to run a tournament")
			return
		}

		// Tournament 2x2
		for i := 0; i < nbots-1; i++ {
			for j := i + 1; j < nbots; j++ {
				fmt.Printf("Tournament: %s vs %s\n", robots[i].Filename, robots[j].Filename)
				robots[i].Count++
				robots[j].Count++
			}
		}
		for _, r := range robots {
			fmt.Printf("%s: %d\n", r.Filename, r.Count)
		}
		for ii := range robots {
			robots[ii].Count = 0
		}

		// Tournament 3x3
		if nbots > 2 {
			for i := 0; i < nbots-2; i++ {
				for j := i + 1; j < nbots-1; j++ {
					for k := j + 1; k < nbots; k++ {
						fmt.Printf("Tournament: %s vs %s vs %s\n", robots[i].Filename, robots[j].Filename, robots[k].Filename)
						robots[i].Count++
						robots[j].Count++
						robots[k].Count++
					}
				}
			}
			for _, r := range robots {
				fmt.Printf("%s: %d\n", r.Filename, r.Count)
			}
		}
		for ii := range robots {
			robots[ii].Count = 0
		}

		// Tournament 4x4
		if nbots > 3 {
			for i := 0; i < nbots-3; i++ {
				for j := i + 1; j < nbots-2; j++ {
					for k := j + 1; k < nbots-1; k++ {
						for l := k + 1; l < nbots; l++ {
							fmt.Printf("Tournament: %s vs %s vs %s vs %s\n", robots[i].Filename, robots[j].Filename, robots[k].Filename, robots[l].Filename)
							robots[i].Count++
							robots[j].Count++
							robots[k].Count++
							robots[l].Count++
						}
					}
				}
			}
			for _, r := range robots {
				fmt.Printf("%s: %d\n", r.Filename, r.Count)
			}
		}
		for ii := range robots {
			robots[ii].Count = 0
		}
	*/

}

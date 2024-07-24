package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"database/sql"

	"github.com/misterunix/sniffle/hashing"
	_ "modernc.org/sqlite"
)

type therobots struct {
	ID           int     // ID
	OwnerID      int     // Owner ID
	Filename     string  // Filename of the robot
	FilenameHash string  // Filename hash
	Code         string  // Code of the robot
	CodeHash     string  // Code hash
	Count        int     // Number of times this robot has competed
	Points       float64 // Points scored by this robot
	Win          int     // Number of wins
	Tie          int     // Number of ties
	Loss         int     // Number of losses
}

// type Challenge struct {
// 	RobotsName []string `json:"robotname"` // List of robots
// 	Wins       []int    `json:"wins"`      // List of wins
// 	Ties       []int    `json:"ties"`      // List of ties
// 	Losses     []int    `json:"losses"`    // List of losses
// }

// type Robot struct {
// 	Filename string    `json:"filename"` // Filename of the robot
// 	Count    int       `json:"count"`    // Number of times this robot has competed
// 	Points   int       `json:"points"`   // Points scored by this robot
// 	Win      int       `json:"win"`      // Number of wins
// 	Tie      int       `json:"tie"`      // Number of ties
// 	Loss     int       `json:"loss"`     // Number of losses
// 	Battles  Challenge `json:"battles"`  // List of battles this robot has competed agaist other robots
// }

var botsTableCreate string
var botsTabeleDelete string
var botsSelect string

func main() {

	botsTabeleDelete = "DROP TABLE IF EXISTS robots;"
	botsTableCreate = "CREATE TABLE IF NOT EXISTS robots (ID INTEGER PRIMARY KEY, filename TEXT, fnhash TEXT, code TEXT, codehash TEXT, count INTEGER, points REAL, win INTEGER, tie INTEGER, loss INTEGER);"
	botsSelect = "SELECT ID,filename,fnhash,code,codehash,count,points,win,tie,loss FROM robots WHERE ID="

	//robots := make(map[int]string)

	db, err := sql.Open("sqlite", "robots.db")
	if err != nil {
		log.Fatal(err)
	}

	robotStorage := make([]therobots, 0)

	file, err := os.ReadDir("../robots/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range file {
		if f.IsDir() {
			continue
		}
		robotFilename := f.Name()
		x := strings.ToLower(robotFilename)

		if !strings.HasSuffix(x, ".bas") {
			continue
		}

		robotFilenameHash := hashing.StringHash(hashing.SHA256, robotFilename)
		fmt.Println("Want", robotFilename)

		tmp := db.ReadStr(robotFilenameHash)
		if len(tmp) == 0 {
			// need to add it
			tr := therobots{}
			tr.Filename = robotFilename
			tr.FilenameHash = robotFilenameHash

			rc, err := os.ReadFile("../robots/" + robotFilename)
			if err != nil {
				log.Fatal(err)
			}

			tr.Code = string(rc)
			tr.CodeHash = hashing.StringHash(hashing.SHA256, tr.Code)
			tr.Count = 0
			tr.Points = 0
			tr.Win = 0
			tr.Tie = 0
			tr.Loss = 0
			robotStorage = append(robotStorage, tr)
			trj, err := json.Marshal(tr)
			if err != nil {
				fmt.Println("Error on Marshal")
				log.Fatal(err)
			}

			fmt.Println("adding", fn)
			db.Map[fnh] = string(trj)
			continue
		}

		tr := therobots{}
		//fmt.Println(len(tmp), tmp)

		err := json.Unmarshal([]byte(tmp), &tr)
		if err != nil {
			fmt.Println("Error on Unmarshal")
			log.Fatal(err)
		}
		robotStorage = append(robotStorage, tr)

		// fmt.Println(tr.Filename)
		// fmt.Println(tr.FilenameHash)
		// fmt.Println(tr.CodeHash)
		// fmt.Println()
		// fmt.Println()
	}
	nbots := len(robotStorage)

	// Tournament 2x2
	for i := 0; i < nbots-1; i++ {
		for j := i + 1; j < nbots; j++ {
			buf := new(bytes.Buffer)
			fmt.Printf("Tournament: %s vs %s\n", robotStorage[i].Filename, robotStorage[j].Filename)
			matches := "111"
			cmd := exec.Command("../bin/basicbots-linux_amd64", "-m", matches, "../robots/"+robotStorage[i].Filename, "../robots/"+robotStorage[j].Filename)
			cmd.Stdout = buf
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error on Run")
				log.Fatal(err)
			}
			lines := strings.Split(buf.String(), "\n")
			for _, line := range lines {
				parts1 := strings.Split(line, "\t")
				parts2 := strings.Split(parts1[1], " ")
				parts3 := strings.Split(parts2[0], ":")
				whoami := hashing.StringHash(hashing.SHA256, parts1[0])
				tt := db.ReadStr(whoami)
				ttu := therobots{}
				err := json.Unmarshal([]byte(tt), &ttu)
				if err != nil {
					fmt.Println("Error on Unmarshal")
					log.Fatal(err)
				}
				ttu.Count++
				w, _ := strconv.Atoi(parts3[1])
				t, _ := strconv.Atoi(parts3[3])
				l, _ := strconv.Atoi(parts3[5])
				p, _ := strconv.ParseFloat(parts3[7], 64)
				ttu.Win += w
				ttu.Tie += t
				ttu.Loss += l
				ttu.Points += p

			}

			fmt.Println(buf.String())
		}
	}

}

// err = db.Close()
// if err != nil {
// 	log.Fatal(err)
// }

// if strings.HasSuffix(fnl, ".bas") {
// 	fileHash, err := hashing.FileHash(hashing.SHA256, fn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	robotsHash[fileHash] = fn
// 	robots[index] = fileHash
// }

// nbots := len(robots)
// if nbots <= 1 {
// 	fmt.Println("Not enough robots to run a tournament")
// 	return
// }

// // Tournament 2x2
// for i := 0; i < nbots-1; i++ {
// 	for j := i + 1; j < nbots; j++ {
// 		ri0 := robots[i]
// 		ri1 := robots[j]
// 		r0 := robotsHash[ri0]
// 		r1 := robotsHash[ri1]
// 		fmt.Printf("Tournament: %s vs %s\n", r0, r1)
// 		//robots[i].Count++
// 		//robots[j].Count++
// 	}
// }
// for _, r := range robots {
// 	fmt.Printf("%s: %d\n", r.Filename, r.Count)
// }
// for ii := range robots {
// 	robots[ii].Count = 0
// }

// // Tournament 3x3
// if nbots > 2 {
// 	for i := 0; i < nbots-2; i++ {
// 		for j := i + 1; j < nbots-1; j++ {
// 			for k := j + 1; k < nbots; k++ {
// 				fmt.Printf("Tournament: %s vs %s vs %s\n", robots[i].Filename, robots[j].Filename, robots[k].Filename)
// 				robots[i].Count++
// 				robots[j].Count++
// 				robots[k].Count++
// 			}
// 		}
// 	}
// 	for _, r := range robots {
// 		fmt.Printf("%s: %d\n", r.Filename, r.Count)
// 	}
// }
// for ii := range robots {
// 	robots[ii].Count = 0
// }

// // Tournament 4x4
// if nbots > 3 {
// 	for i := 0; i < nbots-3; i++ {
// 		for j := i + 1; j < nbots-2; j++ {
// 			for k := j + 1; k < nbots-1; k++ {
// 				for l := k + 1; l < nbots; l++ {
// 					fmt.Printf("Tournament: %s vs %s vs %s vs %s\n", robots[i].Filename, robots[j].Filename, robots[k].Filename, robots[l].Filename)
// 					robots[i].Count++
// 					robots[j].Count++
// 					robots[k].Count++
// 					robots[l].Count++
// 				}
// 			}
// 		}
// 	}
// 	for _, r := range robots {
// 		fmt.Printf("%s: %d\n", r.Filename, r.Count)
// 	}
// }
// for ii := range robots {
// 	robots[ii].Count = 0
// }

package main

import (
	"bytes"

	"fmt"
	"log"
	"os"
	"os/exec"

	"strings"

	"github.com/misterunix/sniffle/hashing"
	sqlhelper "github.com/misterunix/sqlite-helper"
)

func main() {

	//robots := make(map[int]string)

	db = sqlhelper.New()
	db.Path = "."
	db.Filename = "robots.db"

	err := db.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create the table if it doesn't exist
	p := CheckTables()
	if !p {
		err = CreateDB()
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.ReadDir("../robots/")
	if err != nil {
		log.Fatal(err)
	}

	// check to see if the robot is in the database
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

		tempID := -1
		sqlstring := "SELECT ID FROM robots WHERE filenamehash = '" + robotFilenameHash + "';"
		err := db.Db.QueryRow(sqlstring).Scan(&tempID)
		if err != nil {
			log.Fatal(err)
		}

		if tempID != -1 {
			continue // in the DB
		}

		// need to add it
		tr := therobots{}
		tr.Filename = robotFilename
		tr.FilenameHash = robotFilenameHash
		tr.OwnerID = 0
		tr.Count = 0
		tr.Win = 0
		tr.Tie = 0
		tr.Loss = 0
		tr.Points = 0
		rc, err := os.ReadFile("../robots/" + robotFilename)
		if err != nil {
			log.Fatal(err)
		}
		tr.Code = string(rc)
		tr.CodeHash = hashing.StringHash(hashing.SHA256, tr.Code)
		InsertIntoTable("robots", tr)

		// tmp := db.ReadStr(robotFilenameHash)
		// if len(tmp) == 0 {
		// 	// need to add it
		// 	tr := therobots{}
		// 	tr.Filename = robotFilename
		// 	tr.FilenameHash = robotFilenameHash

		// 	rc, err := os.ReadFile("../robots/" + robotFilename)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	tr.Code = string(rc)
		// 	tr.CodeHash = hashing.StringHash(hashing.SHA256, tr.Code)
		// 	tr.Count = 0
		// 	tr.Points = 0
		// 	tr.Win = 0
		// 	tr.Tie = 0
		// 	tr.Loss = 0
		// 	robotStorage = append(robotStorage, tr)
		// 	trj, err := json.Marshal(tr)
		// 	if err != nil {
		// 		fmt.Println("Error on Marshal")
		// 		log.Fatal(err)
		// 	}

		// 	fmt.Println("adding", fn)
		// 	db.Map[fnh] = string(trj)
		// 	continue

		// tr := therobots{}
		// //fmt.Println(len(tmp), tmp)

		// err := json.Unmarshal([]byte(tmp), &tr)
		// if err != nil {
		// 	fmt.Println("Error on Unmarshal")
		// 	log.Fatal(err)
		// }
		// robotStorage = append(robotStorage, tr)

		// fmt.Println(tr.Filename)
		// fmt.Println(tr.FilenameHash)
		// fmt.Println(tr.CodeHash)
		// fmt.Println()
		// fmt.Println()
	}

	// robots table should be complete
	nbots := 0
	sqlstring := "select count(*) as e from robots;"
	err = db.Db.QueryRow(sqlstring).Scan(&nbots)
	if err != nil {
		log.Fatal(err)
	}

	robotStorage := make([]therobots, 0)

	sqlstring = "select ID,Filename from robots;"
	rows, err := db.Db.Query(sqlstring)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var fn therobots
		err = rows.Scan(&fn.ID, &fn.Filename)
		if err != nil {
			log.Fatal(err)
		}
		robotStorage = append(robotStorage, fn)
	}
	rows.Close()

	//nbots := len(robotStorage)

	// Tournament 2x2
	for i := 0; i < nbots-1; i++ {
		for j := i + 1; j < nbots; j++ {
			buf := new(bytes.Buffer)
			fmt.Printf("Tournament: %s vs %s\n", robotStorage[i].Filename, robotStorage[j].Filename)
			matches := "17"
			cmd := exec.Command("../bin/basicbots-linux_amd64", "-tt", "-m", matches, "../robots/"+robotStorage[i].Filename+"../robots/"+robotStorage[j].Filename)
			cmd.Stdout = buf
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error on Run")
				log.Fatal(err)
			}
			lines := strings.Split(buf.String(), "\n")
			for _, line := range lines {
				parts1 := strings.Split(line, " ")
				filename := parts1[0]
				if len(filename) == 0 {
					log.Fatalln("Error on filename, during select")
				}
				w := parts1[1]
				t := parts1[2]
				l := parts1[3]
				p := parts1[4]

				sqlstring = "UPDATE SET	win = win + " + w + ", tie = tie + " + t + ", loss = loss + " + l + ", points = points + " + p + " WHERE filename = '" + filename + "';"
				_, err := db.Db.Exec(sqlstring)
				if err != nil {
					log.Fatal(err)
				}

				//parts2 := strings.Split(parts1[1], " ")
				//parts3 := strings.Split(parts2[0], ":")
				//whoami := hashing.StringHash(hashing.SHA256, parts1[0])

				//tt := db.Db.ReadStr(whoami)
				// ttu := therobots{}
				// err := json.Unmarshal([]byte(tt), &ttu)
				// if err != nil {
				// 	fmt.Println("Error on Unmarshal")
				// 	log.Fatal(err)
				// }
				// ttu.Count++
				// w, _ := strconv.Atoi(parts3[1])
				// t, _ := strconv.Atoi(parts3[3])
				// l, _ := strconv.Atoi(parts3[5])
				// p, _ := strconv.ParseFloat(parts3[7], 64)
				// ttu.Win += w
				// ttu.Tie += t
				// ttu.Loss += l
				// ttu.Points += p

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

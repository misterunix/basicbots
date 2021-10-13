package main

func endGame() {
	a := robotStatusToint()

	if teams {
		if a == 0 { // no survivors - no points
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 0
			Robots[0].Points += 0

			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 0
			Robots[1].Points += 0

			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 0
			Robots[2].Points += 0

			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 0
			Robots[3].Points += 0
		}
		if a == 1 { // Robots[0] is the winner, Robot[2] gets win as well.
			Robots[0].Winner += 1
			Robots[0].Tie += 0
			Robots[0].Lose += 0
			Robots[0].Points += 4

			Robots[1].Winner += 1
			Robots[1].Tie += 0
			Robots[1].Lose += 0
			Robots[1].Points += 4

			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0

			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 2 { // Robots[1] is the winner. Robots[0] gets win as well.
			Robots[0].Winner += 1
			Robots[0].Tie += 0
			Robots[0].Lose += 0
			Robots[0].Points += 4

			Robots[1].Winner += 1
			Robots[1].Tie += 0
			Robots[1].Lose += 0
			Robots[1].Points += 4

			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0

			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 3 { // Robots[0] and Robots[1] are the winners
			Robots[0].Winner += 1
			Robots[0].Tie += 0
			Robots[0].Lose += 0
			Robots[0].Points += 4

			Robots[1].Winner += 1
			Robots[1].Tie += 0
			Robots[1].Lose += 0
			Robots[1].Points += 4

			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0

			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 4 { // Robots[2] is the winner and Robots[3] get carried along as well.
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0

			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0

			Robots[2].Winner += 1
			Robots[2].Tie += 0
			Robots[2].Lose += 0
			Robots[2].Points += 4

			Robots[3].Winner += 1
			Robots[3].Tie += 0
			Robots[3].Lose += 0
			Robots[3].Points += 4
		}
		if a == 8 { // Robots[3] is the winner and Robots[2] get carried along as well.
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0

			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0

			Robots[2].Winner += 1
			Robots[2].Tie += 0
			Robots[2].Lose += 0
			Robots[2].Points += 4

			Robots[3].Winner += 1
			Robots[3].Tie += 0
			Robots[3].Lose += 0
			Robots[3].Points += 4
		}
		if a == 12 { // Robots[2] and Robots[3] are the Winners
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0

			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0

			Robots[2].Winner += 1
			Robots[2].Tie += 0
			Robots[2].Lose += 0
			Robots[2].Points += 4

			Robots[3].Winner += 1
			Robots[3].Tie += 0
			Robots[3].Lose += 0
			Robots[3].Points += 4
		}
		if a == 15 { // All robots are still alive, must be max cycles. Tie for everyone
			Robots[0].Winner += 0
			Robots[0].Tie += 1
			Robots[0].Lose += 0
			Robots[0].Points += 1

			Robots[1].Winner += 0
			Robots[1].Tie += 1
			Robots[1].Lose += 0
			Robots[1].Points += 1

			Robots[2].Winner += 0
			Robots[2].Tie += 1
			Robots[2].Lose += 0
			Robots[2].Points += 1

			Robots[3].Winner += 0
			Robots[3].Tie += 1
			Robots[3].Lose += 0
			Robots[3].Points += 1
		}
	} else {
		if a == 0 { // Every Robot is dead
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0

			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0

			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0

			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 1 { // Robots[0] is the winner
			Robots[0].Winner += 1
			Robots[0].Tie += 0
			Robots[0].Lose += 0
			Robots[0].Points += 4

			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0

			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0

			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 2 { // Robots[1] is the winner
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0
			Robots[1].Winner += 1
			Robots[1].Tie += 0
			Robots[1].Lose += 0
			Robots[1].Points += 4
			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0
			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 3 { // 1 & 0 winner
			Robots[0].Winner += 0
			Robots[0].Tie += 1
			Robots[0].Lose += 0
			Robots[0].Points += 1
			Robots[1].Winner += 0
			Robots[1].Tie += 1
			Robots[1].Lose += 0
			Robots[1].Points += 1
			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0
			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 4 { // Robots[2] is the winner
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0
			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0
			Robots[2].Winner += 1
			Robots[2].Tie += 0
			Robots[2].Lose += 0
			Robots[2].Points += 4
			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 5 { // 0 & 2 winner
			Robots[0].Winner += 0
			Robots[0].Tie += 1
			Robots[0].Lose += 0
			Robots[0].Points += 1
			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0
			Robots[2].Winner += 0
			Robots[2].Tie += 1
			Robots[2].Lose += 0
			Robots[2].Points += 1
			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 6 { // 1 & 2 winner
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0
			Robots[1].Winner += 0
			Robots[1].Tie += 1
			Robots[1].Lose += 0
			Robots[1].Points += 1
			Robots[2].Winner += 0
			Robots[2].Tie += 1
			Robots[2].Lose += 0
			Robots[2].Points += 1
			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 7 { // 0 & 1 & 2 winner
			Robots[0].Winner += 0
			Robots[0].Tie += 1
			Robots[0].Lose += 0
			Robots[0].Points += 1
			Robots[1].Winner += 0
			Robots[1].Tie += 1
			Robots[1].Lose += 0
			Robots[1].Points += 1
			Robots[2].Winner += 0
			Robots[2].Tie += 1
			Robots[2].Lose += 0
			Robots[2].Points += 1
			Robots[3].Winner += 0
			Robots[3].Tie += 0
			Robots[3].Lose += 1
			Robots[3].Points += 0
		}
		if a == 8 { // Robots[3] is the winner
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0
			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0
			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0
			Robots[3].Winner += 1
			Robots[3].Tie += 0
			Robots[3].Lose += 0
			Robots[3].Points += 4
		}
		if a == 9 { // 0 & 3 winner
			Robots[0].Winner += 0
			Robots[0].Tie += 1
			Robots[0].Lose += 0
			Robots[0].Points += 1
			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0
			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0
			Robots[3].Winner += 0
			Robots[3].Tie += 1
			Robots[3].Lose += 0
			Robots[3].Points += 1
		}
		if a == 10 { // 1,3
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0
			Robots[1].Winner += 0
			Robots[1].Tie += 1
			Robots[1].Lose += 0
			Robots[1].Points += 1
			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0
			Robots[3].Winner += 0
			Robots[3].Tie += 1
			Robots[3].Lose += 0
			Robots[3].Points += 1
		}
		if a == 11 { // 0,1,3
			Robots[0].Winner += 0
			Robots[0].Tie += 1
			Robots[0].Lose += 0
			Robots[0].Points += 1
			Robots[1].Winner += 0
			Robots[1].Tie += 1
			Robots[1].Lose += 0
			Robots[1].Points += 1
			Robots[2].Winner += 0
			Robots[2].Tie += 0
			Robots[2].Lose += 1
			Robots[2].Points += 0
			Robots[3].Winner += 0
			Robots[3].Tie += 1
			Robots[3].Lose += 0
			Robots[3].Points += 1
		}
		if a == 12 { // 2,3
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0
			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 0
			Robots[2].Winner += 0
			Robots[2].Tie += 1
			Robots[2].Lose += 0
			Robots[2].Points += 1
			Robots[3].Winner += 0
			Robots[3].Tie += 1
			Robots[3].Lose += 0
			Robots[3].Points += 1
		}
		if a == 13 { // 0,2,3
			Robots[0].Winner += 0
			Robots[0].Tie += 1
			Robots[0].Lose += 0
			Robots[0].Points += 1

			Robots[1].Winner += 0
			Robots[1].Tie += 0
			Robots[1].Lose += 1
			Robots[1].Points += 1

			Robots[2].Winner += 0
			Robots[2].Tie += 1
			Robots[2].Lose += 0
			Robots[2].Points += 0

			Robots[3].Winner += 0
			Robots[3].Tie += 1
			Robots[3].Lose += 0
			Robots[3].Points += 1
		}
		if a == 14 { // 1,2,3
			Robots[0].Winner += 0
			Robots[0].Tie += 0
			Robots[0].Lose += 1
			Robots[0].Points += 0
			Robots[1].Winner += 0
			Robots[1].Tie += 1
			Robots[1].Lose += 0
			Robots[1].Points += 1
			Robots[2].Winner += 0
			Robots[2].Tie += 1
			Robots[2].Lose += 0
			Robots[2].Points += 1
			Robots[3].Winner += 0
			Robots[3].Tie += 1
			Robots[3].Lose += 0
			Robots[3].Points += 1
		}
		if a == 15 { // All robots still alive - tie for all.
			Robots[0].Winner += 0
			Robots[0].Tie += 1
			Robots[0].Lose += 0
			Robots[0].Points += 1
			Robots[1].Winner += 0
			Robots[1].Tie += 1
			Robots[1].Lose += 0
			Robots[1].Points += 1
			Robots[2].Winner += 0
			Robots[2].Tie += 1
			Robots[2].Lose += 0
			Robots[2].Points += 1
			Robots[3].Winner += 0
			Robots[3].Tie += 1
			Robots[3].Lose += 0
			Robots[3].Points += 1
		}
	}
}

// checkAlive : The one and only test to see if a robot is dead and sets status to dead.
func checkAlive(n int) {
	if Robots[n].Status == DEAD {
		return
	}
	if Robots[n].Damage >= 100 {
		Robots[n].Status = DEAD
	}

}
func robotStatusToint() int {
	var a int
	switch numberOfRobots {

	case 4:
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
	case 3:
		if Robots[0].Status == ALIVE {
			a = a | 1
		}
		if Robots[1].Status == ALIVE {
			a = a | 2
		}
		if Robots[2].Status == ALIVE {
			a = a | 4
		}
	case 2:
		if Robots[0].Status == ALIVE {
			a = a | 1
		}
		if Robots[1].Status == ALIVE {
			a = a | 2
		}
	}

	return a
}

// endCondition : Check if its end game
// Last update, I totally screws this up.
func endCondition() bool {

	a := robotStatusToint()
	//fmt.Println(a)
	if teams {
		if a == 0 || a == 1 || a == 2 || a == 3 || a == 4 || a == 8 || a == 12 {
			return true
		} else {
			return false
		}
	} else {
		if a == 0 || a == 1 || a == 2 || a == 4 || a == 8 {
			return true
		} else {
			return false
		}
	}

	//return false

}

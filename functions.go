package main

import (
	"basicbots/builtin"
	"basicbots/object"
	"fmt"
	"math"
)

// FunctionTeam : Return the team number and set it is on a team.
func FunctionTeam(env builtin.Environment, args []object.Object) object.Object {
	var t float64

	if current == 0 {
		t = 1
		teams[0] = 1
	}

	if current == 1 {
		t = 1
		teams[1] = 1
	}

	if current == 2 {
		t = 2
		teams[2] = 2
	}

	if current == 3 {
		t = 2
		teams[3] = 2
	}

	return &object.NumberObject{Value: t}
}

// FunctionLocX : Basic statement. LOCX returns the current Y location.
func FunctionLocX(env builtin.Environment, args []object.Object) object.Object {
	X := Robots[current].X
	//if debug {
	//	fmt.Fprintf(os.Stderr, "Robot:%d FunctionLocX() LOCX:%5.2f\n", current, Robots[current].X)
	//}
	return &object.NumberObject{Value: X}
}

// FunctionLocY : Basic statement. LOCY returns the current Y location
func FunctionLocY(env builtin.Environment, args []object.Object) object.Object {
	y := Robots[current].Y
	//if debug {
	//	fmt.Fprintf(os.Stderr, "Robot:%d FunctionLocY() LOCY:%5.2f\n", current, Robots[current].Y)
	//}
	return &object.NumberObject{Value: y}
}

// FunctionSpeed : Basic statement. SPEED returns the current speed of the robot
func FunctionSpeed(env builtin.Environment, args []object.Object) object.Object {
	speed := Robots[current].Speed
	//if debug {
	//	fmt.Fprintf(os.Stderr, "Robot:%d FunctionSpeed() Speed:%5.2f\n", current, Robots[current].Speed)
	//}
	return &object.NumberObject{Value: speed}
}

// FunctionDamage : Basic statement. DAMAGE returns the current damage of the robot
func FunctionDamage(env builtin.Environment, args []object.Object) object.Object {
	damage := float64(Robots[current].Damage)
	//if debug {
	//	fmt.Fprintf(os.Stderr, "Robot:%d FunctionDamage() Damage:%d\n", current, Robots[current].Damage)
	//}
	return &object.NumberObject{Value: damage}
}

// FunctionDrive : Basic statement. DRIVE direction,speed sets speed and direction
func FunctionDrive(env builtin.Environment, args []object.Object) object.Object {
	var s, d float64

	if args[0].Type() == object.NUMBER {
		d = args[0].(*object.NumberObject).Value
	}
	if args[1].Type() == object.NUMBER {
		s = args[1].(*object.NumberObject).Value
	}

	if s < 0.0 {
		s = 0.0
	}
	if s > 100.0 {
		s = 100.0
	}

	d = math.Mod(d, 360.0)
	if d < 0.0 {
		d += 360.0
	}
	if d > 360.0 {
		d -= 360.0
	}

	Robots[current].SpeedWanted = s
	Robots[current].HeadingWanted = d

	//if debug {
	//	fmt.Fprintf(os.Stderr, "Robot:%d FunctionDrive heading:%5.2f speedwanted:%5.2f\n", current, Robots[current].HeadingWanted, Robots[current].SpeedWanted)
	//}

	return &object.NumberObject{Value: 0.0}
}

// FunctionScan : Basic statement. SCAN direction,width. Scan the battlefield in direction with a width of +/- width
func FunctionScan(env builtin.Environment, args []object.Object) object.Object {
	var angle, width float64

	if args[0].Type() == object.NUMBER {
		angle = args[0].(*object.NumberObject).Value
	}
	if args[1].Type() == object.NUMBER {
		width = args[1].(*object.NumberObject).Value
	}

	angle = math.Mod(angle, 360.0)
	if angle < 0 {
		angle += 360.0
	}
	if angle > 360.0 {
		angle -= 360

	}

	//if debug {
	//	fmt.Fprintf(os.Stderr, "SCAN angle: %f\n", angle)
	//}

	if width > 10.0 {
		width = 10.0
	}
	if width < 2.0 {
		width = 2.0
	}
	Robots[current].Scan = angle
	Robots[current].Width = width

	checkAlive(current)
	if Robots[current].Status == DEAD {
		return &object.NumberObject{Value: 0.0}
	}
	x1 := Robots[current].X
	y1 := Robots[current].Y

	td := 1000.0
	for i := 0; i < numberOfRobots; i++ {
		// fmt.Println("i=", i, numberOfRobots)

		checkAlive(i)

		if Robots[i].Status == DEAD {
			continue
		}

		if i == current {
			continue
		}

		if teams[i] != -1 {
			if teams[i] == teams[current] {
				//fmt.Println(i, current, teams[i], teams[current])
				continue
			}
		}

		x2 := Robots[i].X
		y2 := Robots[i].Y
		t := Scanner(angle, width, x1, y1, x2, y2)
		if t != 0.0 {
			if t < td {
				td = t
			}
		}
	}
	if td == 1000 { // would be fun if the limit of the scan was 100 units longer than missile range. maybe.
		td = 0
	}

	//if debug {
	//	fmt.Fprintf(os.Stderr, "Robot:%d FunctionScan() angle:%5.2f width:%5.2f result:%5.2f\n", current, angle, width, td)
	//}

	return &object.NumberObject{Value: td}

}

// FunctionCannon : Basic statement. CANNON direction, range. Fire the cannon at angle and distance. Do nothing if no missiles available.
func FunctionCannon(env builtin.Environment, args []object.Object) object.Object {

	var angle, rang float64

	if Robots[current].Reload != 0 {
		return &object.NumberObject{Value: 1.0}
	}

	if args[0].Type() == object.NUMBER {
		angle = args[0].(*object.NumberObject).Value
	}
	if args[1].Type() == object.NUMBER {
		rang = args[1].(*object.NumberObject).Value
	}
	if rang < 1 {
		// do nothing
		return &object.NumberObject{Value: 0.0}
	}
	if rang > MISSLERANGE {
		rang = MISSLERANGE
	}
	angle = math.Mod(angle, 360.0)
	if angle < 0 {
		angle += 360.0
	}
	if angle > 360.0 {
		angle -= 360.0
	}

	for m := 0; m < MAXMISSILES; m++ {
		if Missiles[current][m].Status == AVAILABLE && Missiles[current][m].Reload == 0 {
			Missiles[current][m].Status = FLYING
			Missiles[current][m].Distance = rang
			Missiles[current][m].Heading = angle
			Missiles[current][m].XOrigin = Robots[current].X
			Missiles[current][m].YOrigin = Robots[current].Y
			Missiles[current][m].XO = Robots[current].X
			Missiles[current][m].YO = Robots[current].Y
			Missiles[current][m].X = Robots[current].X
			Missiles[current][m].Y = Robots[current].Y
			//if debug {
			//	fmt.Fprintf(os.Stderr, "Robot:%d FunctionCannon() missile:%d status:%d angle:%5.2f range:%5.2f OX:%5.2f OY:%5.2f\n",
			//		current, m, Missiles[current][m].Status, angle, rang, Missiles[current][m].XOrigin, Missiles[current][m].YOrigin)
			//}
			Robots[current].Reload = ROBOTRELOAD
			break
		}
	}

	return &object.NumberObject{Value: 1.0}

}

func FunctionOut(env builtin.Environment, args []object.Object) object.Object {

	var outmessage string

	if args[0].Type() == object.STRING {
		outmessage = args[0].(*object.StringObject).Value
	}

	if current == 0 || current == 1 {
		select {
		case team1 <- outmessage:
		default:
		}
	} else {
		select {
		case team2 <- outmessage:
		default:
		}
	}
	return &object.NumberObject{Value: 0.0}
}

func FunctionIn(env builtin.Environment, args []object.Object) object.Object {
	var msg string

	if current == 0 || current == 1 {
		select {
		case msg = <-team1:
			// fmt.Println("received message", msg)
		default:
		}
	} else {
		select {
		case msg = <-team2:
			// fmt.Println("received message", msg)
		default:
		}
	}

	return &object.StringObject{Value: msg}

}

// STR converts a number to a string
func FunctionSTRC(env builtin.Environment, args []object.Object) object.Object {

	var num, c int
	// Error?
	if args[0].Type() == object.ERROR {
		return args[0]
	}

	// Already a string?
	if args[0].Type() == object.STRING {
		return args[0]
	}

	if args[0].Type() == object.NUMBER {
		i := args[0].(*object.NumberObject).Value
		num = int(i)
	}

	if args[1].Type() == object.NUMBER {
		i := args[1].(*object.NumberObject).Value
		c = int(i)
	}

	// Get the value
	s := fmt.Sprintf("%%%d%d", c, num)

	return &object.StringObject{Value: s}
}

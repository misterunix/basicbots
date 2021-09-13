package main

import (
	"basicbots/builtin"
	"basicbots/object"
	"math"
)

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
		checkAlive(i)
		if Robots[i].Status == DEAD {
			continue
		}
		if i == current {
			continue
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

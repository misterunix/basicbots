package main

// Robot : type struct for holding robot variables
type Robot struct {
	Name string // Name : The name of the robot

	X        float64 // X : Current X of the robot.
	Y        float64 // Y : Current Y of the robot.
	XOrigin  float64 // XOrigin : orgin x location. Used in movement.
	YOrigin  float64 // YOrigin : orgin y location. Used in movement.
	XPlotOld int     // XPlotOld : Last X plot on the battlefield. Used to remove the past marker.
	YPlotOld int     // YPlotOld : Last Y plot on the battlefield. Used to remove the past marker.

	Damage int // Damage : Current state of the damage of the robot

	Speed       float64 // Speed : Current speed of the robot.
	SpeedWanted float64 // SpeedWanted : The desired speed of the robot.
	SpeedHold   float64 // SpeedHold : Holds the previous speed while in a turn.

	Heading       float64 // Heading : Current heading of the robot
	HeadingWanted float64 // HeadingWanted : The disired heading of the robot.
	Distance      float64 // Distance : The Distance the robot has traveled from the origin point.

	Scan  float64 // Scan : Scan heading. 0-360
	Width float64 // Width : Scan width  2-10

	Cannon float64 // Cannon : Vannon heading
	Reload int     // Reload : Countdown until cannon reload is complete

	Status int // Status : Status of the robot. Dead or Alive.

	Winner int // Winner : Number of wins
	Lose   int // Lose : Number off times lost
	Tie    int // Tie : Number of times tied
	Points int // Points : Number of points this run.

	Program []byte // Program : Byte slice holding the program's text file.
}

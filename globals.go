package main

import (
	"basicbots/eval"
	"basicbots/tokenizer"

	"github.com/gdamore/tcell/v2"
)

// Nasty global variables and constants go here

// variables
var debug bool         // debug : Debug flag
var trace bool         // trace : Trace flag
var battledisplay bool // battledisplay : true show graphics

var numberOfRobots int // numberOfRobots : The number of robots in this simulation
var cycles int         // cycles : The number of cpu cycles
var Robots []Robot     // Robots : Array of the robots
var current int        // current : The current active robot

var token []*tokenizer.Tokenizer // token : slice of tokienizers

var evaluator []*eval.Interpreter // evaluator : Slice of Interpreters.

var maxCycles int // maxCycles : Maximum numer of cycles per match. Can change with cli flag.

var Missiles [MAXROBOTS][MAXMISSILES]Missile // Missile : Array of the missiles that can be used.

var defStyle tcell.Style // defStyle : white on blank for text. Used in tcell.
//var boxStyle tcell.Style // boxStyle : Leagacy purple
var scr tcell.Screen // scr : tcell screen interface. Using global to keep from haveing go routines.

var event = make(chan int)    // event : Channel for getting out of tcell with escape key.
var team1 = make(chan string) // team1 : Team 1 comm channel
var team2 = make(chan string) // team2 : Team 2 comm channel

var lox float64     // lox : scaling factor for battlefield to console area.
var loy float64     // loy : scalling factor for battlefield to console area.
var battleSizeX int // battleSizeX : Size of the battlescreen on the console in the X direction.
var battleSizeY int // battleSizeY : Size of the battlescreen on the console in the Y direction.

var cycledelay int64 // cycledelay : Delay in nanoseconds. Used in the battlescreen mode to slow down the play.

var matchcount int // matchcount : Number of matches to play with current robots. Can not be used with 'battledisplay'.

var etype int       // etype : Holds the event type from tcell or function when needing to exit the program.
var exiterror error // exiterror : Holds any error code that is causing the program to end.

/*
	CONSTANTS
*/

const VERSION = "v0.0.1b"

// Constants for cycles routines. Mainly movements.
const (
	MOVECLICKS = 50 // CLICKS : Number of cycles between moverobot and movemissile.
	//CLICK      = 100
)

// Constants for battlefield.
const (
	MAXX = 1000.0 // MAXX : Maximum size of the battlefield in the X direction.
	MAXY = 1000.0 // MAXY : Maximum size of the battlefield in the Y direction.
)

// Constants for robot status.
const (
	MAXROBOTS = 4    // MAXROBOTS : The maximum number of robots allowed.
	ALIVE     = 0    // ALIVE : Robot is functional.
	DEAD      = 1    // DEAD : Robot is dead.
	ACCEL     = 10.0 // ACCEL : Max Acceleration.
)

// Constants for converting degrees to radians and back.
const (
	DEG2RAD = 0.0174532925  // DEG2RAD : multiply 0-360 degrees to get radians.
	RAD2DEG = 57.2957795130 // RAD2GEG : multiply radians to get degrees.
)

// Constants for missile reload and movement.
const (
	MAXMISSILES  = 2     // MAXMISSILES : Maximum number if missiles a robot can have on the battlefield at one time.
	MISSLERANGE  = 700.0 // MISSLERANGE : Maxumum range of a missile.
	RELOAD       = 15    // RELOAD : Number of of movrment cycles for a missile reload.
	MISSILESPEED = 500.0 // MISSILESPEED : Missiles move at full speed. No ramp up. 500 is 500% vs 100% for robots.
	ROBOTRELOAD  = 5     // ROBOTRELOAD : Number of cycles for the robot to reload. Used to slow down the firing of the second missile.
	EXPLODECOUNT = 5     // EXPLODECOUNT : The number of movement cycles for the explosion to show in the battlescreen.
)

// Constants for missile status
const (
	AVAILABLE = 1 // AVAILABLE : Missle is available for firing
	FLYING    = 2 // FLYING : Missile is in flight.
	EXPLODE   = 3 // EXPLODE : Missile is the process of exploding.
	EXPLODING = 4 // EXPLODING : Blowing up
)

// Constants for robots running into things.
const (
	DAMAGEWALL = 2 // DAMAGEWALL : Amount of damage a robot will occure when it hits a wall.
	DAMAGECOL  = 4 // DAMAGECOL : Amount of damage a robot will occure when it hits another robot. Both robots take damamge.
)

// Constants for missle dmaage and blast radius.
const (
	MISFAR   = 40 // MISFAR : Largest blast radius to cause damage.
	MISNEAR  = 20 // MISNEAR : Medium blast radius to cause damage.
	MISCLOSE = 5  // MISCLOSE : Closest blast radius to cause damage.
	DAMFAR   = 3  // DAMFAR : Damage to robot at between MISNEAR and MISFAR
	DAMNEAR  = 5  // DAMNEAR : Damage to robot at between MISNEAR and MISCLOSE
	DAMCLOSE = 10 // DAMCLOSE : Damage to robot when at or below MISCLOSE
)

// Constants for channel events
const (
	ESCKEY     = 99
	BASICERROR = 98
)

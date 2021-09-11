package main

// Missle : type struct for holding missle variables.
type Missile struct {
	X            float64 // X : Current X location of the missile.
	Y            float64 // Y : Current Y location of the missile.
	XOrigin      float64 // XOrigin : Origin point from the last X location. Used in movement for directions.
	YOrigin      float64 // YOrigin : Origin point from the lasy Y location. Used in movement for ditections.
	XO           float64 // XO : Origin point from the last X location. Used in movement for directions.
	YO           float64 // YO : Origin point from the lasy Y location. Used in movement for ditections.
	XPlotOld     int     // XPlotOld : Last battlefield X position. Used for erasing the old marker.
	YPlotOld     int     // YPlotOld : Last battlefield Y position. Used for erasing the old marker.
	Heading      float64 // Heading : Heading of the missile.
	Distance     float64 // Distance : The set distance at which the missile explodes.
	Status       int     // Status : The status of the missile. Available,flying,exploding.
	Reload       int     // Reload : Movement clicks until reloaded.
	ExplodeCount int     // ExplodeCount : Time in motion cycles to have explosion showing.
}

package main

import (
	"math"
)

// GetAngle : Return the angle in degrees from x1,y1 to x2,y2
func GetAngle(x1, y1, x2, y2 float64) float64 {

	//a := math.Atan2(y2-y1, x2-x1) * RAD2DEG
	//a = math.Abs(math.Mod(a, 360.0))
	// fmt.Printf("\n%5.2f %5.2f %5.2f %5.2f %5.2f\n", x1, y1, x2, y2, a)
	//return a

	a := math.Atan2(y2-y1, x2-x1) * RAD2DEG

	if a < 0 {
		a += 360
	}
	if a > 360 {
		a -= 360
	}

	return a

}

// GetDistance : Return the distance between x1,y1 to x2,y2
func GetDistance(x1, y1, x2, y2 float64) float64 {
	d := math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
	return d
}

// Scanner : Scan angle with within x,y of both robots.
// Returns distance if robot found or 0 if none found.
func Scanner(angle, width, b1x, b1y, b2x, b2y float64) float64 {

	a := GetAngle(b1x, b1y, b2x, b2y)
	d := GetDistance(b1x, b1y, b2x, b2y)

	if width > 10.0 {
		width = 10.0
	}

	if width < 2.0 {
		width = 2.0
	}

	widthL := angle - width
	widthH := angle + width

	lowadj := widthL + 360.0
	highadj := widthH + 360.0

	var botAngleAdj float64

	botAngleAdj = a + 360
	if botAngleAdj <= widthH {
		botAngleAdj += 360
	}

	if botAngleAdj >= lowadj && botAngleAdj <= highadj {
		if d < 0 {
			d = math.Abs(d)
		}
		return d
	}

	return 0

}

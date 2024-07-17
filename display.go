package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

// Initialize tcell
func initDisplay() error {
	var err error
	defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	//boxStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	// Creat a new tcell screen
	scr, err = tcell.NewScreen()
	if err != nil {
		return err // Something went wrong, pass the error down stream.
	}

	// Init the tcell screen
	if err = scr.Init(); err != nil {
		return err
	}

	// set the default style, white on black screen.
	scr.SetStyle(defStyle)
	scr.Clear()

	// Draw initial boxes
	battleSizeX = 80 - 20
	battleSizeY = 23
	drawBox(scr, 0, 0, battleSizeX, battleSizeY, defStyle, "Battlefield")
	lox = (float64(battleSizeX) - 2.0) / MAXX
	loy = (float64(battleSizeY) - 2.0) / MAXY

	return nil
}

// Draws a box bording x1,y1 & x2,y2. Uses style for forground and background colors. text puts a string in the upper left corner od the box.
func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	//drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
	drawText(s, x1+2, y1, x1+len(text)+2, y2, style, text)
}

// Puts text onthe screen at x1,y1 to x2,y2 using style for forground and background colors.
func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

// Go routine for processing tcell events. Channel 'event' passes
// back to main program to terminate.
func eventloop() {
	for {
		ev := scr.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				event <- ESCKEY // Escape key
				break
			}
		}
	}

}

// Erase and Draw the robots, missiles and blasts on the
// battlefield as well as the side status.
func plotbattlefield() {
	posx := battleSizeX + 2
	var posy int

	for r := 0; r < numberOfRobots; r++ {

		posy = 5 * r // Y position of each robots running stats. 5 lines per robot

		checkAlive(r)
		//if Robots[r].Status == DEAD {
		//	continue
		//}

		if Robots[r].XPlotOld != 0 && Robots[r].YPlotOld != 0 {
			tg := rune(' ')
			scr.SetCell(Robots[r].XPlotOld, Robots[r].YPlotOld, defStyle, tg)
		}

		xf := Robots[r].X * lox
		yf := Robots[r].Y * loy
		x := int(xf) + 1
		y := int(yf) + 1

		if Robots[r].Status != DEAD {
			t := []rune(strconv.Itoa(r + 1))
			scr.SetCell(x, y, defStyle, t[0])
		}
		Robots[r].XPlotOld = x
		Robots[r].YPlotOld = y

		// Missles

		for m := 0; m < MAXMISSILES; m++ {
			if Missiles[r][m].Status == FLYING {
				mxf := Missiles[r][m].X * lox
				myf := Missiles[r][m].Y * loy
				mx := int(mxf) + 1
				my := int(myf) + 1
				if Missiles[r][m].XPlotOld != 0 || Missiles[r][m].YPlotOld != 0 {
					tgme := rune(' ')
					if isBorder(Missiles[r][m].XPlotOld, Missiles[r][m].YPlotOld) {
						scr.SetCell(Missiles[r][m].XPlotOld, Missiles[r][m].YPlotOld, defStyle, tgme)
					}
				}
				tgm := rune('.')
				if isBorder(mx, my) {
					scr.SetCell(mx, my, defStyle, tgm)
				}
				Missiles[r][m].XPlotOld = mx
				Missiles[r][m].YPlotOld = my
			}
			if Missiles[r][m].Status == EXPLODE {
				mxf := Missiles[r][m].X * lox
				myf := Missiles[r][m].Y * loy
				mx := int(mxf) + 1
				my := int(myf) + 1
				// Switching to function
				//if Missiles[r][m].XPlotOld != 0 || Missiles[r][m].YPlotOld != 0 {
				if isBorder(Missiles[r][m].XPlotOld, Missiles[r][m].YPlotOld) {
					tgme := rune(' ')
					scr.SetCell(Missiles[r][m].XPlotOld, Missiles[r][m].YPlotOld, defStyle, tgme)
				}
				//}
				tgm := rune('*')
				if isBorder(mx, my) {
					scr.SetCell(mx, my, defStyle, tgm)
				}
				Missiles[r][m].XPlotOld = mx
				Missiles[r][m].YPlotOld = my
			}

			if Missiles[r][m].Status == EXPLODING {
				mxf := Missiles[r][m].X * lox
				myf := Missiles[r][m].Y * loy
				mx := int(mxf) + 1
				my := int(myf) + 1
				//	if Missiles[r][m].XPlotOld != 0 && Missiles[r][m].YPlotOld != 0 {
				//		tgme := rune(' ')
				//		scr.SetCell(Missiles[r][m].XPlotOld, Missiles[r][m].YPlotOld, defStyle, tgme)
				//	}
				for j := 0; j < 3; j++ {
					for k := 0; k < 3; k++ {
						tgm := rune('*')
						if isBorder(mx-1+k, my+j) {
							scr.SetCell(mx-1+k, my+j, defStyle, tgm)
						}
					}
				}
				//				tgm := rune('*')
				//scr.SetCell(mx, my, defStyle, tgm)
				//Missiles[r][m].XPlotOld = mx
				//	Missiles[r][m].YPlotOld = my
			}

			if Missiles[r][m].ExplodeCount == EXPLODECOUNT {
				mxf := Missiles[r][m].X * lox
				myf := Missiles[r][m].Y * loy
				mx := int(mxf) + 1
				my := int(myf) + 1
				//	if Missiles[r][m].XPlotOld != 0 && Missiles[r][m].YPlotOld != 0 {
				//		tgme := rune(' ')
				//		scr.SetCell(Missiles[r][m].XPlotOld, Missiles[r][m].YPlotOld, defStyle, tgme)
				//	}
				for j := 0; j < 3; j++ {
					for k := 0; k < 3; k++ {
						tgm := rune(' ')
						if isBorder(mx-1+k, my+j) {
							scr.SetCell(mx-1+k, my+j, defStyle, tgm)
						}
					}
				}
				//				tgm := rune('*')
				//scr.SetCell(mx, my, defStyle, tgm)
				//Missiles[r][m].XPlotOld = mx
				//	Missiles[r][m].YPlotOld = my
			}

		}

		dd := Robots[r].Name // fmt.Sprintf("%s", Robots[r].Name)
		drawText(scr, posx, posy, posx+17, posy, defStyle, dd)
		posy++
		dd = fmt.Sprintf("SCN %03d WTH %03d", int(Robots[r].Scan), int(Robots[r].Width))
		drawText(scr, posx, posy, posx+17, posy, defStyle, dd)
		posy++
		dd = fmt.Sprintf("HNG %03d SPD %03d", int(Robots[r].Heading), int(Robots[r].Speed))
		drawText(scr, posx, posy, posx+17, posy, defStyle, dd)
		posy++
		dd = fmt.Sprintf("CAN %03d DAM %03d", int(Robots[r].Cannon), Robots[r].Damage)
		drawText(scr, posx, posy, posx+17, posy, defStyle, dd)
		posy++
		dd = fmt.Sprintf(" X  %03d  Y  %03d", int(Robots[r].X), int(Robots[r].Y))
		drawText(scr, posx, posy, posx+17, posy, defStyle, dd)
		posy++

		cycleString := fmt.Sprintf("Cycles %d", cycles)
		drawText(scr, posx, battleSizeY-2, posx+len(cycleString)+1, battleSizeY-2, defStyle, cycleString)

		//fps := fmt.Sprintf("%s", timeBucket)
		//drawText(scr, posx, battleSizeY-1, posx+len(fps)+1, battleSizeY-1, defStyle, fps)
	}

}

func isBorder(x, y int) bool {

	if x >= battleSizeX {
		return false
	}
	if x < 1 {
		return false
	}

	if y >= battleSizeY {
		return false
	}
	if y < 1 {
		return false
	}

	return true

}

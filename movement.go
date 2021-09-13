package main

import (
	"math"
)

// moverobot : Move robots every motion cycle.
// Take damage, detect dead robots, detect for collision.
func moverobot() {

	for i := 0; i < numberOfRobots; i++ {
		checkAlive(i)
		if Robots[i].Status == DEAD {
			continue
		}

		//  seems I need to check for walls before doing anything
		if Robots[i].X < 0.0 {
			Robots[i].X = 0.0
			Robots[i].XOrigin = Robots[i].X
			Robots[i].Damage += DAMAGEWALL
			Robots[i].Speed = 0.0
			Robots[i].SpeedWanted = 0.0
			Robots[i].SpeedHold = 0.0
			Robots[i].HeadingWanted = Robots[i].Heading
		}
		if Robots[i].Y < 0.0 {
			Robots[i].Y = 0.0
			Robots[i].YOrigin = Robots[i].Y
			Robots[i].Damage += DAMAGEWALL
			Robots[i].Speed = 0.0
			Robots[i].SpeedWanted = 0.0
			Robots[i].SpeedHold = 0.0
			Robots[i].HeadingWanted = Robots[i].Heading

		}
		if Robots[i].X > MAXX {
			Robots[i].X = MAXX
			Robots[i].XOrigin = Robots[i].X
			Robots[i].Damage += DAMAGEWALL
			Robots[i].Speed = 0.0
			Robots[i].SpeedWanted = 0.0
			Robots[i].SpeedHold = 0.0
			Robots[i].HeadingWanted = Robots[i].Heading
		}
		if Robots[i].Y > MAXY {
			Robots[i].Y = MAXY
			Robots[i].YOrigin = Robots[i].Y
			Robots[i].Damage += DAMAGEWALL
			Robots[i].Speed = 0.0
			Robots[i].SpeedWanted = 0.0
			Robots[i].SpeedHold = 0.0
			Robots[i].HeadingWanted = Robots[i].Heading
		}

		if Robots[i].Heading != Robots[i].HeadingWanted {
			if Robots[i].Speed > 50.0 {
				if Robots[i].SpeedHold == 0.0 {
					Robots[i].SpeedHold = Robots[i].SpeedWanted
				}
				Robots[i].SpeedWanted = 50.0
			} else {
				Robots[i].Heading = Robots[i].HeadingWanted
				if Robots[i].SpeedHold != 0.0 {
					Robots[i].SpeedWanted = Robots[i].SpeedHold
					Robots[i].SpeedHold = 0.0
				}
			}
		}

		if Robots[i].SpeedWanted > Robots[i].Speed {
			if Robots[i].Speed < 1.0 {
				Robots[i].Speed = 1.1
			}

			t := Robots[i].Speed
			velocity := math.Log10(t) / 1.5
			Robots[i].Speed += velocity
			if Robots[i].Speed > Robots[i].SpeedWanted {
				Robots[i].Speed = Robots[i].SpeedWanted
			}
		}

		if Robots[i].SpeedWanted < Robots[i].Speed {
			t := Robots[i].Speed
			if t > 1 && Robots[i].SpeedWanted != 0 {
				velocity := math.Log10(t) / 1.2
				Robots[i].Speed -= velocity
				if Robots[i].Speed < Robots[i].SpeedWanted {
					Robots[i].Speed = Robots[i].SpeedWanted
				}
			} else {
				Robots[i].Speed = 0.0
			}
		}

		if Robots[i].Speed < 0.0 {
			Robots[i].Speed = 0.0
		}
		if Robots[i].Speed > 100.0 {
			Robots[i].Speed = 100.0
		}

		Robots[i].X = Robots[i].XOrigin + math.Cos(Robots[i].Heading*DEG2RAD)*(Robots[i].Speed/100.0)
		Robots[i].Y = Robots[i].YOrigin + math.Sin(Robots[i].Heading*DEG2RAD)*(Robots[i].Speed/100.0)

		Robots[i].XOrigin = Robots[i].X
		Robots[i].YOrigin = Robots[i].Y

		//if debug {
		//			fmt.Fprintf(os.Stderr, "Robot:%d cycle:%d Hd:%5.2f HdW:%5.2f x:%5.2f y:%5.2f \n",
		//i, cycles, Robots[i].Heading, Robots[i].HeadingWanted, Robots[i].X, Robots[i].Y)
		//}

		for rn := 0; rn < numberOfRobots; rn++ {
			if rn == i {
				continue
			}
			checkAlive(rn)
			if Robots[rn].Status == DEAD {
				continue
			}

			if int(Robots[i].X) == int(Robots[rn].X) && int(Robots[i].Y) == int(Robots[rn].Y) {
				a1 := Robots[i].Heading
				a2 := a1 - 180
				if a2 < 0 {
					a2 += 360.0
				}
				if a2 > 360 {
					a2 -= 360.0
				}

				// get the offset of the current robot and add the inverse to the one hit
				tx := math.Cos(Robots[i].Heading*DEG2RAD) * 2.0
				ty := math.Sin(Robots[i].Heading*DEG2RAD) * 2.0
				Robots[rn].X += tx
				Robots[rn].Y += ty
				Robots[rn].XOrigin = Robots[rn].X
				Robots[rn].YOrigin = Robots[rn].Y
				Robots[rn].Speed = 0.0
				Robots[rn].SpeedWanted = 0.0
				Robots[i].Speed = 0.0
				Robots[i].SpeedWanted = 0.0
				Robots[i].Damage += DAMAGECOL
				Robots[rn].Damage += DAMAGECOL
				// move unit back one from the direction they came and set speed to 0
			}

			// Check current robot damage and break out of current collision loop
			checkAlive(i)
			if Robots[i].Status == DEAD {
				break
			}
		}
	}

}

// movemissile : move all flying missiles per motion click. Check for blast damage, htting wall, etc.
func movemissile() {
	//if debug {
	//	fmt.Fprintf(os.Stderr, "Cycles %d movemissile()\n", cycles)
	//}
	// missiles fly even if the robot that shot it is dead.
	for r := 0; r < numberOfRobots; r++ {

		for m := 0; m < MAXMISSILES; m++ {

			if Missiles[r][m].Status == AVAILABLE {
				if Missiles[r][m].Reload > 0 {
					Missiles[r][m].Reload--
					continue
				} else {
					Missiles[r][m].Reload = RELOAD
				}
			}

			if Missiles[r][m].Status == FLYING {
				// 	v := math.Log10(MISSILESPEED)
				Missiles[r][m].X = Missiles[r][m].XO + math.Cos(Missiles[r][m].Heading*DEG2RAD)*(MISSILESPEED/100.0)
				Missiles[r][m].Y = Missiles[r][m].YO + math.Sin(Missiles[r][m].Heading*DEG2RAD)*(MISSILESPEED/100.0)
				Missiles[r][m].XO = Missiles[r][m].X
				Missiles[r][m].YO = Missiles[r][m].Y
				//if debug {
				//		fmt.Fprintf(os.Stderr, "Robot:%d movemissile() move Missile:%d X:%5.2f Y:%5.2f xo:%5.2f yo:%5.2f range:%5.2f\n",
				//		r, m, Missiles[r][m].X, Missiles[r][m].Y, Missiles[r][m].XOrigin, Missiles[r][m].YOrigin, Missiles[r][m].Distance)
				//	}
				if Missiles[r][m].X <= 0.0 {
					Missiles[r][m].X = 0.0
					Missiles[r][m].Status = EXPLODE
				}
				if Missiles[r][m].Y <= 0.0 {
					Missiles[r][m].Y = 0.0
					Missiles[r][m].Status = EXPLODE
				}
				if Missiles[r][m].X >= MAXX {
					Missiles[r][m].X = MAXX - 1
					Missiles[r][m].Status = EXPLODE
				}
				if Missiles[r][m].Y >= MAXY {
					Missiles[r][m].Y = MAXY - 1
					Missiles[r][m].Status = EXPLODE
				}

				tx := Missiles[r][m].X - Missiles[r][m].XOrigin
				ty := Missiles[r][m].Y - Missiles[r][m].YOrigin
				d := math.Sqrt((tx * tx) + (ty * ty))
				if d >= Missiles[r][m].Distance || d >= MISSLERANGE {
					Missiles[r][m].Status = EXPLODE
					//if debug {
					//	fmt.Fprintf(os.Stderr, "Robot:%d movemissile() Distance test Missile:%d distance %5.2f range:%5.2f x:%5.2f xo:%5.2f y:%5.2f yo:%5.2f tx:%5.2f ty:%5.2f\n",
					//		r, m, d, Missiles[r][m].Distance, Missiles[r][m].X, Missiles[r][m].XOrigin, Missiles[r][m].Y, Missiles[r][m].YOrigin, tx, ty)
					//}
				}

			}

			if Missiles[r][m].Status == EXPLODING {
				//if debug {
				//	fmt.Fprintf(os.Stderr, "Robot:%d movemissile() Missile:%d Exploding1 X:%5.2f Y:%5.2f\n", r, m, Missiles[r][m].X, Missiles[r][m].Y)
				//}
				if Missiles[r][m].ExplodeCount == EXPLODECOUNT {
					Missiles[r][m].Reload = RELOAD
					Missiles[r][m].Status = AVAILABLE
					Missiles[r][m].ExplodeCount = 0
				} else {
					Missiles[r][m].ExplodeCount++
				}
			}

			if Missiles[r][m].Status == EXPLODE {
				Missiles[r][m].Status = EXPLODING
				//if debug {
				//		fmt.Fprintf(os.Stderr, "Robot:%d movemissile() Missile:%d Exploding X:%5.2f Y:%5.2f\n", r, m, Missiles[r][m].X, Missiles[r][m].Y)
				//	}

				for i := 0; i < numberOfRobots; i++ {
					checkAlive(i)
					if Robots[i].Status == DEAD {
						continue
					}
					if i == r {
						continue
					}
					tx := Missiles[r][m].X - Robots[i].X
					ty := Missiles[r][m].Y - Robots[i].Y
					d := math.Sqrt((tx * tx) + (ty * ty))
					switch {
					case d <= MISCLOSE:
						Robots[i].Damage += DAMCLOSE
					case d <= MISNEAR:
						Robots[i].Damage += DAMNEAR
					case d <= MISFAR:
						Robots[i].Damage += DAMFAR
					}

					checkAlive(i)

				}

			}

		}

	}

}

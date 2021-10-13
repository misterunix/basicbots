//go:build darwin

/* delay
Unix code for timing delay. Only included when OS is linux.
Sleep is not accurate and the delay needs to be accurate for setting the number of times the battlescreen is updated, thus setting the speed of the game.
*/
package delay

import "time"

// Delay : delay for d nanoseconds
func Delay(d int64) {
	time.Sleep(time.Duration(d) * time.Nanosecond)
	/*
		var t syscall.Timespec
		var b syscall.Timespec
		t.Sec = 0
		t.Nsec = d
		//starttime := time.Now()
		//for i := 0; i < 1000; i++ {
		_ = syscall.Nanosleep(&t, &b)
		//	_ = syscall.Times
		//}
		//d := time.Since(starttime)
		//fmt.Println("E", d)
		//fmt.Println(b)
	*/
}

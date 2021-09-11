//go:build linux || darwin

/* delay
Unix code for timing delay. Only included when OS is linux.
*/
package delay

import (
	"syscall"
)

// Delay : delay for d nanoseconds
func Delay(d int64) {
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
}

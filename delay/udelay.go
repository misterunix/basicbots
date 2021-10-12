//go:build darwin || linux

/* delay
Unix code for timing delay. Only included when OS is linux or darwin
*/
package delay

import (
	"time"
)

func Delay(d int64) {
	time.Sleep(time.Duration(d))
}

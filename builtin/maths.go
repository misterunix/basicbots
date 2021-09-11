// The builtin package provides the ability to register our built-in functions.
//
// maths.go implements our math-related primitives
// All trig funcitons are passed and returned as degrees
//gobasic is Copyrighted by Steve Kemp.
//Licensed under gpl-v2. Used by permission.
//You can find the original 'gobasic' at https://github.com/skx/gobasic

package builtin

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"basicbots/object"
)

// For math
const (
	DEG2RAD = 0.0174532925
	RAD2DEG = 57.2957795130
)

// init ensures that we've initialized our random-number state
func init() {
	rand.Seed(time.Now().UnixNano())
}

// ABS implements ABS
func ABS(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// If less than zero make it positive.
	if i < 0 {
		return &object.NumberObject{Value: -1 * i}
	}

	// Otherwise return as-is.
	return &object.NumberObject{Value: i}
}

// ACS (arccosine)
func ACS(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value
	r := math.Acos(i * DEG2RAD)
	r = math.Mod(r, 360)
	if r < 0.0 {
		r += 360.0
	}
	if r > 360.0 {
		r -= 360.0
	}
	return &object.NumberObject{Value: r}
}

// ASN (arcsine)
func ASN(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value
	r := math.Asin(i * DEG2RAD)
	r = math.Mod(r, 360)
	if r < 0.0 {
		r += 360.0
	}
	if r > 360.0 {
		r -= 360.0
	}
	return &object.NumberObject{Value: r}
}

// ATN (arctan)
func ATN(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value
	r := math.Atan(i * DEG2RAD)
	r = math.Mod(r, 360)
	if r < 0.0 {
		r += 360.0
	}
	if r > 360.0 {
		r -= 360.0
	}
	return &object.NumberObject{Value: r}
}

// BIN converts a number from binary.
func BIN(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	s := fmt.Sprintf("%d", int(i))

	b, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return object.Error("BIN:%s", err.Error())
	}

	return &object.NumberObject{Value: float64(b)}

}

// COS implements the COS function..
func COS(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value
	r := math.Cos(i * DEG2RAD)
	r = math.Mod(r, 360)
	if r < 0.0 {
		r += 360.0
	}
	if r > 360.0 {
		r -= 360.0
	}
	return &object.NumberObject{Value: r}
}

// EXP x=e^x EXP
func EXP(env Environment, args []object.Object) object.Object {
	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Exp(i)}
}

// INT implements INT
func INT(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Truncate.
	return &object.NumberObject{Value: float64(int(i))}
}

// LN calculates logarithms to the base e - LN
func LN(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Log(i)}
}

// PI returns the value of PI
func PI(env Environment, args []object.Object) object.Object {
	return &object.NumberObject{Value: math.Pi}
}

// RND implements RND
func RND(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// convert to int
	n := int(i)

	// Ensure it is valid.
	if n < 1 {
		return object.Error("Argument to RND must be >0")
	}

	// Return the random number
	return &object.NumberObject{Value: float64(rand.Intn(n))}
}

// SGN is the sign function (sometimes called signum).
func SGN(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	if i < 0 {
		return &object.NumberObject{Value: -1}
	}
	if i == 0 {
		return &object.NumberObject{Value: 0}
	}
	return &object.NumberObject{Value: 1}

}

// SIN operats the sin function.
func SIN(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value
	r := math.Sin(i * DEG2RAD)
	r = math.Mod(r, 360)
	if r < 0.0 {
		r += 360.0
	}
	if r > 360.0 {
		r -= 360.0
	}
	return &object.NumberObject{Value: r}
}

// SQR implements square root.
// Modified 2021 by Bill Jones for use in basicbots.
func SQR(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Ensure it is valid.
	if i < 1 {
		// return object.Error("Argument to SQR must be >0")
		return &object.NumberObject{Value: 0}
	}
	return &object.NumberObject{Value: math.Sqrt(i)}
}

// TAN implements the tan function.
func TAN(env Environment, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	r := math.Tan(i * DEG2RAD)
	r = math.Mod(r, 360)
	if r < 0.0 {
		r += 360.0
	}
	if r > 360.0 {
		r -= 360.0
	}
	return &object.NumberObject{Value: r}
}

// ATN2 (arctan)
// Added 2021 by Bill Jones for use in basicbots.
func ATN2(env Environment, args []object.Object) object.Object {

	// Get the (x float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	y := args[0].(*object.NumberObject).Value

	// Get the (x float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	x := args[1].(*object.NumberObject).Value

	r := math.Atan2(y, x) * RAD2DEG
	r = math.Mod(r, 360.0)
	if r < 0.0 {
		r += 360.0
	}
	if r > 360.0 {
		r -= 360.0
	}

	return &object.NumberObject{Value: r}
}

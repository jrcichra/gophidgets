package phidgets

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"

//Common functions that convert different types for this package

func float32ToCdouble(f float32) C.double {
	return C.double(f)
}

func boolToCInt(b bool) C.int {
	var r C.int
	if b {
		r = 1
	} else {
		r = 0
	}
	return r
}

func intToBool(i int) bool {
	var b bool
	if i > 0 {
		b = true
	} else {
		b = false
	}
	return b
}

func intToCInt(i int) C.int {
	var c C.int
	c = (C.int)(i)
	return c
}

func cIntToint(c C.int) int {
	var i int
	i = (int)(c)
	return i
}

func cIntTobool(c C.int) bool {
	i := cIntToint(c)
	return intToBool(i)
}

func uintToCUInt(i uint) C.uint {
	var c C.uint
	c = (C.uint)(i)
	return c
}

func cDoubleTofloat32(d C.double) float32 {
	var f float32
	f = (float32)(d)
	return f
}

func stringToCCharArray(s string) *C.char {
	return C.CString(s)
}

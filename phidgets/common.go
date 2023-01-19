package phidgets

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L . -lphidget22
#include <stdlib.h>
#include <phidget22.h>
*/
import "C"
import (
	"reflect"
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

// Passthrough - Go struct that passes through the phidget context callback, giving us a Go phidget pointer and the function we should callback to
type Passthrough struct {
	f func(float64)
}

// MotionPassthrough - has more than one float64 value as a parameter
type MotionPassthrough struct {
	f func([]float64, float64)
}

// SoundPassthrough - has more than one float64 value as a parameter
type SoundPassthrough struct {
	f func(float64, float64, float64, []float64)
}

// DistancePassthrough
type DistancePassthrough struct {
	f func(uint32)
}

// ReflectionPassthrough
type ReflectionPassthrough struct {
	f func([8]uint32, [8]uint32, uint32)
}

//export callback
func callback(handle unsafe.Pointer, ctx unsafe.Pointer, value C.double) {
	passthrough := gopointer.Restore(ctx).(Passthrough)
	p2 := passthrough.f
	p2(float64(value))
}

func carray2slice(array *C.double, len int) []C.double {
	var list []C.double
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&list)))
	sliceHeader.Cap = len
	sliceHeader.Len = len
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	return list
}

//export soundcallback
func soundcallback(handle unsafe.Pointer, ctx unsafe.Pointer, dB C.double, dBA C.double, dBC C.double, octaves *C.double) {
	passthrough := gopointer.Restore(ctx).(SoundPassthrough)
	p2 := passthrough.f
	var slce []float64
	length := 10
	cslce := carray2slice(octaves, length)
	for i := 0; i < length; i++ {
		slce = append(slce, float64(cslce[i]))
	}

	p2(float64(dB), float64(dBA), float64(dBC), slce)
}

// Common functions that convert different types for this package
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

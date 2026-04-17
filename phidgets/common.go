package phidgets

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L . -lphidget22
#include <stdlib.h>
#include <phidget22.h>
void callback(void*, void*, double);
void soundcallback(void*, void*, double, double, double, double*);
void motioncallback(void*, void*, double*, double);
void quaternioncallback(void*, void*, double*, double);
void voidcallback(void*, void*);
void encodercallback(void*, void*, int, double, int);
void countcallback(void*, void*, uint64_t, double);
void twofloatcallback(void*, void*, double, double);
void threefloatcallback(void*, void*, double, double, double);
void spatialcallback(void*, void*, double*, double*, double*, double);
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

// MotionPassthrough - has more than one float64 value as a parameter (array + timestamp)
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

// VoidPassthrough - callback with no arguments
type VoidPassthrough struct {
	f func()
}

// EncoderPassthrough - callback for encoder position changes (positionChange, timeChange, indexTriggered)
type EncoderPassthrough struct {
	f func(int, float64, bool)
}

// CountPassthrough - callback for frequency counter count changes (counts, timeChange)
type CountPassthrough struct {
	f func(uint64, float64)
}

// TwoFloatPassthrough - callback with two float64 values
type TwoFloatPassthrough struct {
	f func(float64, float64)
}

// ThreeFloatPassthrough - callback with three float64 values
type ThreeFloatPassthrough struct {
	f func(float64, float64, float64)
}

// SpatialPassthrough - callback for spatial sensor data (acceleration, angularRate, magneticField, timestamp)
type SpatialPassthrough struct {
	f func([]float64, []float64, []float64, float64)
}

//export callback
func callback(handle unsafe.Pointer, ctx unsafe.Pointer, value C.double) {
	passthrough := gopointer.Restore(ctx).(Passthrough)
	p2 := passthrough.f
	p2(float64(value))
}

//export motioncallback
func motioncallback(handle unsafe.Pointer, ctx unsafe.Pointer, arr *C.double, timestamp C.double) {
	passthrough := gopointer.Restore(ctx).(MotionPassthrough)
	passthrough.f(cDoubleArrayToSlice(arr, 3), float64(timestamp))
}

//export quaternioncallback
func quaternioncallback(handle unsafe.Pointer, ctx unsafe.Pointer, arr *C.double, timestamp C.double) {
	passthrough := gopointer.Restore(ctx).(MotionPassthrough)
	passthrough.f(cDoubleArrayToSlice(arr, 4), float64(timestamp))
}

//export voidcallback
func voidcallback(handle unsafe.Pointer, ctx unsafe.Pointer) {
	passthrough := gopointer.Restore(ctx).(VoidPassthrough)
	passthrough.f()
}

//export encodercallback
func encodercallback(handle unsafe.Pointer, ctx unsafe.Pointer, positionChange C.int, timeChange C.double, indexTriggered C.int) {
	passthrough := gopointer.Restore(ctx).(EncoderPassthrough)
	passthrough.f(int(positionChange), float64(timeChange), indexTriggered != 0)
}

//export countcallback
func countcallback(handle unsafe.Pointer, ctx unsafe.Pointer, counts C.uint64_t, timeChange C.double) {
	passthrough := gopointer.Restore(ctx).(CountPassthrough)
	passthrough.f(uint64(counts), float64(timeChange))
}

//export twofloatcallback
func twofloatcallback(handle unsafe.Pointer, ctx unsafe.Pointer, a C.double, b C.double) {
	passthrough := gopointer.Restore(ctx).(TwoFloatPassthrough)
	passthrough.f(float64(a), float64(b))
}

//export threefloatcallback
func threefloatcallback(handle unsafe.Pointer, ctx unsafe.Pointer, a C.double, b C.double, c C.double) {
	passthrough := gopointer.Restore(ctx).(ThreeFloatPassthrough)
	passthrough.f(float64(a), float64(b), float64(c))
}

//export spatialcallback
func spatialcallback(handle unsafe.Pointer, ctx unsafe.Pointer, accel *C.double, angularRate *C.double, magneticField *C.double, timestamp C.double) {
	passthrough := gopointer.Restore(ctx).(SpatialPassthrough)
	passthrough.f(
		cDoubleArrayToSlice(accel, 3),
		cDoubleArrayToSlice(angularRate, 3),
		cDoubleArrayToSlice(magneticField, 3),
		float64(timestamp),
	)
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

// cDoubleArrayToSlice converts a C double pointer to a Go float64 slice of the given length.
func cDoubleArrayToSlice(arr *C.double, length int) []float64 {
	result := make([]float64, length)
	base := uintptr(unsafe.Pointer(arr))
	size := unsafe.Sizeof(*arr)
	for i := 0; i < length; i++ {
		result[i] = float64(*(*C.double)(unsafe.Pointer(base + uintptr(i)*size)))
	}
	return result
}

func carray2slice(array *C.double, len int) []C.double {
	var list []C.double
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&list)))
	sliceHeader.Cap = len
	sliceHeader.Len = len
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	return list
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

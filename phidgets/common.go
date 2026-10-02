package phidgets

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L . -lphidget22
#include <stdlib.h>
#include <phidget22.h>
*/
import "C"
import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

//export callback
func callback(handle unsafe.Pointer, ctx unsafe.Pointer, value C.double) {
	gopointer.Restore(ctx).(func(float64))(float64(value))
}

//export motioncallback
func motioncallback(handle unsafe.Pointer, ctx unsafe.Pointer, arr *C.double, timestamp C.double) {
	gopointer.Restore(ctx).(func([]float64, float64))(cDoubles(arr, 3), float64(timestamp))
}

//export quaternioncallback
func quaternioncallback(handle unsafe.Pointer, ctx unsafe.Pointer, arr *C.double, timestamp C.double) {
	gopointer.Restore(ctx).(func([]float64, float64))(cDoubles(arr, 4), float64(timestamp))
}

//export voidcallback
func voidcallback(handle unsafe.Pointer, ctx unsafe.Pointer) {
	gopointer.Restore(ctx).(func())()
}

//export encodercallback
func encodercallback(handle unsafe.Pointer, ctx unsafe.Pointer, positionChange C.int, timeChange C.double, indexTriggered C.int) {
	gopointer.Restore(ctx).(func(int, float64, bool))(int(positionChange), float64(timeChange), indexTriggered != 0)
}

//export countcallback
func countcallback(handle unsafe.Pointer, ctx unsafe.Pointer, counts C.uint64_t, timeChange C.double) {
	gopointer.Restore(ctx).(func(uint64, float64))(uint64(counts), float64(timeChange))
}

//export twofloatcallback
func twofloatcallback(handle unsafe.Pointer, ctx unsafe.Pointer, a C.double, b C.double) {
	gopointer.Restore(ctx).(func(float64, float64))(float64(a), float64(b))
}

//export threefloatcallback
func threefloatcallback(handle unsafe.Pointer, ctx unsafe.Pointer, a C.double, b C.double, c C.double) {
	gopointer.Restore(ctx).(func(float64, float64, float64))(float64(a), float64(b), float64(c))
}

//export spatialcallback
func spatialcallback(handle unsafe.Pointer, ctx unsafe.Pointer, accel *C.double, angularRate *C.double, magneticField *C.double, timestamp C.double) {
	gopointer.Restore(ctx).(func([]float64, []float64, []float64, float64))(
		cDoubles(accel, 3), cDoubles(angularRate, 3), cDoubles(magneticField, 3), float64(timestamp))
}

//export statecallback
func statecallback(handle unsafe.Pointer, ctx unsafe.Pointer, state C.int) {
	gopointer.Restore(ctx).(func(bool))(state != 0)
}

//export uint32callback
func uint32callback(handle unsafe.Pointer, ctx unsafe.Pointer, value C.uint32_t) {
	gopointer.Restore(ctx).(func(uint32))(uint32(value))
}

//export errorcallback
func errorcallback(handle unsafe.Pointer, ctx unsafe.Pointer, code C.int, message *C.char) {
	gopointer.Restore(ctx).(func(int, string))(int(code), C.GoString(message))
}

//export rfidtagcallback
func rfidtagcallback(handle unsafe.Pointer, ctx unsafe.Pointer, tag *C.char, protocol C.int) {
	gopointer.Restore(ctx).(func(string, int))(C.GoString(tag), int(protocol))
}

//export dictkvcallback
func dictkvcallback(handle unsafe.Pointer, ctx unsafe.Pointer, key, value *C.char) {
	gopointer.Restore(ctx).(func(string, string))(C.GoString(key), C.GoString(value))
}

//export dictkeycallback
func dictkeycallback(handle unsafe.Pointer, ctx unsafe.Pointer, key *C.char) {
	gopointer.Restore(ctx).(func(string))(C.GoString(key))
}

//export ircodecallback
func ircodecallback(handle unsafe.Pointer, ctx unsafe.Pointer, code *C.char, bitCount C.uint32_t, isRepeat C.int) {
	gopointer.Restore(ctx).(func(string, uint32, bool))(C.GoString(code), uint32(bitCount), isRepeat != 0)
}

//export irrawcallback
func irrawcallback(handle unsafe.Pointer, ctx unsafe.Pointer, data *C.uint32_t, dataLen C.size_t) {
	n := int(dataLen)
	out := make([]uint32, n)
	for i, v := range unsafe.Slice(data, n) {
		out[i] = uint32(v)
	}
	gopointer.Restore(ctx).(func([]uint32))(out)
}

//export irlearncallback
func irlearncallback(handle unsafe.Pointer, ctx unsafe.Pointer, code *C.char, info unsafe.Pointer) {
	gopointer.Restore(ctx).(func(string, IRCodeInfo))(C.GoString(code), irCodeInfoFromC((*C.PhidgetIR_CodeInfo)(info)))
}

//export propcallback
func propcallback(handle unsafe.Pointer, ctx unsafe.Pointer, name *C.char) {
	gopointer.Restore(ctx).(func(string))(C.GoString(name))
}

//export soundcallback
func soundcallback(handle unsafe.Pointer, ctx unsafe.Pointer, dB C.double, dBA C.double, dBC C.double, octaves *C.double) {
	gopointer.Restore(ctx).(func(float64, float64, float64, []float64))(float64(dB), float64(dBA), float64(dBC), cDoubles(octaves, 10))
}

// cDoubles copies n C doubles into a Go slice (the C buffer is only valid during the callback).
func cDoubles(arr *C.double, n int) []float64 {
	out := make([]float64, n)
	for i, v := range unsafe.Slice(arr, n) {
		out[i] = float64(v)
	}
	return out
}

func getVec3(p *phidget, f func(*[3]C.double) C.PhidgetReturnCode) ([]float64, error) {
	return get(p, f, func(r [3]C.double) []float64 { return []float64{float64(r[0]), float64(r[1]), float64(r[2])} })
}

// get runs a C getter and converts its result, wrapping any error.
func get[R, T any](p *phidget, f func(*R) C.PhidgetReturnCode, conv func(R) T) (T, error) {
	var r R
	if cerr := f(&r); cerr != C.EPHIDGET_OK {
		var zero T
		return zero, p.phidgetError(cerr)
	}
	return conv(r), nil
}

func getDouble(p *phidget, f func(*C.double) C.PhidgetReturnCode) (float64, error) {
	return get(p, f, func(r C.double) float64 { return float64(r) })
}

func getUint32(p *phidget, f func(*C.uint32_t) C.PhidgetReturnCode) (uint32, error) {
	return get(p, f, func(r C.uint32_t) uint32 { return uint32(r) })
}

func getInt(p *phidget, f func(*C.int) C.PhidgetReturnCode) (int, error) {
	return get(p, f, func(r C.int) int { return int(r) })
}

func getBool(p *phidget, f func(*C.int) C.PhidgetReturnCode) (bool, error) {
	return get(p, f, func(r C.int) bool { return r != 0 })
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

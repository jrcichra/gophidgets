package phidgets

/*
#include <phidget22.h>
#include <stdlib.h>
#include <string.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetIR wraps a Phidget IR Transceiver
type PhidgetIR struct {
	phidget
	handle C.PhidgetIRHandle
}

// IRCodeInfo describes an IR code's transmission parameters
type IRCodeInfo struct {
	Encoding         IREncoding
	Length           IRLength
	Gap              uint32     // gap time (us)
	Trail            uint32     // trail time (us); 0 for none
	Header           [2]uint32  // header pulse and space (us)
	One              [2]uint32  // pulse and space times for a '1' bit (us)
	Zero             [2]uint32  // pulse and space times for a '0' bit (us)
	Repeat           [26]uint32 // repeat pulse/space times (us), null terminated
	MinRepeat        uint32     // minimum number of repeats on transmit
	DutyCycle        float64    // duty cycle (fraction)
	CarrierFrequency uint32     // carrier frequency (Hz)
	ToggleMask       string     // bit toggle mask
}

func uint32ToC(dst *C.uint32_t, src []uint32) {
	d := unsafe.Slice(dst, len(src))
	for i, v := range src {
		d[i] = C.uint32_t(v)
	}
}

func uint32FromC(src *C.uint32_t, dst []uint32) {
	s := unsafe.Slice(src, len(dst))
	for i := range dst {
		dst[i] = uint32(s[i])
	}
}

func irCodeInfoToC(info IRCodeInfo) *C.PhidgetIR_CodeInfo {
	c := (*C.PhidgetIR_CodeInfo)(C.malloc(C.size_t(unsafe.Sizeof(C.PhidgetIR_CodeInfo{}))))
	c.encoding = C.PhidgetIR_Encoding(info.Encoding)
	c.length = C.PhidgetIR_Length(info.Length)
	c.gap = C.uint32_t(info.Gap)
	c.trail = C.uint32_t(info.Trail)
	uint32ToC(&c.header[0], info.Header[:])
	uint32ToC(&c.one[0], info.One[:])
	uint32ToC(&c.zero[0], info.Zero[:])
	uint32ToC(&c.repeat[0], info.Repeat[:])
	c.minRepeat = C.uint32_t(info.MinRepeat)
	c.dutyCycle = C.double(info.DutyCycle)
	c.carrierFrequency = C.uint32_t(info.CarrierFrequency)
	if info.ToggleMask != "" {
		C.strcpy(&c.toggleMask[0], C.CString(info.ToggleMask))
	}
	return c
}

func irCodeInfoFromC(c *C.PhidgetIR_CodeInfo) IRCodeInfo {
	if c == nil {
		return IRCodeInfo{}
	}
	h := [2]uint32{}
	uint32FromC(&c.header[0], h[:])
	o := [2]uint32{}
	uint32FromC(&c.one[0], o[:])
	z := [2]uint32{}
	uint32FromC(&c.zero[0], z[:])
	re := [26]uint32{}
	uint32FromC(&c.repeat[0], re[:])
	return IRCodeInfo{
		Encoding:         IREncoding(c.encoding),
		Length:           IRLength(c.length),
		Gap:              uint32(c.gap),
		Trail:            uint32(c.trail),
		Header:           h,
		One:              o,
		Zero:             z,
		Repeat:           re,
		MinRepeat:        uint32(c.minRepeat),
		DutyCycle:        float64(c.dutyCycle),
		CarrierFrequency: uint32(c.carrierFrequency),
		ToggleMask:       C.GoString(&c.toggleMask[0]),
	}
}

// Create creates a PhidgetIR handle
func (p *PhidgetIR) Create() {
	C.PhidgetIR_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetLastCode returns the most recently received IR code (a hex string) and
// its bit count.
func (p *PhidgetIR) GetLastCode() (string, uint32, error) {
	var bitCount C.uint32_t
	code, err := readString(&p.phidget, C.IR_MAX_CODE_STR_LENGTH, func(buf *C.char, n C.size_t) C.PhidgetReturnCode {
		return C.PhidgetIR_getLastCode(p.handle, buf, n, &bitCount)
	})
	return code, uint32(bitCount), err
}

// GetLastLearnedCode returns the most recently learned IR code and its
// transmission parameters.
func (p *PhidgetIR) GetLastLearnedCode() (string, IRCodeInfo, error) {
	var info C.PhidgetIR_CodeInfo
	code, err := readString(&p.phidget, C.IR_MAX_CODE_STR_LENGTH, func(buf *C.char, n C.size_t) C.PhidgetReturnCode {
		return C.PhidgetIR_getLastLearnedCode(p.handle, buf, n, &info)
	})
	if err != nil {
		return "", IRCodeInfo{}, err
	}
	return code, irCodeInfoFromC(&info), nil
}

// Transmit sends an IR code with the given transmission parameters
func (p *PhidgetIR) Transmit(code string, info IRCodeInfo) error {
	ccode := C.CString(code)
	defer C.free(unsafe.Pointer(ccode))
	cinfo := irCodeInfoToC(info)
	defer C.free(unsafe.Pointer(cinfo))
	return p.phidgetError(C.PhidgetIR_transmit(p.handle, ccode, cinfo))
}

// TransmitRaw sends a raw sequence of IR timing samples
func (p *PhidgetIR) TransmitRaw(data []uint32, carrierFrequency uint32, dutyCycle float64, gap uint32) error {
	if len(data) == 0 {
		return p.phidgetError(C.EPHIDGET_INVALIDARG)
	}
	ptr := (*C.uint32_t)(unsafe.Pointer(&data[0]))
	return p.phidgetError(C.PhidgetIR_transmitRaw(p.handle, ptr, C.size_t(len(data)), C.uint32_t(carrierFrequency), C.double(dutyCycle), C.uint32_t(gap)))
}

// TransmitRepeat sends the most recently transmitted code as a repeat
func (p *PhidgetIR) TransmitRepeat() error {
	return p.phidgetError(C.PhidgetIR_transmitRepeat(p.handle))
}

// SetOnCodeHandler sets a callback that fires when an IR code is received.
// The callback receives the code, its bit count, and whether it is a repeat.
func (p *PhidgetIR) SetOnCodeHandler(f func(code string, bitCount uint32, isRepeat bool)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetIR_setOnCodeHandler(
		p.handle, (C.phidget_ircode_fcn)(unsafe.Pointer(C.ircodecallback)), ctx))
}

// SetOnLearnHandler sets a callback that fires when an IR code is learned.
// The callback receives the code and its transmission parameters.
func (p *PhidgetIR) SetOnLearnHandler(f func(code string, info IRCodeInfo)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetIR_setOnLearnHandler(
		p.handle, (C.phidget_irlearn_fcn)(unsafe.Pointer(C.irlearncallback)), ctx))
}

// SetOnRawDataHandler sets a callback that fires for raw IR timing samples.
// The callback receives the sample array.
func (p *PhidgetIR) SetOnRawDataHandler(f func(samples []uint32)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetIR_setOnRawDataHandler(
		p.handle, (C.phidget_irraw_fcn)(unsafe.Pointer(C.irrawcallback)), ctx))
}

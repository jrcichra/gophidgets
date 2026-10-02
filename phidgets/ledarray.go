package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetLEDArray wraps a Phidget LED Array
type PhidgetLEDArray struct {
	phidget
	handle C.PhidgetLEDArrayHandle
}

// LEDColor is an RGBW color value
type LEDColor struct {
	R uint8
	G uint8
	B uint8
	W uint8
}

// LEDAnimation describes an LED animation
type LEDAnimation struct {
	StartAddress uint32
	EndAddress   uint32
	Time         uint32 // time between changes (ms)
	Type         int    // PhidgetLEDArray_AnimationType
}

func ledColorToC(c LEDColor) C.PhidgetLEDArray_Color {
	var cc C.PhidgetLEDArray_Color
	cc.r = C.uint8_t(c.R)
	cc.g = C.uint8_t(c.G)
	cc.b = C.uint8_t(c.B)
	cc.w = C.uint8_t(c.W)
	return cc
}

// Create creates a PhidgetLEDArray handle
func (p *PhidgetLEDArray) Create() {
	C.PhidgetLEDArray_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetLED sets a single LED color
func (p *PhidgetLEDArray) SetLED(address uint32, color LEDColor, fadeTimeMS uint32) error {
	cc := ledColorToC(color)
	return p.phidgetError(C.PhidgetLEDArray_setLED(p.handle, C.uint32_t(address), &cc, C.uint32_t(fadeTimeMS)))
}

// SetLEDs sets a range of LED colors
func (p *PhidgetLEDArray) SetLEDs(startAddress, endAddress uint32, colors []LEDColor, fadeTimeMS uint32) error {
	if len(colors) == 0 {
		return p.phidgetError(C.EPHIDGET_INVALIDARG)
	}
	cbuf := make([]C.PhidgetLEDArray_Color, len(colors))
	for i, c := range colors {
		cbuf[i] = ledColorToC(c)
	}
	return p.phidgetError(C.PhidgetLEDArray_setLEDs(
		p.handle, C.uint32_t(startAddress), C.uint32_t(endAddress), &cbuf[0], C.size_t(len(cbuf)), C.uint32_t(fadeTimeMS)))
}

// ClearLEDs clears all LEDs to black
func (p *PhidgetLEDArray) ClearLEDs() error {
	return p.phidgetError(C.PhidgetLEDArray_clearLEDs(p.handle))
}

// SetAnimation starts an animation over the given LED color pattern
func (p *PhidgetLEDArray) SetAnimation(animationID int, pattern []LEDColor, animation LEDAnimation) error {
	if len(pattern) == 0 {
		return p.phidgetError(C.EPHIDGET_INVALIDARG)
	}
	cbuf := make([]C.PhidgetLEDArray_Color, len(pattern))
	for i, c := range pattern {
		cbuf[i] = ledColorToC(c)
	}
	var ca C.PhidgetLEDArray_Animation
	ca.startAddress = C.uint32_t(animation.StartAddress)
	ca.endAddress = C.uint32_t(animation.EndAddress)
	ca.time = C.uint32_t(animation.Time)
	ca.animationType = C.PhidgetLEDArray_AnimationType(animation.Type)
	return p.phidgetError(C.PhidgetLEDArray_setAnimation(
		p.handle, C.int32_t(animationID), &cbuf[0], C.size_t(len(cbuf)), &ca))
}

// StopAnimation stops the animation with the given ID
func (p *PhidgetLEDArray) StopAnimation(animationID int) error {
	return p.phidgetError(C.PhidgetLEDArray_stopAnimation(p.handle, C.int32_t(animationID)))
}

// SynchronizeAnimations synchronizes the playback of all active animations
func (p *PhidgetLEDArray) SynchronizeAnimations() error {
	return p.phidgetError(C.PhidgetLEDArray_synchronizeAnimations(p.handle))
}

// GetMinAnimationID returns the minimum valid animation ID
func (p *PhidgetLEDArray) GetMinAnimationID() (int, error) {
	var r C.int32_t
	if cerr := C.PhidgetLEDArray_getMinAnimationID(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// GetMaxAnimationID returns the maximum valid animation ID
func (p *PhidgetLEDArray) GetMaxAnimationID() (int, error) {
	var r C.int32_t
	if cerr := C.PhidgetLEDArray_getMaxAnimationID(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// GetMinAnimationPatternCount returns the minimum animation pattern length
func (p *PhidgetLEDArray) GetMinAnimationPatternCount() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetLEDArray_getMinAnimationPatternCount(p.handle, r)
	})
}

// GetMaxAnimationPatternCount returns the maximum animation pattern length
func (p *PhidgetLEDArray) GetMaxAnimationPatternCount() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetLEDArray_getMaxAnimationPatternCount(p.handle, r)
	})
}

// GetBrightness returns the LED array brightness
func (p *PhidgetLEDArray) GetBrightness() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLEDArray_getBrightness(p.handle, r) })
}

// SetBrightness sets the LED array brightness
func (p *PhidgetLEDArray) SetBrightness(brightness float64) error {
	return p.phidgetError(C.PhidgetLEDArray_setBrightness(p.handle, C.double(brightness)))
}

// GetMinBrightness returns the minimum brightness
func (p *PhidgetLEDArray) GetMinBrightness() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMinBrightness(p.handle, r) })
}

// GetMaxBrightness returns the maximum brightness
func (p *PhidgetLEDArray) GetMaxBrightness() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMaxBrightness(p.handle, r) })
}

// GetGamma returns the LED array gamma
func (p *PhidgetLEDArray) GetGamma() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLEDArray_getGamma(p.handle, r) })
}

// SetGamma sets the LED array gamma
func (p *PhidgetLEDArray) SetGamma(gamma float64) error {
	return p.phidgetError(C.PhidgetLEDArray_setGamma(p.handle, C.double(gamma)))
}

// GetMinGamma returns the minimum gamma
func (p *PhidgetLEDArray) GetMinGamma() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMinGamma(p.handle, r) })
}

// GetMaxGamma returns the maximum gamma
func (p *PhidgetLEDArray) GetMaxGamma() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMaxGamma(p.handle, r) })
}

// GetColorOrder returns the LED color order (PhidgetLEDArray_ColorOrder)
func (p *PhidgetLEDArray) GetColorOrder() (int, error) {
	return get(&p.phidget, func(r *C.PhidgetLEDArray_ColorOrder) C.PhidgetReturnCode {
		return C.PhidgetLEDArray_getColorOrder(p.handle, r)
	}, func(r C.PhidgetLEDArray_ColorOrder) int { return int(r) })
}

// SetColorOrder sets the LED color order (PhidgetLEDArray_ColorOrder)
func (p *PhidgetLEDArray) SetColorOrder(order int) error {
	return p.phidgetError(C.PhidgetLEDArray_setColorOrder(p.handle, C.PhidgetLEDArray_ColorOrder(order)))
}

// GetMinAddress returns the minimum LED address
func (p *PhidgetLEDArray) GetMinAddress() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMinAddress(p.handle, r) })
}

// GetMaxAddress returns the maximum LED address
func (p *PhidgetLEDArray) GetMaxAddress() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMaxAddress(p.handle, r) })
}

// GetMinLEDCount returns the minimum number of LEDs
func (p *PhidgetLEDArray) GetMinLEDCount() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMinLEDCount(p.handle, r) })
}

// GetMaxLEDCount returns the maximum number of LEDs
func (p *PhidgetLEDArray) GetMaxLEDCount() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMaxLEDCount(p.handle, r) })
}

// GetMinFadeTime returns the minimum fade time in milliseconds
func (p *PhidgetLEDArray) GetMinFadeTime() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMinFadeTime(p.handle, r) })
}

// GetMaxFadeTime returns the maximum fade time in milliseconds
func (p *PhidgetLEDArray) GetMaxFadeTime() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetLEDArray_getMaxFadeTime(p.handle, r) })
}

// SetPowerEnabled enables or disables power to the LEDs
func (p *PhidgetLEDArray) SetPowerEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetLEDArray_setPowerEnabled(p.handle, boolToCInt(enabled)))
}

// GetPowerEnabled returns whether power to the LEDs is enabled
func (p *PhidgetLEDArray) GetPowerEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetLEDArray_getPowerEnabled(p.handle, r) })
}

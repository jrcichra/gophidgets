package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetCapacitiveTouch wraps a Phidget capacitive touch sensor
type PhidgetCapacitiveTouch struct {
	phidget
	handle C.PhidgetCapacitiveTouchHandle
}

// Create creates a PhidgetCapacitiveTouch handle
func (p *PhidgetCapacitiveTouch) Create() {
	C.PhidgetCapacitiveTouch_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetTouchValue returns the current touch position (0.0–1.0)
func (p *PhidgetCapacitiveTouch) GetTouchValue() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetCapacitiveTouch_getTouchValue(p.handle, r) })
}

// GetIsTouched returns whether the sensor is currently being touched
func (p *PhidgetCapacitiveTouch) GetIsTouched() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetCapacitiveTouch_getIsTouched(p.handle, r) })
}

// SetSensitivity sets the touch sensitivity (0.0–1.0)
func (p *PhidgetCapacitiveTouch) SetSensitivity(sensitivity float64) error {
	return p.phidgetError(C.PhidgetCapacitiveTouch_setSensitivity(p.handle, C.double(sensitivity)))
}

// GetSensitivity returns the current touch sensitivity
func (p *PhidgetCapacitiveTouch) GetSensitivity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetCapacitiveTouch_getSensitivity(p.handle, r) })
}

// SetTouchValueChangeTrigger sets the change threshold that triggers the touch callback
func (p *PhidgetCapacitiveTouch) SetTouchValueChangeTrigger(trigger float64) error {
	return p.phidgetError(C.PhidgetCapacitiveTouch_setTouchValueChangeTrigger(p.handle, C.double(trigger)))
}

// GetTouchValueChangeTrigger returns the current touch value change trigger
func (p *PhidgetCapacitiveTouch) GetTouchValueChangeTrigger() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetCapacitiveTouch_getTouchValueChangeTrigger(p.handle, r)
	})
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetCapacitiveTouch) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetCapacitiveTouch_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetCapacitiveTouch) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetCapacitiveTouch_getDataInterval(p.handle, r) })
}

// SetOnTouchHandler sets a callback that fires when the sensor is touched; receives the touch value
func (p *PhidgetCapacitiveTouch) SetOnTouchHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetCapacitiveTouch_setOnTouchHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.callback)), ctx))
}

// SetOnTouchEndHandler sets a callback that fires when touch ends
func (p *PhidgetCapacitiveTouch) SetOnTouchEndHandler(f func()) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetCapacitiveTouch_setOnTouchEndHandler(
		p.handle, (C.phidget_void_fcn)(unsafe.Pointer(C.voidcallback)), ctx))
}

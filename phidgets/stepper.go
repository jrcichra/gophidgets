package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetStepper wraps a Phidget stepper motor controller
type PhidgetStepper struct {
	phidget
	handle C.PhidgetStepperHandle
}

// Create creates a PhidgetStepper handle
func (p *PhidgetStepper) Create() {
	C.PhidgetStepper_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetTargetPosition sets the target position in steps
func (p *PhidgetStepper) SetTargetPosition(pos float64) error {
	return p.phidgetError(C.PhidgetStepper_setTargetPosition(p.handle, C.double(pos)))
}

// GetTargetPosition returns the currently set target position
func (p *PhidgetStepper) GetTargetPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getTargetPosition(p.handle, r) })
}

// GetPosition returns the current motor position in steps
func (p *PhidgetStepper) GetPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getPosition(p.handle, r) })
}

// GetMinPosition returns the minimum position limit
func (p *PhidgetStepper) GetMinPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getMinPosition(p.handle, r) })
}

// GetMaxPosition returns the maximum position limit
func (p *PhidgetStepper) GetMaxPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getMaxPosition(p.handle, r) })
}

// AddPositionOffset adds an offset to the current position counter without moving
func (p *PhidgetStepper) AddPositionOffset(offset float64) error {
	return p.phidgetError(C.PhidgetStepper_addPositionOffset(p.handle, C.double(offset)))
}

// SetEngaged engages or disengages the stepper coils
func (p *PhidgetStepper) SetEngaged(engaged bool) error {
	return p.phidgetError(C.PhidgetStepper_setEngaged(p.handle, boolToCInt(engaged)))
}

// GetEngaged returns whether the stepper coils are engaged
func (p *PhidgetStepper) GetEngaged() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetStepper_getEngaged(p.handle, r) })
}

// GetIsMoving returns whether the stepper is currently moving
func (p *PhidgetStepper) GetIsMoving() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetStepper_getIsMoving(p.handle, r) })
}

// SetAcceleration sets the motor acceleration in steps/s²
func (p *PhidgetStepper) SetAcceleration(accel float64) error {
	return p.phidgetError(C.PhidgetStepper_setAcceleration(p.handle, C.double(accel)))
}

// GetAcceleration returns the current acceleration setting
func (p *PhidgetStepper) GetAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getAcceleration(p.handle, r) })
}

// SetVelocityLimit sets the maximum velocity in steps/s
func (p *PhidgetStepper) SetVelocityLimit(limit float64) error {
	return p.phidgetError(C.PhidgetStepper_setVelocityLimit(p.handle, C.double(limit)))
}

// GetVelocityLimit returns the velocity limit
func (p *PhidgetStepper) GetVelocityLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getVelocityLimit(p.handle, r) })
}

// GetVelocity returns the current velocity in steps/s
func (p *PhidgetStepper) GetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getVelocity(p.handle, r) })
}

// SetCurrentLimit sets the motor current limit in amps
func (p *PhidgetStepper) SetCurrentLimit(amps float64) error {
	return p.phidgetError(C.PhidgetStepper_setCurrentLimit(p.handle, C.double(amps)))
}

// GetCurrentLimit returns the current limit setting
func (p *PhidgetStepper) GetCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getCurrentLimit(p.handle, r) })
}

// SetHoldingCurrentLimit sets the current limit when the motor is holding position
func (p *PhidgetStepper) SetHoldingCurrentLimit(amps float64) error {
	return p.phidgetError(C.PhidgetStepper_setHoldingCurrentLimit(p.handle, C.double(amps)))
}

// SetRescaleFactor sets a scaling factor applied to positions and velocities
func (p *PhidgetStepper) SetRescaleFactor(factor float64) error {
	return p.phidgetError(C.PhidgetStepper_setRescaleFactor(p.handle, C.double(factor)))
}

// GetRescaleFactor returns the current rescale factor
func (p *PhidgetStepper) GetRescaleFactor() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetStepper_getRescaleFactor(p.handle, r) })
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetStepper) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetStepper_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetStepper) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetStepper_getDataInterval(p.handle, r) })
}

// SetOnPositionChangeHandler sets a callback that fires when the position changes
func (p *PhidgetStepper) SetOnPositionChangeHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetStepper_setOnPositionChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.callback)), ctx))
}

// SetOnVelocityChangeHandler sets a callback that fires when the velocity changes
func (p *PhidgetStepper) SetOnVelocityChangeHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetStepper_setOnVelocityChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.callback)), ctx))
}

// SetOnStoppedHandler sets a callback that fires when the stepper has stopped moving
func (p *PhidgetStepper) SetOnStoppedHandler(f func()) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetStepper_setOnStoppedHandler(
		p.handle, (C.phidget_void_fcn)(unsafe.Pointer(C.voidcallback)), ctx))
}

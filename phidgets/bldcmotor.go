package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

// PhidgetBLDCMotor wraps a Phidget brushless DC motor controller
type PhidgetBLDCMotor struct {
	phidget
	handle C.PhidgetBLDCMotorHandle
}

// Create creates a PhidgetBLDCMotor handle
func (p *PhidgetBLDCMotor) Create() {
	C.PhidgetBLDCMotor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetTargetVelocity sets the target velocity (-1.0 to 1.0; negative = reverse)
func (p *PhidgetBLDCMotor) SetTargetVelocity(velocity float64) error {
	return p.phidgetError(C.PhidgetBLDCMotor_setTargetVelocity(p.handle, C.double(velocity)))
}

// GetTargetVelocity returns the currently set target velocity
func (p *PhidgetBLDCMotor) GetTargetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getTargetVelocity(p.handle, r) })
}

// GetVelocity returns the current measured velocity
func (p *PhidgetBLDCMotor) GetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getVelocity(p.handle, r) })
}

// GetMinVelocity returns the minimum velocity
func (p *PhidgetBLDCMotor) GetMinVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getMinVelocity(p.handle, r) })
}

// GetMaxVelocity returns the maximum velocity
func (p *PhidgetBLDCMotor) GetMaxVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getMaxVelocity(p.handle, r) })
}

// SetAcceleration sets the rate of velocity change
func (p *PhidgetBLDCMotor) SetAcceleration(accel float64) error {
	return p.phidgetError(C.PhidgetBLDCMotor_setAcceleration(p.handle, C.double(accel)))
}

// GetAcceleration returns the current acceleration setting
func (p *PhidgetBLDCMotor) GetAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getAcceleration(p.handle, r) })
}

// GetMinAcceleration returns the minimum settable acceleration
func (p *PhidgetBLDCMotor) GetMinAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getMinAcceleration(p.handle, r) })
}

// GetMaxAcceleration returns the maximum settable acceleration
func (p *PhidgetBLDCMotor) GetMaxAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getMaxAcceleration(p.handle, r) })
}

// SetTargetBrakingStrength sets the braking strength when the motor stops (0.0–1.0)
func (p *PhidgetBLDCMotor) SetTargetBrakingStrength(strength float64) error {
	return p.phidgetError(C.PhidgetBLDCMotor_setTargetBrakingStrength(p.handle, C.double(strength)))
}

// GetBrakingStrength returns the current braking strength
func (p *PhidgetBLDCMotor) GetBrakingStrength() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getBrakingStrength(p.handle, r) })
}

// SetCurrentLimit sets the motor current limit in amps
func (p *PhidgetBLDCMotor) SetCurrentLimit(amps float64) error {
	return p.phidgetError(C.PhidgetBLDCMotor_setCurrentLimit(p.handle, C.double(amps)))
}

// GetCurrentLimit returns the current limit setting
func (p *PhidgetBLDCMotor) GetCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getCurrentLimit(p.handle, r) })
}

// GetPosition returns the current rotor position (cumulative)
func (p *PhidgetBLDCMotor) GetPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getPosition(p.handle, r) })
}

// AddPositionOffset adds an offset to the position counter without moving
func (p *PhidgetBLDCMotor) AddPositionOffset(offset float64) error {
	return p.phidgetError(C.PhidgetBLDCMotor_addPositionOffset(p.handle, C.double(offset)))
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetBLDCMotor) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetBLDCMotor_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetBLDCMotor) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetBLDCMotor_getDataInterval(p.handle, r) })
}

// SetOnVelocityUpdateHandler sets a callback that fires on each velocity update
func (p *PhidgetBLDCMotor) SetOnVelocityUpdateHandler(f func(float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetBLDCMotor_setOnVelocityUpdateHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// SetOnPositionChangeHandler sets a callback that fires when the position changes
func (p *PhidgetBLDCMotor) SetOnPositionChangeHandler(f func(float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetBLDCMotor_setOnPositionChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// SetOnBrakingStrengthChangeHandler sets a callback that fires when braking strength changes
func (p *PhidgetBLDCMotor) SetOnBrakingStrengthChangeHandler(f func(float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetBLDCMotor_setOnBrakingStrengthChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetBLDCMotor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetBLDCMotor_delete(&p.handle))
}

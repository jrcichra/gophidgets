package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetDCMotor wraps a Phidget DC motor controller
type PhidgetDCMotor struct {
	phidget
	handle C.PhidgetDCMotorHandle
}

// Create creates a PhidgetDCMotor handle
func (p *PhidgetDCMotor) Create() {
	C.PhidgetDCMotor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetTargetVelocity sets the target velocity (-1.0 to 1.0; negative = reverse)
func (p *PhidgetDCMotor) SetTargetVelocity(velocity float64) error {
	return p.phidgetError(C.PhidgetDCMotor_setTargetVelocity(p.handle, C.double(velocity)))
}

// GetTargetVelocity returns the currently set target velocity
func (p *PhidgetDCMotor) GetTargetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getTargetVelocity(p.handle, r) })
}

// GetVelocity returns the current measured velocity
func (p *PhidgetDCMotor) GetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getVelocity(p.handle, r) })
}

// GetMinVelocity returns the minimum velocity
func (p *PhidgetDCMotor) GetMinVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getMinVelocity(p.handle, r) })
}

// GetMaxVelocity returns the maximum velocity
func (p *PhidgetDCMotor) GetMaxVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getMaxVelocity(p.handle, r) })
}

// SetAcceleration sets the rate of velocity change (units/s²)
func (p *PhidgetDCMotor) SetAcceleration(accel float64) error {
	return p.phidgetError(C.PhidgetDCMotor_setAcceleration(p.handle, C.double(accel)))
}

// GetAcceleration returns the current acceleration setting
func (p *PhidgetDCMotor) GetAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getAcceleration(p.handle, r) })
}

// GetMinAcceleration returns the minimum settable acceleration
func (p *PhidgetDCMotor) GetMinAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getMinAcceleration(p.handle, r) })
}

// GetMaxAcceleration returns the maximum settable acceleration
func (p *PhidgetDCMotor) GetMaxAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getMaxAcceleration(p.handle, r) })
}

// SetTargetBrakingStrength sets the braking strength when the motor stops (0.0–1.0)
func (p *PhidgetDCMotor) SetTargetBrakingStrength(strength float64) error {
	return p.phidgetError(C.PhidgetDCMotor_setTargetBrakingStrength(p.handle, C.double(strength)))
}

// GetBrakingStrength returns the current braking strength
func (p *PhidgetDCMotor) GetBrakingStrength() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getBrakingStrength(p.handle, r) })
}

// SetCurrentLimit sets the motor current limit in amps
func (p *PhidgetDCMotor) SetCurrentLimit(amps float64) error {
	return p.phidgetError(C.PhidgetDCMotor_setCurrentLimit(p.handle, C.double(amps)))
}

// GetCurrentLimit returns the current limit setting
func (p *PhidgetDCMotor) GetCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getCurrentLimit(p.handle, r) })
}

// GetBackEMF returns the measured back-EMF voltage
func (p *PhidgetDCMotor) GetBackEMF() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDCMotor_getBackEMF(p.handle, r) })
}

// SetBackEMFSensingState enables or disables back-EMF sensing
func (p *PhidgetDCMotor) SetBackEMFSensingState(enabled bool) error {
	return p.phidgetError(C.PhidgetDCMotor_setBackEMFSensingState(p.handle, boolToCInt(enabled)))
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetDCMotor) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetDCMotor_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetDCMotor) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetDCMotor_getDataInterval(p.handle, r) })
}

// SetOnVelocityUpdateHandler sets a callback that fires on each velocity update
func (p *PhidgetDCMotor) SetOnVelocityUpdateHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDCMotor_setOnVelocityUpdateHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// SetOnBackEMFChangeHandler sets a callback that fires when back-EMF changes
func (p *PhidgetDCMotor) SetOnBackEMFChangeHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDCMotor_setOnBackEMFChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// SetOnBrakingStrengthChangeHandler sets a callback that fires when braking strength changes
func (p *PhidgetDCMotor) SetOnBrakingStrengthChangeHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDCMotor_setOnBrakingStrengthChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetDCMotor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetDCMotor_delete(&p.handle))
}

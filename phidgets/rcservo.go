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

// PhidgetRCServo wraps a Phidget RC servo controller channel
type PhidgetRCServo struct {
	phidget
	handle C.PhidgetRCServoHandle
}

// Create creates a PhidgetRCServo handle
func (p *PhidgetRCServo) Create() {
	C.PhidgetRCServo_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetTargetPosition sets the target position in degrees
func (p *PhidgetRCServo) SetTargetPosition(degrees float64) error {
	return p.phidgetError(C.PhidgetRCServo_setTargetPosition(p.handle, C.double(degrees)))
}

// GetTargetPosition returns the currently set target position
func (p *PhidgetRCServo) GetTargetPosition() (float64, error) {
	var r C.double
	if cerr := C.PhidgetRCServo_getTargetPosition(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetPosition returns the current servo position in degrees
func (p *PhidgetRCServo) GetPosition() (float64, error) {
	var r C.double
	if cerr := C.PhidgetRCServo_getPosition(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetMinPosition returns the minimum position limit
func (p *PhidgetRCServo) GetMinPosition() (float64, error) {
	var r C.double
	if cerr := C.PhidgetRCServo_getMinPosition(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetMaxPosition returns the maximum position limit
func (p *PhidgetRCServo) GetMaxPosition() (float64, error) {
	var r C.double
	if cerr := C.PhidgetRCServo_getMaxPosition(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetMinPosition overrides the minimum position limit
func (p *PhidgetRCServo) SetMinPosition(min float64) error {
	return p.phidgetError(C.PhidgetRCServo_setMinPosition(p.handle, C.double(min)))
}

// SetMaxPosition overrides the maximum position limit
func (p *PhidgetRCServo) SetMaxPosition(max float64) error {
	return p.phidgetError(C.PhidgetRCServo_setMaxPosition(p.handle, C.double(max)))
}

// SetEngaged engages or disengages the servo (must be engaged to move)
func (p *PhidgetRCServo) SetEngaged(engaged bool) error {
	return p.phidgetError(C.PhidgetRCServo_setEngaged(p.handle, boolToCInt(engaged)))
}

// GetEngaged returns whether the servo is engaged
func (p *PhidgetRCServo) GetEngaged() (bool, error) {
	var r C.int
	if cerr := C.PhidgetRCServo_getEngaged(p.handle, &r); cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

// GetIsMoving returns whether the servo is currently moving
func (p *PhidgetRCServo) GetIsMoving() (bool, error) {
	var r C.int
	if cerr := C.PhidgetRCServo_getIsMoving(p.handle, &r); cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

// SetAcceleration sets the servo acceleration in degrees/s²
func (p *PhidgetRCServo) SetAcceleration(accel float64) error {
	return p.phidgetError(C.PhidgetRCServo_setAcceleration(p.handle, C.double(accel)))
}

// GetAcceleration returns the current acceleration setting
func (p *PhidgetRCServo) GetAcceleration() (float64, error) {
	var r C.double
	if cerr := C.PhidgetRCServo_getAcceleration(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetVelocityLimit sets the maximum velocity in degrees/s
func (p *PhidgetRCServo) SetVelocityLimit(limit float64) error {
	return p.phidgetError(C.PhidgetRCServo_setVelocityLimit(p.handle, C.double(limit)))
}

// GetVelocityLimit returns the velocity limit
func (p *PhidgetRCServo) GetVelocityLimit() (float64, error) {
	var r C.double
	if cerr := C.PhidgetRCServo_getVelocityLimit(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetVelocity returns the current servo velocity in degrees/s
func (p *PhidgetRCServo) GetVelocity() (float64, error) {
	var r C.double
	if cerr := C.PhidgetRCServo_getVelocity(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetTorque sets the torque limit (0.0–1.0)
func (p *PhidgetRCServo) SetTorque(torque float64) error {
	return p.phidgetError(C.PhidgetRCServo_setTorque(p.handle, C.double(torque)))
}

// GetTorque returns the torque limit
func (p *PhidgetRCServo) GetTorque() (float64, error) {
	var r C.double
	if cerr := C.PhidgetRCServo_getTorque(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetSpeedRampingState enables or disables speed ramping
func (p *PhidgetRCServo) SetSpeedRampingState(enabled bool) error {
	return p.phidgetError(C.PhidgetRCServo_setSpeedRampingState(p.handle, boolToCInt(enabled)))
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetRCServo) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetRCServo_setDataInterval(p.handle, C.uint32_t(ms)))
}

// SetOnPositionChangeHandler sets a callback that fires when the position changes
func (p *PhidgetRCServo) SetOnPositionChangeHandler(f func(float64)) error {
	var pt Passthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetRCServo_setOnPositionChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// SetOnVelocityChangeHandler sets a callback that fires when the velocity changes
func (p *PhidgetRCServo) SetOnVelocityChangeHandler(f func(float64)) error {
	var pt Passthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetRCServo_setOnVelocityChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// SetOnTargetPositionReachedHandler sets a callback that fires when the target position is reached
func (p *PhidgetRCServo) SetOnTargetPositionReachedHandler(f func(float64)) error {
	var pt Passthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetRCServo_setOnTargetPositionReachedHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetRCServo) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetRCServo_delete(&p.handle))
}

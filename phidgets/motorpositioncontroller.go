package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetMotorPositionController wraps a Phidget motor position controller (stepper or brushless)
type PhidgetMotorPositionController struct {
	phidget
	handle C.PhidgetMotorPositionControllerHandle
}

// Create creates a PhidgetMotorPositionController handle
func (p *PhidgetMotorPositionController) Create() {
	C.PhidgetMotorPositionController_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetTargetPosition sets the target position
func (p *PhidgetMotorPositionController) SetTargetPosition(pos float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setTargetPosition(p.handle, C.double(pos)))
}

// GetTargetPosition returns the currently set target position
func (p *PhidgetMotorPositionController) GetTargetPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getTargetPosition(p.handle, r)
	})
}

// GetPosition returns the current motor position
func (p *PhidgetMotorPositionController) GetPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getPosition(p.handle, r)
	})
}

// GetMinPosition returns the minimum position limit
func (p *PhidgetMotorPositionController) GetMinPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getMinPosition(p.handle, r)
	})
}

// GetMaxPosition returns the maximum position limit
func (p *PhidgetMotorPositionController) GetMaxPosition() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getMaxPosition(p.handle, r)
	})
}

// SetEngaged engages or disengages the motor windings
func (p *PhidgetMotorPositionController) SetEngaged(engaged bool) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setEngaged(p.handle, boolToCInt(engaged)))
}

// GetEngaged returns whether the motor windings are engaged
func (p *PhidgetMotorPositionController) GetEngaged() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetMotorPositionController_getEngaged(p.handle, r) })
}

// SetAcceleration sets the rate of position change
func (p *PhidgetMotorPositionController) SetAcceleration(accel float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setAcceleration(p.handle, C.double(accel)))
}

// GetAcceleration returns the current acceleration setting
func (p *PhidgetMotorPositionController) GetAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getAcceleration(p.handle, r)
	})
}

// SetVelocityLimit sets the maximum velocity
func (p *PhidgetMotorPositionController) SetVelocityLimit(limit float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setVelocityLimit(p.handle, C.double(limit)))
}

// GetVelocityLimit returns the maximum velocity limit
func (p *PhidgetMotorPositionController) GetVelocityLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getVelocityLimit(p.handle, r)
	})
}

// SetCurrentLimit sets the motor current limit
func (p *PhidgetMotorPositionController) SetCurrentLimit(limit float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setCurrentLimit(p.handle, C.double(limit)))
}

// GetCurrentLimit returns the motor current limit
func (p *PhidgetMotorPositionController) GetCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getCurrentLimit(p.handle, r)
	})
}

// SetKp sets the proportional controller gain
func (p *PhidgetMotorPositionController) SetKp(kp float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setKp(p.handle, C.double(kp)))
}

// GetKp returns the proportional controller gain
func (p *PhidgetMotorPositionController) GetKp() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetMotorPositionController_getKp(p.handle, r) })
}

// SetKi sets the integral controller gain
func (p *PhidgetMotorPositionController) SetKi(ki float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setKi(p.handle, C.double(ki)))
}

// GetKi returns the integral controller gain
func (p *PhidgetMotorPositionController) GetKi() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetMotorPositionController_getKi(p.handle, r) })
}

// SetKd sets the derivative controller gain
func (p *PhidgetMotorPositionController) SetKd(kd float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setKd(p.handle, C.double(kd)))
}

// GetKd returns the derivative controller gain
func (p *PhidgetMotorPositionController) GetKd() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetMotorPositionController_getKd(p.handle, r) })
}

// SetDeadBand sets the deadband (units within the target that is considered on-target)
func (p *PhidgetMotorPositionController) SetDeadBand(deadband float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setDeadBand(p.handle, C.double(deadband)))
}

// GetDeadBand returns the deadband
func (p *PhidgetMotorPositionController) GetDeadBand() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getDeadBand(p.handle, r)
	})
}

// SetStallVelocity sets the velocity at which the motor is considered stalled
func (p *PhidgetMotorPositionController) SetStallVelocity(velocity float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setStallVelocity(p.handle, C.double(velocity)))
}

// GetStallVelocity returns the stall velocity
func (p *PhidgetMotorPositionController) GetStallVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getStallVelocity(p.handle, r)
	})
}

// SetRescaleFactor sets a scaling factor applied to positions and velocities
func (p *PhidgetMotorPositionController) SetRescaleFactor(factor float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setRescaleFactor(p.handle, C.double(factor)))
}

// GetRescaleFactor returns the current rescale factor
func (p *PhidgetMotorPositionController) GetRescaleFactor() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getRescaleFactor(p.handle, r)
	})
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetMotorPositionController) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetMotorPositionController) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getDataInterval(p.handle, r)
	})
}

// SetDataRate sets the data rate in Hertz
func (p *PhidgetMotorPositionController) SetDataRate(hz float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setDataRate(p.handle, C.double(hz)))
}

// GetDataRate returns the data rate in Hertz
func (p *PhidgetMotorPositionController) GetDataRate() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getDataRate(p.handle, r)
	})
}

// GetMinFailsafeTime returns the minimum settable failsafe time in milliseconds
func (p *PhidgetMotorPositionController) GetMinFailsafeTime() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getMinFailsafeTime(p.handle, r)
	})
}

// GetMaxFailsafeTime returns the maximum settable failsafe time in milliseconds
func (p *PhidgetMotorPositionController) GetMaxFailsafeTime() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getMaxFailsafeTime(p.handle, r)
	})
}

// SetFailsafeBrakingEnabled sets whether to hold the motor in place when the failsafe triggers
func (p *PhidgetMotorPositionController) SetFailsafeBrakingEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setFailsafeBrakingEnabled(p.handle, boolToCInt(enabled)))
}

// GetFailsafeBrakingEnabled returns whether braking is enabled on failsafe
func (p *PhidgetMotorPositionController) GetFailsafeBrakingEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getFailsafeBrakingEnabled(p.handle, r)
	})
}

// SetFailsafeCurrentLimit sets the current limit applied when the failsafe triggers
func (p *PhidgetMotorPositionController) SetFailsafeCurrentLimit(amps float64) error {
	return p.phidgetError(C.PhidgetMotorPositionController_setFailsafeCurrentLimit(p.handle, C.double(amps)))
}

// GetFailsafeCurrentLimit returns the failsafe current limit
func (p *PhidgetMotorPositionController) GetFailsafeCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorPositionController_getFailsafeCurrentLimit(p.handle, r)
	})
}

// EnableFailsafe enables the failsafe with the given time in milliseconds
func (p *PhidgetMotorPositionController) EnableFailsafe(ms uint32) error {
	return p.phidgetError(C.PhidgetMotorPositionController_enableFailsafe(p.handle, C.uint32_t(ms)))
}

// ResetFailsafe resets the failsafe state
func (p *PhidgetMotorPositionController) ResetFailsafe() error {
	return p.phidgetError(C.PhidgetMotorPositionController_resetFailsafe(p.handle))
}

// SetOnPositionChangeHandler sets a callback that fires when the position changes
func (p *PhidgetMotorPositionController) SetOnPositionChangeHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetMotorPositionController_setOnPositionChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// SetOnDutyCycleUpdateHandler sets a callback that fires on each duty cycle update
func (p *PhidgetMotorPositionController) SetOnDutyCycleUpdateHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetMotorPositionController_setOnDutyCycleUpdateHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetMotorPositionController) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetMotorPositionController_delete(&p.handle))
}

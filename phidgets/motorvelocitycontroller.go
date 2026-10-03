package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetMotorVelocityController wraps a Phidget motor velocity controller (brushless)
type PhidgetMotorVelocityController struct {
	phidget
	handle C.PhidgetMotorVelocityControllerHandle
}

// Create creates a PhidgetMotorVelocityController handle
func (p *PhidgetMotorVelocityController) Create() {
	C.PhidgetMotorVelocityController_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetTargetVelocity sets the target velocity
func (p *PhidgetMotorVelocityController) SetTargetVelocity(velocity float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setTargetVelocity(p.handle, C.double(velocity)))
}

// GetTargetVelocity returns the currently set target velocity
func (p *PhidgetMotorVelocityController) GetTargetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getTargetVelocity(p.handle, r)
	})
}

// GetVelocity returns the current measured velocity
func (p *PhidgetMotorVelocityController) GetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getVelocity(p.handle, r)
	})
}

// GetMaxTargetVelocity returns the maximum settable target velocity
func (p *PhidgetMotorVelocityController) GetMaxTargetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getMaxTargetVelocity(p.handle, r)
	})
}

// GetMinTargetVelocity returns the minimum settable target velocity
func (p *PhidgetMotorVelocityController) GetMinTargetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getMinTargetVelocity(p.handle, r)
	})
}

// SetAcceleration sets the rate of velocity change
func (p *PhidgetMotorVelocityController) SetAcceleration(accel float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setAcceleration(p.handle, C.double(accel)))
}

// GetAcceleration returns the current acceleration setting
func (p *PhidgetMotorVelocityController) GetAcceleration() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getAcceleration(p.handle, r)
	})
}

// SetCurrentLimit sets the maximum current limit
func (p *PhidgetMotorVelocityController) SetCurrentLimit(amps float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setCurrentLimit(p.handle, C.double(amps)))
}

// GetCurrentLimit returns the maximum current limit setting
func (p *PhidgetMotorVelocityController) GetCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getCurrentLimit(p.handle, r)
	})
}

// GetActiveCurrentLimit returns the current limit actually in effect
func (p *PhidgetMotorVelocityController) GetActiveCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getActiveCurrentLimit(p.handle, r)
	})
}

// SetSurgeCurrentLimit sets the peak (surge) current limit
func (p *PhidgetMotorVelocityController) SetSurgeCurrentLimit(amps float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setSurgeCurrentLimit(p.handle, C.double(amps)))
}

// GetSurgeCurrentLimit returns the peak (surge) current limit
func (p *PhidgetMotorVelocityController) GetSurgeCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getSurgeCurrentLimit(p.handle, r)
	})
}

// SetInductance sets the motor inductance
func (p *PhidgetMotorVelocityController) SetInductance(inductance float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setInductance(p.handle, C.double(inductance)))
}

// GetInductance returns the motor inductance
func (p *PhidgetMotorVelocityController) GetInductance() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getInductance(p.handle, r)
	})
}

// SetKp sets the proportional controller gain
func (p *PhidgetMotorVelocityController) SetKp(kp float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setKp(p.handle, C.double(kp)))
}

// GetKp returns the proportional controller gain
func (p *PhidgetMotorVelocityController) GetKp() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetMotorVelocityController_getKp(p.handle, r) })
}

// SetKi sets the integral controller gain
func (p *PhidgetMotorVelocityController) SetKi(ki float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setKi(p.handle, C.double(ki)))
}

// GetKi returns the integral controller gain
func (p *PhidgetMotorVelocityController) GetKi() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetMotorVelocityController_getKi(p.handle, r) })
}

// SetKd sets the derivative controller gain
func (p *PhidgetMotorVelocityController) SetKd(kd float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setKd(p.handle, C.double(kd)))
}

// GetKd returns the derivative controller gain
func (p *PhidgetMotorVelocityController) GetKd() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetMotorVelocityController_getKd(p.handle, r) })
}

// SetDeadBand sets the velocity deadband
func (p *PhidgetMotorVelocityController) SetDeadBand(deadband float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setDeadBand(p.handle, C.double(deadband)))
}

// GetDeadBand returns the velocity deadband
func (p *PhidgetMotorVelocityController) GetDeadBand() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getDeadBand(p.handle, r)
	})
}

// SetStallVelocity sets the velocity at which the motor is considered stalled
func (p *PhidgetMotorVelocityController) SetStallVelocity(velocity float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setStallVelocity(p.handle, C.double(velocity)))
}

// GetStallVelocity returns the stall velocity
func (p *PhidgetMotorVelocityController) GetStallVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getStallVelocity(p.handle, r)
	})
}

// SetEngaged engages or disengages the motor windings
func (p *PhidgetMotorVelocityController) SetEngaged(engaged bool) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setEngaged(p.handle, boolToCInt(engaged)))
}

// GetEngaged returns whether the motor windings are engaged
func (p *PhidgetMotorVelocityController) GetEngaged() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetMotorVelocityController_getEngaged(p.handle, r) })
}

// SetRescaleFactor sets a scaling factor applied to velocities
func (p *PhidgetMotorVelocityController) SetRescaleFactor(factor float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setRescaleFactor(p.handle, C.double(factor)))
}

// GetRescaleFactor returns the current rescale factor
func (p *PhidgetMotorVelocityController) GetRescaleFactor() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getRescaleFactor(p.handle, r)
	})
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetMotorVelocityController) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetMotorVelocityController) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getDataInterval(p.handle, r)
	})
}

// SetDataRate sets the data rate in Hertz
func (p *PhidgetMotorVelocityController) SetDataRate(hz float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setDataRate(p.handle, C.double(hz)))
}

// GetDataRate returns the data rate in Hertz
func (p *PhidgetMotorVelocityController) GetDataRate() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getDataRate(p.handle, r)
	})
}

// SetPositionType selects the position measurement source (Phidget_PositionType)
func (p *PhidgetMotorVelocityController) SetPositionType(positionType PositionType) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setPositionType(p.handle, C.Phidget_PositionType(positionType)))
}

// GetPositionType returns the position measurement source (Phidget_PositionType)
func (p *PhidgetMotorVelocityController) GetPositionType() (PositionType, error) {
	return get(&p.phidget, func(r *C.Phidget_PositionType) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getPositionType(p.handle, r)
	}, func(r C.Phidget_PositionType) PositionType { return PositionType(r) })
}

// SetFailsafeBrakingEnabled sets whether to hold the motor in place when the failsafe triggers
func (p *PhidgetMotorVelocityController) SetFailsafeBrakingEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setFailsafeBrakingEnabled(p.handle, boolToCInt(enabled)))
}

// GetFailsafeBrakingEnabled returns whether braking is enabled on failsafe
func (p *PhidgetMotorVelocityController) GetFailsafeBrakingEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getFailsafeBrakingEnabled(p.handle, r)
	})
}

// SetFailsafeCurrentLimit sets the current limit applied when the failsafe triggers
func (p *PhidgetMotorVelocityController) SetFailsafeCurrentLimit(amps float64) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_setFailsafeCurrentLimit(p.handle, C.double(amps)))
}

// GetFailsafeCurrentLimit returns the failsafe current limit
func (p *PhidgetMotorVelocityController) GetFailsafeCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMotorVelocityController_getFailsafeCurrentLimit(p.handle, r)
	})
}

// EnableFailsafe enables the failsafe with the given time in milliseconds
func (p *PhidgetMotorVelocityController) EnableFailsafe(ms uint32) error {
	return p.phidgetError(C.PhidgetMotorVelocityController_enableFailsafe(p.handle, C.uint32_t(ms)))
}

// ResetFailsafe resets the failsafe state
func (p *PhidgetMotorVelocityController) ResetFailsafe() error {
	return p.phidgetError(C.PhidgetMotorVelocityController_resetFailsafe(p.handle))
}

// SetOnVelocityChangeHandler sets a callback that fires when the velocity changes
func (p *PhidgetMotorVelocityController) SetOnVelocityChangeHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetMotorVelocityController_setOnVelocityChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.callback)), ctx))
}

// SetOnDutyCycleUpdateHandler sets a callback that fires on each duty cycle update
func (p *PhidgetMotorVelocityController) SetOnDutyCycleUpdateHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetMotorVelocityController_setOnDutyCycleUpdateHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.callback)), ctx))
}

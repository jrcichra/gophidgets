package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetPowerGuard wraps a Phidget Power Guard
type PhidgetPowerGuard struct {
	phidget
	handle C.PhidgetPowerGuardHandle
}

// Create creates a PhidgetPowerGuard handle
func (p *PhidgetPowerGuard) Create() {
	C.PhidgetPowerGuard_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetPowerEnabled enables or disables the output power
func (p *PhidgetPowerGuard) SetPowerEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetPowerGuard_setPowerEnabled(p.handle, boolToCInt(enabled)))
}

// GetPowerEnabled returns whether the output power is enabled
func (p *PhidgetPowerGuard) GetPowerEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetPowerGuard_getPowerEnabled(p.handle, r) })
}

// SetOverVoltage sets the over-voltage trip point in volts
func (p *PhidgetPowerGuard) SetOverVoltage(volts float64) error {
	return p.phidgetError(C.PhidgetPowerGuard_setOverVoltage(p.handle, C.double(volts)))
}

// GetOverVoltage returns the over-voltage trip point in volts
func (p *PhidgetPowerGuard) GetOverVoltage() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPowerGuard_getOverVoltage(p.handle, r) })
}

// GetMinOverVoltage returns the minimum settable over-voltage trip point in volts
func (p *PhidgetPowerGuard) GetMinOverVoltage() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPowerGuard_getMinOverVoltage(p.handle, r) })
}

// GetMaxOverVoltage returns the maximum settable over-voltage trip point in volts
func (p *PhidgetPowerGuard) GetMaxOverVoltage() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPowerGuard_getMaxOverVoltage(p.handle, r) })
}

// SetFanMode sets the fan mode (Phidget_FanMode)
func (p *PhidgetPowerGuard) SetFanMode(mode int) error {
	return p.phidgetError(C.PhidgetPowerGuard_setFanMode(p.handle, C.Phidget_FanMode(mode)))
}

// GetFanMode returns the fan mode (Phidget_FanMode)
func (p *PhidgetPowerGuard) GetFanMode() (int, error) {
	var r C.Phidget_FanMode
	if cerr := C.PhidgetPowerGuard_getFanMode(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// GetMinFailsafeTime returns the minimum settable failsafe time in milliseconds
func (p *PhidgetPowerGuard) GetMinFailsafeTime() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetPowerGuard_getMinFailsafeTime(p.handle, r) })
}

// GetMaxFailsafeTime returns the maximum settable failsafe time in milliseconds
func (p *PhidgetPowerGuard) GetMaxFailsafeTime() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetPowerGuard_getMaxFailsafeTime(p.handle, r) })
}

// EnableFailsafe enables the failsafe with the given time in milliseconds
func (p *PhidgetPowerGuard) EnableFailsafe(ms uint32) error {
	return p.phidgetError(C.PhidgetPowerGuard_enableFailsafe(p.handle, C.uint32_t(ms)))
}

// ResetFailsafe resets the failsafe state
func (p *PhidgetPowerGuard) ResetFailsafe() error {
	return p.phidgetError(C.PhidgetPowerGuard_resetFailsafe(p.handle))
}

// Close closes the handle and deletes it
func (p *PhidgetPowerGuard) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetPowerGuard_delete(&p.handle))
}

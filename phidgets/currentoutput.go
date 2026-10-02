package phidgets

/*
#include <phidget22.h>
*/
import "C"
import (
	"unsafe"
)

// PhidgetCurrentOutput wraps a Phidget current output (voltage-to-current converter)
type PhidgetCurrentOutput struct {
	phidget
	handle C.PhidgetCurrentOutputHandle
}

// Create creates a PhidgetCurrentOutput handle
func (p *PhidgetCurrentOutput) Create() {
	C.PhidgetCurrentOutput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetCurrent sets the output current in amps, bounded by the device's
// minimum and maximum current limits.
func (p *PhidgetCurrentOutput) SetCurrent(amps float64) error {
	return p.phidgetError(C.PhidgetCurrentOutput_setCurrent(p.handle, C.double(amps)))
}

// GetCurrent returns the output current in amps
func (p *PhidgetCurrentOutput) GetCurrent() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetCurrentOutput_getCurrent(p.handle, r) })
}

// GetMinCurrent returns the minimum settable current in amps
func (p *PhidgetCurrentOutput) GetMinCurrent() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetCurrentOutput_getMinCurrent(p.handle, r) })
}

// GetMaxCurrent returns the maximum settable current in amps
func (p *PhidgetCurrentOutput) GetMaxCurrent() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetCurrentOutput_getMaxCurrent(p.handle, r) })
}

// GetMinFailsafeTime returns the minimum settable failsafe time in milliseconds
func (p *PhidgetCurrentOutput) GetMinFailsafeTime() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetCurrentOutput_getMinFailsafeTime(p.handle, r) })
}

// GetMaxFailsafeTime returns the maximum settable failsafe time in milliseconds
func (p *PhidgetCurrentOutput) GetMaxFailsafeTime() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetCurrentOutput_getMaxFailsafeTime(p.handle, r) })
}

// EnableFailsafe enables the failsafe with the given time in milliseconds
func (p *PhidgetCurrentOutput) EnableFailsafe(ms uint32) error {
	return p.phidgetError(C.PhidgetCurrentOutput_enableFailsafe(p.handle, C.uint32_t(ms)))
}

// ResetFailsafe resets the failsafe state
func (p *PhidgetCurrentOutput) ResetFailsafe() error {
	return p.phidgetError(C.PhidgetCurrentOutput_resetFailsafe(p.handle))
}

// Close closes the handle and deletes it
func (p *PhidgetCurrentOutput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetCurrentOutput_delete(&p.handle))
}

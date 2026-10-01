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

// PhidgetResistanceInput wraps a Phidget resistance input
type PhidgetResistanceInput struct {
	phidget
	handle C.PhidgetResistanceInputHandle
}

// Create creates a PhidgetResistanceInput handle
func (p *PhidgetResistanceInput) Create() {
	C.PhidgetResistanceInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetResistance returns the current resistance in ohms
func (p *PhidgetResistanceInput) GetResistance() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetResistanceInput_getResistance(p.handle, r) })
}

// GetMinResistance returns the minimum measurable resistance
func (p *PhidgetResistanceInput) GetMinResistance() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetResistanceInput_getMinResistance(p.handle, r) })
}

// GetMaxResistance returns the maximum measurable resistance
func (p *PhidgetResistanceInput) GetMaxResistance() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetResistanceInput_getMaxResistance(p.handle, r) })
}

// SetResistanceChangeTrigger sets the change threshold that triggers the callback
func (p *PhidgetResistanceInput) SetResistanceChangeTrigger(trigger float64) error {
	return p.phidgetError(C.PhidgetResistanceInput_setResistanceChangeTrigger(p.handle, C.double(trigger)))
}

// GetResistanceChangeTrigger returns the current change trigger
func (p *PhidgetResistanceInput) GetResistanceChangeTrigger() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetResistanceInput_getResistanceChangeTrigger(p.handle, r)
	})
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetResistanceInput) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetResistanceInput_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetResistanceInput) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetResistanceInput_getDataInterval(p.handle, r) })
}

// SetOnResistanceChangeHandler sets a callback that fires when resistance changes beyond the trigger
func (p *PhidgetResistanceInput) SetOnResistanceChangeHandler(f func(float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetResistanceInput_setOnResistanceChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetResistanceInput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetResistanceInput_delete(&p.handle))
}

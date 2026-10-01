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

// PhidgetPressureSensor wraps a Phidget pressure sensor
type PhidgetPressureSensor struct {
	phidget
	handle C.PhidgetPressureSensorHandle
}

// Create creates a PhidgetPressureSensor handle
func (p *PhidgetPressureSensor) Create() {
	C.PhidgetPressureSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetPressure returns the current pressure in kPa
func (p *PhidgetPressureSensor) GetPressure() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPressureSensor_getPressure(p.handle, r) })
}

// GetMinPressure returns the minimum measurable pressure
func (p *PhidgetPressureSensor) GetMinPressure() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPressureSensor_getMinPressure(p.handle, r) })
}

// GetMaxPressure returns the maximum measurable pressure
func (p *PhidgetPressureSensor) GetMaxPressure() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPressureSensor_getMaxPressure(p.handle, r) })
}

// SetPressureChangeTrigger sets the change threshold that triggers the callback
func (p *PhidgetPressureSensor) SetPressureChangeTrigger(trigger float64) error {
	return p.phidgetError(C.PhidgetPressureSensor_setPressureChangeTrigger(p.handle, C.double(trigger)))
}

// GetPressureChangeTrigger returns the current change trigger
func (p *PhidgetPressureSensor) GetPressureChangeTrigger() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetPressureSensor_getPressureChangeTrigger(p.handle, r)
	})
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetPressureSensor) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetPressureSensor_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetPressureSensor) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetPressureSensor_getDataInterval(p.handle, r) })
}

// SetOnPressureChangeHandler sets a callback that fires when pressure changes beyond the trigger
func (p *PhidgetPressureSensor) SetOnPressureChangeHandler(f func(float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetPressureSensor_setOnPressureChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetPressureSensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetPressureSensor_delete(&p.handle))
}

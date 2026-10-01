package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetPHSensor wraps a Phidget pH sensor
type PhidgetPHSensor struct {
	phidget
	handle C.PhidgetPHSensorHandle
}

// Create creates a PhidgetPHSensor handle
func (p *PhidgetPHSensor) Create() {
	C.PhidgetPHSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetPH returns the current pH value (0–14)
func (p *PhidgetPHSensor) GetPH() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPHSensor_getPH(p.handle, r) })
}

// GetMinPH returns the minimum measurable pH
func (p *PhidgetPHSensor) GetMinPH() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPHSensor_getMinPH(p.handle, r) })
}

// GetMaxPH returns the maximum measurable pH
func (p *PhidgetPHSensor) GetMaxPH() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPHSensor_getMaxPH(p.handle, r) })
}

// SetPHChangeTrigger sets the change threshold that triggers the callback
func (p *PhidgetPHSensor) SetPHChangeTrigger(trigger float64) error {
	return p.phidgetError(C.PhidgetPHSensor_setPHChangeTrigger(p.handle, C.double(trigger)))
}

// GetPHChangeTrigger returns the current change trigger
func (p *PhidgetPHSensor) GetPHChangeTrigger() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPHSensor_getPHChangeTrigger(p.handle, r) })
}

// SetCorrectionTemperature sets the temperature used to compensate pH readings
func (p *PhidgetPHSensor) SetCorrectionTemperature(tempC float64) error {
	return p.phidgetError(C.PhidgetPHSensor_setCorrectionTemperature(p.handle, C.double(tempC)))
}

// GetCorrectionTemperature returns the current correction temperature
func (p *PhidgetPHSensor) GetCorrectionTemperature() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetPHSensor_getCorrectionTemperature(p.handle, r) })
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetPHSensor) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetPHSensor_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetPHSensor) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetPHSensor_getDataInterval(p.handle, r) })
}

// SetOnPHChangeHandler sets a callback that fires when pH changes beyond the trigger
func (p *PhidgetPHSensor) SetOnPHChangeHandler(f func(float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetPHSensor_setOnPHChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetPHSensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetPHSensor_delete(&p.handle))
}

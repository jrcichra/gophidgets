package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*callback_fcn)(void* handle, void* ctx, double value);
void ccallback(void* handle, void* ctx, double value);  // Forward declaration.
*/
import "C"
import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
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
	var r C.double
	if cerr := C.PhidgetPHSensor_getPH(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetMinPH returns the minimum measurable pH
func (p *PhidgetPHSensor) GetMinPH() (float64, error) {
	var r C.double
	if cerr := C.PhidgetPHSensor_getMinPH(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetMaxPH returns the maximum measurable pH
func (p *PhidgetPHSensor) GetMaxPH() (float64, error) {
	var r C.double
	if cerr := C.PhidgetPHSensor_getMaxPH(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetPHChangeTrigger sets the change threshold that triggers the callback
func (p *PhidgetPHSensor) SetPHChangeTrigger(trigger float64) error {
	return p.phidgetError(C.PhidgetPHSensor_setPHChangeTrigger(p.handle, C.double(trigger)))
}

// GetPHChangeTrigger returns the current change trigger
func (p *PhidgetPHSensor) GetPHChangeTrigger() (float64, error) {
	var r C.double
	if cerr := C.PhidgetPHSensor_getPHChangeTrigger(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetCorrectionTemperature sets the temperature used to compensate pH readings
func (p *PhidgetPHSensor) SetCorrectionTemperature(tempC float64) error {
	return p.phidgetError(C.PhidgetPHSensor_setCorrectionTemperature(p.handle, C.double(tempC)))
}

// GetCorrectionTemperature returns the current correction temperature
func (p *PhidgetPHSensor) GetCorrectionTemperature() (float64, error) {
	var r C.double
	if cerr := C.PhidgetPHSensor_getCorrectionTemperature(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetPHSensor) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetPHSensor_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetPHSensor) GetDataInterval() (uint32, error) {
	var r C.uint32_t
	if cerr := C.PhidgetPHSensor_getDataInterval(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return uint32(r), nil
}

// SetOnPHChangeHandler sets a callback that fires when pH changes beyond the trigger
func (p *PhidgetPHSensor) SetOnPHChangeHandler(f func(float64)) error {
	var pt Passthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetPHSensor_setOnPHChangeHandler(
		p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetPHSensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetPHSensor_delete(&p.handle))
}

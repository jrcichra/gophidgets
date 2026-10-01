package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetTemperatureSensor is the struct that is a phidget temperature sensor
type PhidgetTemperatureSensor struct {
	phidget
	handle C.PhidgetTemperatureSensorHandle
}

// Create creates a phidget temperature sensor
func (p *PhidgetTemperatureSensor) Create() {
	C.PhidgetTemperatureSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the temperature from a phidget temperature sensor
func (p *PhidgetTemperatureSensor) GetValue() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetTemperatureSensor_getTemperature(p.handle, r) })
}

// SetOnTemperatureChangeHandler - interrupt for temperature changes calls a function
func (p *PhidgetTemperatureSensor) SetOnTemperatureChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	pt := p.save(f)
	cerr := C.PhidgetTemperatureSensor_setOnTemperatureChangeHandler(p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), pt)
	return p.phidgetError(cerr)
}

func (p *PhidgetTemperatureSensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetTemperatureSensor_delete(&p.handle))
}

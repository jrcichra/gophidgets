package phidgets

/*
#include <phidget22.h>
*/
import "C"
import (
	"unsafe"
)

// PhidgetDigitalOutput is the struct that is a phidget digital output
type PhidgetDigitalOutput struct {
	phidget
	handle C.PhidgetDigitalOutputHandle
}

// Create creates a phidget digital output
func (p *PhidgetDigitalOutput) Create() {
	C.PhidgetDigitalOutput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the state from a phidget digital output
func (p *PhidgetDigitalOutput) GetState() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetDigitalOutput_getState(p.handle, r) })
}

// SetValue gets the state from a phidget digital output
func (p *PhidgetDigitalOutput) SetState(state bool) error {
	return p.phidgetError(C.PhidgetDigitalOutput_setState(p.handle, boolToCInt(state)))
}

// SetDutyCycle sets the duty cycle of the digital output (0.0–1.0)
func (p *PhidgetDigitalOutput) SetDutyCycle(dutyCycle float64) error {
	return p.phidgetError(C.PhidgetDigitalOutput_setDutyCycle(p.handle, C.double(dutyCycle)))
}

// GetDutyCycle gets the duty cycle of the digital output (0.0–1.0)
func (p *PhidgetDigitalOutput) GetDutyCycle() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDigitalOutput_getDutyCycle(p.handle, r) })
}

// SetLEDCurrentLimit sets the LED current limit in milliamps
func (p *PhidgetDigitalOutput) SetLEDCurrentLimit(milliamps float64) error {
	return p.phidgetError(C.PhidgetDigitalOutput_setLEDCurrentLimit(p.handle, C.double(milliamps)))
}

// GetLEDCurrentLimit gets the LED current limit in milliamps
func (p *PhidgetDigitalOutput) GetLEDCurrentLimit() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetDigitalOutput_getLEDCurrentLimit(p.handle, r) })
}

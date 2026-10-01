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

// Close - close the handle and delete it
func (p *PhidgetDigitalOutput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetDigitalOutput_delete(&p.handle))
}

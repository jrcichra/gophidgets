package phidgets

/*
#include <phidget22.h>
*/
import "C"
import (
	"unsafe"
)

// PhidgetDigitalInput is the struct that is a phidget input sensor
type PhidgetDigitalInput struct {
	phidget
	handle C.PhidgetDigitalInputHandle
}

// Create creates a phidget input sensor
func (p *PhidgetDigitalInput) Create() {
	C.PhidgetDigitalInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the input from a phidget input sensor
func (p *PhidgetDigitalInput) GetState() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetDigitalInput_getState(p.handle, r) })
}

// Close - close the handle and delete it
func (p *PhidgetDigitalInput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetDigitalInput_delete(&p.handle))
}

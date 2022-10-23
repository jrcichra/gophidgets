package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
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
	var r C.int
	cerr := C.PhidgetDigitalOutput_getState(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
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

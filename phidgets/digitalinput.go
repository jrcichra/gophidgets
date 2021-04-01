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

//PhidgetDigitalInput is the struct that is a phidget current sensor
type PhidgetDigitalInput struct {
	phidget
	handle C.PhidgetDigitalInputHandle
}

//Create creates a phidget current sensor
func (p *PhidgetDigitalInput) Create() {
	C.PhidgetDigitalInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//GetValue gets the current from a phidget current sensor
func (p *PhidgetDigitalInput) GetState() (bool, error) {
	var r C.int
	cerr := C.PhidgetDigitalInput_getState(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

//Close - close the handle and delete it
func (p *PhidgetDigitalInput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetDigitalInput_delete(&p.handle))
}

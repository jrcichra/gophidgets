package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
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

// SetOnStateChangeHandler sets a callback that is called when the
// digital input state changes. The callback receives the new state.
func (p *PhidgetDigitalInput) SetOnStateChangeHandler(f func(bool)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDigitalInput_setOnStateChangeHandler(
		p.handle, (C.phidget_state_fcn)(unsafe.Pointer(C.statecallback)), ctx))
}

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

// PhidgetCurrentInput is the struct that is a phidget current sensor
type PhidgetCurrentInput struct {
	phidget
	handle C.PhidgetCurrentInputHandle
}

// Create creates a phidget current sensor
func (p *PhidgetCurrentInput) Create() {
	C.PhidgetCurrentInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the current from a phidget current sensor
func (p *PhidgetCurrentInput) GetValue() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetCurrentInput_getCurrent(p.handle, r) })
}

// SetOnCurrentChangeHandler - interrupt for current changes calls a function
func (p *PhidgetCurrentInput) SetOnCurrentChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	pt := gopointer.Save(f)
	cerr := C.PhidgetCurrentInput_setOnCurrentChangeHandler(p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

// Close - close the handle and delete it
func (p *PhidgetCurrentInput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetCurrentInput_delete(&p.handle))
}

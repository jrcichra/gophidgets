package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*callback_fcn)(void* handle, void* ctx, double b);
void ccallback(void* handle, void* ctx, double b);  // Forward declaration.
*/
import "C"
import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

//PhidgetCurrentInput is the struct that is a phidget current sensor
type PhidgetCurrentInput struct {
	Phidget
	handle C.PhidgetCurrentInputHandle
}

//Create creates a phidget current sensor
func (p *PhidgetCurrentInput) Create() {
	C.PhidgetCurrentInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//GetValue gets the current from a phidget current sensor
func (p *PhidgetCurrentInput) GetValue() (float64, error) {
	var r C.double
	cerr := C.PhidgetCurrentInput_getCurrent(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

//SetOnCurrentChangeHandler - interrupt for current changes calls a function
func (p *PhidgetCurrentInput) SetOnCurrentChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetCurrentInput_setOnCurrentChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

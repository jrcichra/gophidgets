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

//PhidgetVoltageInput is the struct that is a phidget voltageinput sensor
type PhidgetVoltageInput struct {
	phidget
	handle C.PhidgetVoltageInputHandle
}

//Create creates a phidget voltageinput sensor
func (p *PhidgetVoltageInput) Create() {
	C.PhidgetVoltageInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//GetValue gets the voltageinput from a phidget voltageinput sensor
func (p *PhidgetVoltageInput) GetValue() (float64, error) {
	var r C.double
	cerr := C.PhidgetVoltageInput_getVoltage(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

//SetOnVoltageChangeHandler - voltage input for temperature changes calls a function
func (p *PhidgetVoltageInput) SetOnVoltageChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetVoltageInput_setOnVoltageChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

//Close - close the handle and delete it
func (p *PhidgetVoltageInput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetVoltageInput_delete(&p.handle))
}

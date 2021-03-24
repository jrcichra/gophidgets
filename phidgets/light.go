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

//PhidgetLightSensor is the struct that is a phidget lumenance sensor
type PhidgetLightSensor struct {
	Phidget
	handle C.PhidgetLightSensorHandle
}

//Create creates a phidget lumenance sensor
func (p *PhidgetLightSensor) Create() {
	C.PhidgetLightSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//GetValue gets the lumenance from a phidget lumenance sensor
func (p *PhidgetLightSensor) GetValue() (float64, error) {
	var r C.double
	cerr := C.PhidgetLightSensor_getIlluminance(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

//SetOnIlluminanceChangeHandler - interrupt for illumiance changes calls a function
func (p *PhidgetLightSensor) SetOnIlluminanceChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetLightSensor_setOnIlluminanceChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

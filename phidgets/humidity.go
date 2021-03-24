package phidgets

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L . -lphidget22
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

//PhidgetHumiditySensor is the struct that is a phidget humidity sensor
type PhidgetHumiditySensor struct {
	Phidget
	handle C.PhidgetHumiditySensorHandle
}

//Create creates a phidget humidity sensor
func (p *PhidgetHumiditySensor) Create() {
	C.PhidgetHumiditySensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//GetValue gets the humidity from a phidget humidity sensor
func (p *PhidgetHumiditySensor) GetValue() (float64, error) {
	var r C.double
	cerr := C.PhidgetHumiditySensor_getHumidity(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

//SetOnHumidityChangeHandler - interrupt for humdity changes calls a function
func (p *PhidgetHumiditySensor) SetOnHumidityChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetHumiditySensor_setOnHumidityChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

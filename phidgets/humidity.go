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

// PhidgetHumiditySensor is the struct that is a phidget humidity sensor
type PhidgetHumiditySensor struct {
	phidget
	handle C.PhidgetHumiditySensorHandle
}

// Create creates a phidget humidity sensor
func (p *PhidgetHumiditySensor) Create() {
	C.PhidgetHumiditySensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the humidity from a phidget humidity sensor
func (p *PhidgetHumiditySensor) GetValue() (float64, error) {
	var r C.double
	cerr := C.PhidgetHumiditySensor_getHumidity(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetOnHumidityChangeHandler - interrupt for humdity changes calls a function
func (p *PhidgetHumiditySensor) SetOnHumidityChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetHumiditySensor_setOnHumidityChangeHandler(p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

// Close - close the handle and delete it
func (p *PhidgetHumiditySensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetHumiditySensor_delete(&p.handle))
}

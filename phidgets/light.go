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

// PhidgetLightSensor is the struct that is a phidget lumenance sensor
type PhidgetLightSensor struct {
	phidget
	handle C.PhidgetLightSensorHandle
}

// Create creates a phidget lumenance sensor
func (p *PhidgetLightSensor) Create() {
	C.PhidgetLightSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the lumenance from a phidget lumenance sensor
func (p *PhidgetLightSensor) GetValue() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLightSensor_getIlluminance(p.handle, r) })
}

// SetOnIlluminanceChangeHandler - interrupt for illumiance changes calls a function
func (p *PhidgetLightSensor) SetOnIlluminanceChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	pt := gopointer.Save(f)
	cerr := C.PhidgetLightSensor_setOnIlluminanceChangeHandler(p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

// Close - close the handle and delete it
func (p *PhidgetLightSensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetLightSensor_delete(&p.handle))
}

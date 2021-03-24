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

//PhidgetTemperatureSensor is the struct that is a phidget temperature sensor
type PhidgetTemperatureSensor struct {
	Phidget
	handle C.PhidgetTemperatureSensorHandle
}

//Create creates a phidget temperature sensor
func (p *PhidgetTemperatureSensor) Create() {
	C.PhidgetTemperatureSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//GetValue gets the temperature from a phidget temperature sensor
func (p *PhidgetTemperatureSensor) GetValue() (float64, error) {
	var r C.double
	if cerr := C.PhidgetTemperatureSensor_getTemperature(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

//SetOnTemperatureChangeHandler - interrupt for temperature changes calls a function
func (p *PhidgetTemperatureSensor) SetOnTemperatureChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetTemperatureSensor_setOnTemperatureChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	return p.phidgetError(cerr)
}

func (p *PhidgetTemperatureSensor) Close() error {
	if err := p.Phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetTemperatureSensor_delete(&p.handle))
}

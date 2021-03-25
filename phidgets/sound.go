package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*sound_callback_fcn)(void* handle, void* ctx, double dB, double dBA, double dBC, const double octaves[10]);
void csoundcallback(void* handle, void* ctx, double dB, double dBA, double dBC, const double octaves[10]);  // Forward declaration.
*/
import "C"
import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

//PhidgetSoundSensor is the struct that is a phidget sound sensor
type PhidgetSoundSensor struct {
	phidget
	handle C.PhidgetSoundSensorHandle
}

//Create creates a phidget sound sensor
func (p *PhidgetSoundSensor) Create() {
	C.PhidgetSoundSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//GetValue gets the decibels from a phidget sound sensor
func (p *PhidgetSoundSensor) GetValue() (float64, error) {
	var r C.double
	cerr := C.PhidgetSoundSensor_getdB(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

//SetSPLChangeTrigger sets the interrupt trigger point
func (p *PhidgetSoundSensor) SetSPLChangeTrigger(dBs float64) error {
	return p.phidgetError(C.PhidgetSoundSensor_setSPLChangeTrigger(p.handle, C.double(dBs)))
}

//SetOnSPLChangeHandler - interrupt for sound changes calls a function
func (p *PhidgetSoundSensor) SetOnSPLChangeHandler(f func(float64, float64, float64, []float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough SoundPassthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetSoundSensor_setOnSPLChangeHandler(p.handle, (C.sound_callback_fcn)(unsafe.Pointer(C.csoundcallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

//Close - close the handle and delete it
func (p *PhidgetSoundSensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetSoundSensor_delete(&p.handle))
}

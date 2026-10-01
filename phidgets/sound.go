package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetSoundSensor is the struct that is a phidget sound sensor
type PhidgetSoundSensor struct {
	phidget
	handle C.PhidgetSoundSensorHandle
}

// Create creates a phidget sound sensor
func (p *PhidgetSoundSensor) Create() {
	C.PhidgetSoundSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the decibels from a phidget sound sensor
func (p *PhidgetSoundSensor) GetValue() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetSoundSensor_getdB(p.handle, r) })
}

// SetSPLChangeTrigger sets the interrupt trigger point
func (p *PhidgetSoundSensor) SetSPLChangeTrigger(dBs float64) error {
	return p.phidgetError(C.PhidgetSoundSensor_setSPLChangeTrigger(p.handle, C.double(dBs)))
}

// SetOnSPLChangeHandler - interrupt for sound changes calls a function
func (p *PhidgetSoundSensor) SetOnSPLChangeHandler(f func(float64, float64, float64, []float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	pt := p.save(f)
	cerr := C.PhidgetSoundSensor_setOnSPLChangeHandler(p.handle, (C.phidget_sound_fcn)(unsafe.Pointer(C.csoundcallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

// Close - close the handle and delete it
func (p *PhidgetSoundSensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetSoundSensor_delete(&p.handle))
}

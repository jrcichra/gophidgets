package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetAccelerometer is the struct that is a phidget motion sensor
type PhidgetAccelerometer struct {
	phidget
	handle C.PhidgetAccelerometerHandle
}

// Create creates a phidget motion sensor
func (p *PhidgetAccelerometer) Create() {
	C.PhidgetAccelerometer_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetAcceleration gets the acceleration from a phidget motion sensor
func (p *PhidgetAccelerometer) GetAcceleration() ([]float64, error) {
	var r [3]C.double
	cerr := C.PhidgetAccelerometer_getAcceleration(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return []float64{(float64)(r[0]), (float64)(r[1]), (float64)(r[2])}, nil
}

// GetMinAcceleration gets the min acceleration value from a phidget motion sensor
func (p *PhidgetAccelerometer) GetMinAcceleration() ([]float64, error) {
	var r [3]C.double
	cerr := C.PhidgetAccelerometer_getMinAcceleration(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return []float64{(float64)(r[0]), (float64)(r[1]), (float64)(r[2])}, nil
}

// GetMaxAcceleration gets the max acceleration value from a phidget motion sensor
func (p *PhidgetAccelerometer) GetMaxAcceleration() ([]float64, error) {
	var r [3]C.double
	cerr := C.PhidgetAccelerometer_getMaxAcceleration(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return []float64{(float64)(r[0]), (float64)(r[1]), (float64)(r[2])}, nil
}

// GetAccelerationChangeTrigger gets the acceleration from a phidget temperature sensor
func (p *PhidgetAccelerometer) GetAccelerationChangeTrigger() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetAccelerometer_getAccelerationChangeTrigger(p.handle, r)
	})
}

// SetAccelerationChangeTrigger sets the acceleration trigger in the phidget temperature sensor
func (p *PhidgetAccelerometer) SetAccelerationChangeTrigger(value float64) error {
	return p.phidgetError(C.PhidgetAccelerometer_setAccelerationChangeTrigger(p.handle, C.double(value)))
}

// GetMinAccelerationChangeTrigger sets the min acceleration trigger in the phidget temperature sensor
func (p *PhidgetAccelerometer) GetMinAccelerationChangeTrigger() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetAccelerometer_getMinAccelerationChangeTrigger(p.handle, r)
	})
}

// GetMaxAccelerationChangeTrigger sets the min acceleration trigger in the phidget temperature sensor
func (p *PhidgetAccelerometer) GetMaxAccelerationChangeTrigger() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetAccelerometer_getMaxAccelerationChangeTrigger(p.handle, r)
	})
}

// GetAxisCount return the number of axis of the motion sensor
func (p *PhidgetAccelerometer) GetAxisCount() (int, error) {
	return getInt(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetAccelerometer_getAxisCount(p.handle, r) })
}

// GetDataInterval return the number of axis of the motion sensor
func (p *PhidgetAccelerometer) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetAccelerometer_getDataInterval(p.handle, r) })
}

// SetDataInterval sets the interval between OnAccelerationChange callback calls
func (p *PhidgetAccelerometer) SetDataInterval(value uint32) error {
	cerr := C.PhidgetAccelerometer_setDataInterval(p.handle, (C.uint32_t)(value))
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

// GetMinDataInterval return the number of axis of the motion sensor
func (p *PhidgetAccelerometer) GetMinDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetAccelerometer_getMinDataInterval(p.handle, r) })
}

// GetMaxDataInterval return the number of axis of the motion sensor
func (p *PhidgetAccelerometer) GetMaxDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetAccelerometer_getMaxDataInterval(p.handle, r) })
}

// SetOnAccelerationChangeHandler - interrupt for motion changes calls a function
func (p *PhidgetAccelerometer) SetOnAccelerationChangeHandler(f func([]float64, float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	pt := p.save(f)
	cerr := C.PhidgetAccelerometer_setOnAccelerationChangeHandler(p.handle, (C.phidget_motion_fcn)(unsafe.Pointer(C.motioncallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

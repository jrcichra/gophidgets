package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetSpatial wraps a Phidget spatial sensor (combined accel + gyro + magnetometer)
type PhidgetSpatial struct {
	phidget
	handle C.PhidgetSpatialHandle
}

// Create creates a PhidgetSpatial handle
func (p *PhidgetSpatial) Create() {
	C.PhidgetSpatial_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetEulerAngles returns the current orientation as Euler angles (roll, pitch, heading) in degrees
func (p *PhidgetSpatial) GetEulerAngles() (roll, pitch, heading float64, err error) {
	var angles C.PhidgetSpatial_SpatialEulerAngles
	if cerr := C.PhidgetSpatial_getEulerAngles(p.handle, &angles); cerr != C.EPHIDGET_OK {
		return 0, 0, 0, p.phidgetError(cerr)
	}
	return float64(angles.roll), float64(angles.pitch), float64(angles.heading), nil
}

// GetQuaternion returns the current orientation as a quaternion [w, x, y, z]
func (p *PhidgetSpatial) GetQuaternion() ([]float64, error) {
	var q C.PhidgetSpatial_SpatialQuaternion
	if cerr := C.PhidgetSpatial_getQuaternion(p.handle, &q); cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return []float64{float64(q.w), float64(q.x), float64(q.y), float64(q.z)}, nil
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetSpatial) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetSpatial_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetSpatial) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetSpatial_getDataInterval(p.handle, r) })
}

// SetHeatingEnabled enables or disables the internal heater for temperature stability
func (p *PhidgetSpatial) SetHeatingEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetSpatial_setHeatingEnabled(p.handle, boolToCInt(enabled)))
}

// GetHeatingEnabled returns whether the internal heater is enabled
func (p *PhidgetSpatial) GetHeatingEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetSpatial_getHeatingEnabled(p.handle, r) })
}

// ZeroGyro zeroes the gyroscope. The sensor must be stationary.
func (p *PhidgetSpatial) ZeroGyro() error {
	return p.phidgetError(C.PhidgetSpatial_zeroGyro(p.handle))
}

// ZeroAlgorithm resets the AHRS/orientation algorithm
func (p *PhidgetSpatial) ZeroAlgorithm() error {
	return p.phidgetError(C.PhidgetSpatial_zeroAlgorithm(p.handle))
}

// SetAlgorithm sets the orientation algorithm (use the SpatialAlgorithm constants)
func (p *PhidgetSpatial) SetAlgorithm(algo int) error {
	return p.phidgetError(C.PhidgetSpatial_setAlgorithm(p.handle, C.Phidget_SpatialAlgorithm(algo)))
}

// SetOnSpatialDataHandler sets a callback that fires on each spatial data update.
// The callback receives: acceleration [x,y,z], angularRate [x,y,z], magneticField [x,y,z], timestamp.
func (p *PhidgetSpatial) SetOnSpatialDataHandler(f func([]float64, []float64, []float64, float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetSpatial_setOnSpatialDataHandler(
		p.handle, (C.phidget_spatial_fcn)(unsafe.Pointer(C.spatialcallback)), ctx))
}

// SetOnAlgorithmDataHandler sets a callback that fires on each orientation algorithm update.
// The callback receives: quaternion [w,x,y,z], timestamp.
func (p *PhidgetSpatial) SetOnAlgorithmDataHandler(f func([]float64, float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetSpatial_setOnAlgorithmDataHandler(
		p.handle, (C.phidget_motion_fcn)(unsafe.Pointer(C.quaternioncallback)), ctx))
}

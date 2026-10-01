package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetGyroscope wraps a Phidget gyroscope (angular rate sensor)
type PhidgetGyroscope struct {
	phidget
	handle C.PhidgetGyroscopeHandle
}

// Create creates a PhidgetGyroscope handle
func (p *PhidgetGyroscope) Create() {
	C.PhidgetGyroscope_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetAngularRate returns the current angular rate in degrees/second as [x, y, z]
func (p *PhidgetGyroscope) GetAngularRate() ([]float64, error) {
	return getVec3(&p.phidget, func(r *[3]C.double) C.PhidgetReturnCode { return C.PhidgetGyroscope_getAngularRate(p.handle, r) })
}

// GetMinAngularRate returns the minimum measurable angular rate as [x, y, z]
func (p *PhidgetGyroscope) GetMinAngularRate() ([]float64, error) {
	return getVec3(&p.phidget, func(r *[3]C.double) C.PhidgetReturnCode { return C.PhidgetGyroscope_getMinAngularRate(p.handle, r) })
}

// GetMaxAngularRate returns the maximum measurable angular rate as [x, y, z]
func (p *PhidgetGyroscope) GetMaxAngularRate() ([]float64, error) {
	return getVec3(&p.phidget, func(r *[3]C.double) C.PhidgetReturnCode { return C.PhidgetGyroscope_getMaxAngularRate(p.handle, r) })
}

// GetAxisCount returns the number of axes
func (p *PhidgetGyroscope) GetAxisCount() (int, error) {
	return getInt(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetGyroscope_getAxisCount(p.handle, r) })
}

// GetTimestamp returns the timestamp of the most recent angular rate reading
func (p *PhidgetGyroscope) GetTimestamp() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetGyroscope_getTimestamp(p.handle, r) })
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetGyroscope) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetGyroscope_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetGyroscope) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetGyroscope_getDataInterval(p.handle, r) })
}

// SetHeatingEnabled enables or disables the internal heater for temperature stability
func (p *PhidgetGyroscope) SetHeatingEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetGyroscope_setHeatingEnabled(p.handle, boolToCInt(enabled)))
}

// GetHeatingEnabled returns whether the internal heater is enabled
func (p *PhidgetGyroscope) GetHeatingEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetGyroscope_getHeatingEnabled(p.handle, r) })
}

// Zero zeroes the gyroscope output. The sensor must be stationary for accurate calibration.
func (p *PhidgetGyroscope) Zero() error {
	return p.phidgetError(C.PhidgetGyroscope_zero(p.handle))
}

// SetOnAngularRateUpdateHandler sets a callback that fires on each angular rate update.
// The callback receives the angular rate as [x, y, z] and the timestamp in seconds.
func (p *PhidgetGyroscope) SetOnAngularRateUpdateHandler(f func([]float64, float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetGyroscope_setOnAngularRateUpdateHandler(
		p.handle, (C.phidget_motion_fcn)(unsafe.Pointer(C.cmotioncallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetGyroscope) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetGyroscope_delete(&p.handle))
}

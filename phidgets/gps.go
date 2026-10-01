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

// PhidgetGPS wraps a Phidget GPS receiver
type PhidgetGPS struct {
	phidget
	handle C.PhidgetGPSHandle
}

// Create creates a PhidgetGPS handle
func (p *PhidgetGPS) Create() {
	C.PhidgetGPS_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetLatitude returns the current latitude in degrees
func (p *PhidgetGPS) GetLatitude() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetGPS_getLatitude(p.handle, r) })
}

// GetLongitude returns the current longitude in degrees
func (p *PhidgetGPS) GetLongitude() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetGPS_getLongitude(p.handle, r) })
}

// GetAltitude returns the current altitude in meters
func (p *PhidgetGPS) GetAltitude() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetGPS_getAltitude(p.handle, r) })
}

// GetHeading returns the current heading in degrees (0–360)
func (p *PhidgetGPS) GetHeading() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetGPS_getHeading(p.handle, r) })
}

// GetVelocity returns the current ground speed in km/h
func (p *PhidgetGPS) GetVelocity() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetGPS_getVelocity(p.handle, r) })
}

// GetPositionFixState returns whether the GPS has a position fix
func (p *PhidgetGPS) GetPositionFixState() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetGPS_getPositionFixState(p.handle, r) })
}

// SetOnPositionChangeHandler sets a callback that fires when the GPS position changes.
// The callback receives latitude, longitude, and altitude.
func (p *PhidgetGPS) SetOnPositionChangeHandler(f func(float64, float64, float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetGPS_setOnPositionChangeHandler(
		p.handle, (C.phidget_threefloat_fcn)(unsafe.Pointer(C.cthreefloatcallback)), ctx))
}

// SetOnHeadingChangeHandler sets a callback that fires when the heading or velocity changes.
// The callback receives heading (degrees) and velocity (km/h).
func (p *PhidgetGPS) SetOnHeadingChangeHandler(f func(float64, float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetGPS_setOnHeadingChangeHandler(
		p.handle, (C.phidget_twofloat_fcn)(unsafe.Pointer(C.ctwofloatcallback)), ctx))
}

// SetOnPositionFixStateChangeHandler sets a callback that fires when the fix state changes.
// The callback receives 1 for fix acquired, 0 for fix lost.
func (p *PhidgetGPS) SetOnPositionFixStateChangeHandler(f func(float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetGPS_setOnPositionFixStateChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetGPS) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetGPS_delete(&p.handle))
}

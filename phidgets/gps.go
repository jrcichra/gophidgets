package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*callback_fcn)(void* handle, void* ctx, double value);
void ccallback(void* handle, void* ctx, double value);  // Forward declaration.
typedef void (*twofloat_callback_fcn)(void* handle, void* ctx, double a, double b);
void ctwofloatcallback(void* handle, void* ctx, double a, double b);  // Forward declaration.
typedef void (*threefloat_callback_fcn)(void* handle, void* ctx, double a, double b, double c);
void cthreefloatcallback(void* handle, void* ctx, double a, double b, double c);  // Forward declaration.
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
	var r C.double
	if cerr := C.PhidgetGPS_getLatitude(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetLongitude returns the current longitude in degrees
func (p *PhidgetGPS) GetLongitude() (float64, error) {
	var r C.double
	if cerr := C.PhidgetGPS_getLongitude(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetAltitude returns the current altitude in meters
func (p *PhidgetGPS) GetAltitude() (float64, error) {
	var r C.double
	if cerr := C.PhidgetGPS_getAltitude(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetHeading returns the current heading in degrees (0–360)
func (p *PhidgetGPS) GetHeading() (float64, error) {
	var r C.double
	if cerr := C.PhidgetGPS_getHeading(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetVelocity returns the current ground speed in km/h
func (p *PhidgetGPS) GetVelocity() (float64, error) {
	var r C.double
	if cerr := C.PhidgetGPS_getVelocity(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetPositionFixState returns whether the GPS has a position fix
func (p *PhidgetGPS) GetPositionFixState() (bool, error) {
	var r C.int
	if cerr := C.PhidgetGPS_getPositionFixState(p.handle, &r); cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

// SetOnPositionChangeHandler sets a callback that fires when the GPS position changes.
// The callback receives latitude, longitude, and altitude.
func (p *PhidgetGPS) SetOnPositionChangeHandler(f func(float64, float64, float64)) error {
	var pt ThreeFloatPassthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetGPS_setOnPositionChangeHandler(
		p.handle, (C.threefloat_callback_fcn)(unsafe.Pointer(C.cthreefloatcallback)), ctx))
}

// SetOnHeadingChangeHandler sets a callback that fires when the heading or velocity changes.
// The callback receives heading (degrees) and velocity (km/h).
func (p *PhidgetGPS) SetOnHeadingChangeHandler(f func(float64, float64)) error {
	var pt TwoFloatPassthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetGPS_setOnHeadingChangeHandler(
		p.handle, (C.twofloat_callback_fcn)(unsafe.Pointer(C.ctwofloatcallback)), ctx))
}

// SetOnPositionFixStateChangeHandler sets a callback that fires when the fix state changes.
// The callback receives 1 for fix acquired, 0 for fix lost.
func (p *PhidgetGPS) SetOnPositionFixStateChangeHandler(f func(float64)) error {
	var pt Passthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetGPS_setOnPositionFixStateChangeHandler(
		p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetGPS) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetGPS_delete(&p.handle))
}

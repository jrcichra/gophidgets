package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*callback_fcn)(void* handle, void* ctx, const double magneticField[3], double timestamp);
void cmotioncallback(void* handle, void* ctx, const double magneticField[3], double timestamp);  // Forward declaration.
*/
import "C"
import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

// PhidgetMagnetometer wraps a Phidget magnetometer (compass)
type PhidgetMagnetometer struct {
	phidget
	handle C.PhidgetMagnetometerHandle
}

// Create creates a PhidgetMagnetometer handle
func (p *PhidgetMagnetometer) Create() {
	C.PhidgetMagnetometer_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetMagneticField returns the current magnetic field in Gauss as [x, y, z]
func (p *PhidgetMagnetometer) GetMagneticField() ([]float64, error) {
	var r [3]C.double
	if cerr := C.PhidgetMagnetometer_getMagneticField(p.handle, &r); cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return []float64{float64(r[0]), float64(r[1]), float64(r[2])}, nil
}

// GetMinMagneticField returns the minimum measurable magnetic field as [x, y, z]
func (p *PhidgetMagnetometer) GetMinMagneticField() ([]float64, error) {
	var r [3]C.double
	if cerr := C.PhidgetMagnetometer_getMinMagneticField(p.handle, &r); cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return []float64{float64(r[0]), float64(r[1]), float64(r[2])}, nil
}

// GetMaxMagneticField returns the maximum measurable magnetic field as [x, y, z]
func (p *PhidgetMagnetometer) GetMaxMagneticField() ([]float64, error) {
	var r [3]C.double
	if cerr := C.PhidgetMagnetometer_getMaxMagneticField(p.handle, &r); cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return []float64{float64(r[0]), float64(r[1]), float64(r[2])}, nil
}

// GetAxisCount returns the number of axes
func (p *PhidgetMagnetometer) GetAxisCount() (int, error) {
	var r C.int
	if cerr := C.PhidgetMagnetometer_getAxisCount(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// GetTimestamp returns the timestamp of the most recent reading
func (p *PhidgetMagnetometer) GetTimestamp() (float64, error) {
	var r C.double
	if cerr := C.PhidgetMagnetometer_getTimestamp(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetMagneticFieldChangeTrigger sets the change threshold for callbacks
func (p *PhidgetMagnetometer) SetMagneticFieldChangeTrigger(trigger float64) error {
	return p.phidgetError(C.PhidgetMagnetometer_setMagneticFieldChangeTrigger(p.handle, C.double(trigger)))
}

// GetMagneticFieldChangeTrigger returns the current change trigger
func (p *PhidgetMagnetometer) GetMagneticFieldChangeTrigger() (float64, error) {
	var r C.double
	if cerr := C.PhidgetMagnetometer_getMagneticFieldChangeTrigger(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetMagnetometer) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetMagnetometer_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetMagnetometer) GetDataInterval() (uint32, error) {
	var r C.uint32_t
	if cerr := C.PhidgetMagnetometer_getDataInterval(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return uint32(r), nil
}

// SetHeatingEnabled enables or disables the internal heater for temperature stability
func (p *PhidgetMagnetometer) SetHeatingEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetMagnetometer_setHeatingEnabled(p.handle, boolToCInt(enabled)))
}

// GetHeatingEnabled returns whether the internal heater is enabled
func (p *PhidgetMagnetometer) GetHeatingEnabled() (bool, error) {
	var r C.int
	if cerr := C.PhidgetMagnetometer_getHeatingEnabled(p.handle, &r); cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

// SetOnMagneticFieldChangeHandler sets a callback that fires when the magnetic field changes.
// The callback receives the field as [x, y, z] and the timestamp in seconds.
func (p *PhidgetMagnetometer) SetOnMagneticFieldChangeHandler(f func([]float64, float64)) error {
	var pt MotionPassthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetMagnetometer_setOnMagneticFieldChangeHandler(
		p.handle, (C.callback_fcn)(unsafe.Pointer(C.cmotioncallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetMagnetometer) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetMagnetometer_delete(&p.handle))
}

package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
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
	return getVec3(&p.phidget, func(r *[3]C.double) C.PhidgetReturnCode { return C.PhidgetMagnetometer_getMagneticField(p.handle, r) })
}

// GetMinMagneticField returns the minimum measurable magnetic field as [x, y, z]
func (p *PhidgetMagnetometer) GetMinMagneticField() ([]float64, error) {
	return getVec3(&p.phidget, func(r *[3]C.double) C.PhidgetReturnCode {
		return C.PhidgetMagnetometer_getMinMagneticField(p.handle, r)
	})
}

// GetMaxMagneticField returns the maximum measurable magnetic field as [x, y, z]
func (p *PhidgetMagnetometer) GetMaxMagneticField() ([]float64, error) {
	return getVec3(&p.phidget, func(r *[3]C.double) C.PhidgetReturnCode {
		return C.PhidgetMagnetometer_getMaxMagneticField(p.handle, r)
	})
}

// GetAxisCount returns the number of axes
func (p *PhidgetMagnetometer) GetAxisCount() (int, error) {
	return getInt(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetMagnetometer_getAxisCount(p.handle, r) })
}

// GetTimestamp returns the timestamp of the most recent reading
func (p *PhidgetMagnetometer) GetTimestamp() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetMagnetometer_getTimestamp(p.handle, r) })
}

// SetMagneticFieldChangeTrigger sets the change threshold for callbacks
func (p *PhidgetMagnetometer) SetMagneticFieldChangeTrigger(trigger float64) error {
	return p.phidgetError(C.PhidgetMagnetometer_setMagneticFieldChangeTrigger(p.handle, C.double(trigger)))
}

// GetMagneticFieldChangeTrigger returns the current change trigger
func (p *PhidgetMagnetometer) GetMagneticFieldChangeTrigger() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetMagnetometer_getMagneticFieldChangeTrigger(p.handle, r)
	})
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetMagnetometer) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetMagnetometer_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetMagnetometer) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetMagnetometer_getDataInterval(p.handle, r) })
}

// SetHeatingEnabled enables or disables the internal heater for temperature stability
func (p *PhidgetMagnetometer) SetHeatingEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetMagnetometer_setHeatingEnabled(p.handle, boolToCInt(enabled)))
}

// GetHeatingEnabled returns whether the internal heater is enabled
func (p *PhidgetMagnetometer) GetHeatingEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetMagnetometer_getHeatingEnabled(p.handle, r) })
}

// SetOnMagneticFieldChangeHandler sets a callback that fires when the magnetic field changes.
// The callback receives the field as [x, y, z] and the timestamp in seconds.
func (p *PhidgetMagnetometer) SetOnMagneticFieldChangeHandler(f func([]float64, float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetMagnetometer_setOnMagneticFieldChangeHandler(
		p.handle, (C.phidget_motion_fcn)(unsafe.Pointer(C.motioncallback)), ctx))
}

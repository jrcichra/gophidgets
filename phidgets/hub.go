package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetHub wraps a Phidget Vint Hub
type PhidgetHub struct {
	phidget
	handle C.PhidgetHubHandle
}

// Create creates a PhidgetHub handle
func (p *PhidgetHub) Create() {
	C.PhidgetHub_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetPortMaxSpeed returns the maximum speed for the given port
func (p *PhidgetHub) GetPortMaxSpeed(port int) (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetHub_getPortMaxSpeed(p.handle, C.int(port), r)
	})
}

// GetPortMode returns the mode of the given port (PhidgetHub_PortMode)
func (p *PhidgetHub) GetPortMode(port int) (HubPortMode, error) {
	var r C.PhidgetHub_PortMode
	if cerr := C.PhidgetHub_getPortMode(p.handle, C.int(port), &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return HubPortMode(r), nil
}

// SetPortMode sets the mode of the given port (PhidgetHub_PortMode)
func (p *PhidgetHub) SetPortMode(port int, mode HubPortMode) error {
	return p.phidgetError(C.PhidgetHub_setPortMode(p.handle, C.int(port), C.PhidgetHub_PortMode(mode)))
}

// GetPortPower returns whether power is enabled on the given port
func (p *PhidgetHub) GetPortPower(port int) (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode {
		return C.PhidgetHub_getPortPower(p.handle, C.int(port), r)
	})
}

// SetPortPower enables or disables power on the given port
func (p *PhidgetHub) SetPortPower(port int, enabled bool) error {
	return p.phidgetError(C.PhidgetHub_setPortPower(p.handle, C.int(port), boolToCInt(enabled)))
}

// GetPortSupportsAutoSetSpeed returns whether the port supports automatic speed setting
func (p *PhidgetHub) GetPortSupportsAutoSetSpeed(port int) (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode {
		return C.PhidgetHub_getPortSupportsAutoSetSpeed(p.handle, C.int(port), r)
	})
}

// GetPortSupportsSetSpeed returns whether the port supports speed setting
func (p *PhidgetHub) GetPortSupportsSetSpeed(port int) (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode {
		return C.PhidgetHub_getPortSupportsSetSpeed(p.handle, C.int(port), r)
	})
}

// SetPortAutoSetSpeed enables or disables automatic speed setting on the port
func (p *PhidgetHub) SetPortAutoSetSpeed(port int, enabled bool) error {
	return p.phidgetError(C.PhidgetHub_setPortAutoSetSpeed(p.handle, C.int(port), boolToCInt(enabled)))
}

// SetFirmwareUpgradeFlag sets the firmware upgrade flag on the given port
func (p *PhidgetHub) SetFirmwareUpgradeFlag(port int, timeoutMS uint32) error {
	return p.phidgetError(C.PhidgetHub_setFirmwareUpgradeFlag(p.handle, C.int(port), C.uint32_t(timeoutMS)))
}

// SetADCCalibrationValues sets the ADC calibration values for the hub's
// voltage input and voltage ratio input ports (six values each).
func (p *PhidgetHub) SetADCCalibrationValues(voltageInputGain, voltageRatioGain [6]float64) error {
	var vi [6]C.double
	var vr [6]C.double
	for i, v := range voltageInputGain {
		vi[i] = C.double(v)
	}
	for i, v := range voltageRatioGain {
		vr[i] = C.double(v)
	}
	return p.phidgetError(C.PhidgetHub_setADCCalibrationValues(p.handle, &vi[0], &vr[0]))
}

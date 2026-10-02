package phidgets

/*
#include <phidget22.h>
#include <stdlib.h>
#include "phidgets.h"
*/
import "C"
import (
	"unsafe"
)

// PhidgetFirmwareUpgrade wraps a Phidget Firmware Upgrade channel
type PhidgetFirmwareUpgrade struct {
	phidget
	handle C.PhidgetFirmwareUpgradeHandle
}

// Create creates a PhidgetFirmwareUpgrade handle
func (p *PhidgetFirmwareUpgrade) Create() {
	C.PhidgetFirmwareUpgrade_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetActualDeviceID returns the device ID of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceID() (int, error) {
	return get(&p.phidget, func(r *C.Phidget_DeviceID) C.PhidgetReturnCode {
		return C.PhidgetFirmwareUpgrade_getActualDeviceID(p.handle, r)
	}, func(r C.Phidget_DeviceID) int { return int(r) })
}

// GetActualDeviceName returns the name of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceName() (string, error) {
	return getString(&p.phidget, func(r **C.char) C.PhidgetReturnCode { return C.PhidgetFirmwareUpgrade_getActualDeviceName(p.handle, r) })
}

// GetActualDeviceSKU returns the SKU of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceSKU() (string, error) {
	return getString(&p.phidget, func(r **C.char) C.PhidgetReturnCode { return C.PhidgetFirmwareUpgrade_getActualDeviceSKU(p.handle, r) })
}

// GetActualDeviceVersion returns the firmware version of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceVersion() (int, error) {
	return getInt(&p.phidget, func(r *C.int) C.PhidgetReturnCode {
		return C.PhidgetFirmwareUpgrade_getActualDeviceVersion(p.handle, r)
	})
}

// GetActualDeviceVINTID returns the VINT ID of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceVINTID() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetFirmwareUpgrade_getActualDeviceVINTID(p.handle, r)
	})
}

// GetProgress returns the firmware upgrade progress as a fraction (0.0–1.0)
func (p *PhidgetFirmwareUpgrade) GetProgress() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetFirmwareUpgrade_getProgress(p.handle, r) })
}

// SendFirmware sends the firmware image to the device and begins the upgrade
func (p *PhidgetFirmwareUpgrade) SendFirmware(data []byte) error {
	if len(data) == 0 {
		return p.phidgetError(C.EPHIDGET_INVALIDARG)
	}
	ptr := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	return p.phidgetError(C.PhidgetFirmwareUpgrade_sendFirmware(p.handle, ptr, C.size_t(len(data))))
}

// SetOnProgressChangeHandler sets a callback that fires as the upgrade progresses
func (p *PhidgetFirmwareUpgrade) SetOnProgressChangeHandler(f func(progress float64)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetFirmwareUpgrade_setOnProgressChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.callback)), ctx))
}

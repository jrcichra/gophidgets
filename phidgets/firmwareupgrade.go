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
	var r C.Phidget_DeviceID
	if cerr := C.PhidgetFirmwareUpgrade_getActualDeviceID(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// GetActualDeviceName returns the name of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceName() (string, error) {
	var cstr *C.char
	if cerr := C.PhidgetFirmwareUpgrade_getActualDeviceName(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		return "", p.phidgetError(cerr)
	}
	return C.GoString(cstr), nil
}

// GetActualDeviceSKU returns the SKU of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceSKU() (string, error) {
	var cstr *C.char
	if cerr := C.PhidgetFirmwareUpgrade_getActualDeviceSKU(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		return "", p.phidgetError(cerr)
	}
	return C.GoString(cstr), nil
}

// GetActualDeviceVersion returns the firmware version of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceVersion() (int, error) {
	var r C.int
	if cerr := C.PhidgetFirmwareUpgrade_getActualDeviceVersion(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// GetActualDeviceVINTID returns the VINT ID of the device behind the upgrade channel
func (p *PhidgetFirmwareUpgrade) GetActualDeviceVINTID() (uint32, error) {
	var r C.uint32_t
	if cerr := C.PhidgetFirmwareUpgrade_getActualDeviceVINTID(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return uint32(r), nil
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
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetFirmwareUpgrade) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetFirmwareUpgrade_delete(&p.handle))
}

package phidgets

/*
#include <phidget22.h>
*/
import "C"
import (
	"unsafe"
)

// PhidgetVoltageOutput wraps a Phidget voltage output channel
type PhidgetVoltageOutput struct {
	phidget
	handle C.PhidgetVoltageOutputHandle
}

// Create creates a PhidgetVoltageOutput handle
func (p *PhidgetVoltageOutput) Create() {
	C.PhidgetVoltageOutput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// SetVoltage sets the output voltage
func (p *PhidgetVoltageOutput) SetVoltage(voltage float64) error {
	return p.phidgetError(C.PhidgetVoltageOutput_setVoltage(p.handle, C.double(voltage)))
}

// GetVoltage returns the current output voltage setting
func (p *PhidgetVoltageOutput) GetVoltage() (float64, error) {
	var r C.double
	if cerr := C.PhidgetVoltageOutput_getVoltage(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetMinVoltage returns the minimum settable voltage
func (p *PhidgetVoltageOutput) GetMinVoltage() (float64, error) {
	var r C.double
	if cerr := C.PhidgetVoltageOutput_getMinVoltage(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// GetMaxVoltage returns the maximum settable voltage
func (p *PhidgetVoltageOutput) GetMaxVoltage() (float64, error) {
	var r C.double
	if cerr := C.PhidgetVoltageOutput_getMaxVoltage(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetEnabled enables or disables the voltage output
func (p *PhidgetVoltageOutput) SetEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetVoltageOutput_setEnabled(p.handle, boolToCInt(enabled)))
}

// GetEnabled returns whether the voltage output is enabled
func (p *PhidgetVoltageOutput) GetEnabled() (bool, error) {
	var r C.int
	if cerr := C.PhidgetVoltageOutput_getEnabled(p.handle, &r); cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

// SetVoltageOutputRange sets the output voltage range (use PhidgetVoltageOutput_OutputRange constants)
func (p *PhidgetVoltageOutput) SetVoltageOutputRange(outputRange C.PhidgetVoltageOutput_VoltageOutputRange) error {
	return p.phidgetError(C.PhidgetVoltageOutput_setVoltageOutputRange(p.handle, outputRange))
}

// GetVoltageOutputRange returns the current output voltage range
func (p *PhidgetVoltageOutput) GetVoltageOutputRange() (C.PhidgetVoltageOutput_VoltageOutputRange, error) {
	var r C.PhidgetVoltageOutput_VoltageOutputRange
	if cerr := C.PhidgetVoltageOutput_getVoltageOutputRange(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return r, nil
}

// Close closes the handle and deletes it
func (p *PhidgetVoltageOutput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetVoltageOutput_delete(&p.handle))
}

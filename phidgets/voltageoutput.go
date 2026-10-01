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
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetVoltageOutput_getVoltage(p.handle, r) })
}

// GetMinVoltage returns the minimum settable voltage
func (p *PhidgetVoltageOutput) GetMinVoltage() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetVoltageOutput_getMinVoltage(p.handle, r) })
}

// GetMaxVoltage returns the maximum settable voltage
func (p *PhidgetVoltageOutput) GetMaxVoltage() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetVoltageOutput_getMaxVoltage(p.handle, r) })
}

// SetEnabled enables or disables the voltage output
func (p *PhidgetVoltageOutput) SetEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetVoltageOutput_setEnabled(p.handle, boolToCInt(enabled)))
}

// GetEnabled returns whether the voltage output is enabled
func (p *PhidgetVoltageOutput) GetEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetVoltageOutput_getEnabled(p.handle, r) })
}

// SetVoltageOutputRange sets the output voltage range (use the VoltageOutputRange constants)
func (p *PhidgetVoltageOutput) SetVoltageOutputRange(outputRange int) error {
	return p.phidgetError(C.PhidgetVoltageOutput_setVoltageOutputRange(p.handle, C.PhidgetVoltageOutput_VoltageOutputRange(outputRange)))
}

// GetVoltageOutputRange returns the current output voltage range
func (p *PhidgetVoltageOutput) GetVoltageOutputRange() (int, error) {
	return get(&p.phidget, func(r *C.PhidgetVoltageOutput_VoltageOutputRange) C.PhidgetReturnCode {
		return C.PhidgetVoltageOutput_getVoltageOutputRange(p.handle, r)
	}, func(r C.PhidgetVoltageOutput_VoltageOutputRange) int { return int(r) })
}

// Close closes the handle and deletes it
func (p *PhidgetVoltageOutput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetVoltageOutput_delete(&p.handle))
}

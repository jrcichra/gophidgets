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

// PhidgetFrequencyCounter wraps a Phidget frequency counter
type PhidgetFrequencyCounter struct {
	phidget
	handle C.PhidgetFrequencyCounterHandle
}

// Create creates a PhidgetFrequencyCounter handle
func (p *PhidgetFrequencyCounter) Create() {
	C.PhidgetFrequencyCounter_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetFrequency returns the measured frequency in Hz
func (p *PhidgetFrequencyCounter) GetFrequency() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetFrequencyCounter_getFrequency(p.handle, r) })
}

// GetMaxFrequency returns the maximum measurable frequency
func (p *PhidgetFrequencyCounter) GetMaxFrequency() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetFrequencyCounter_getMaxFrequency(p.handle, r) })
}

// GetCount returns the total number of pulses counted since the last reset
func (p *PhidgetFrequencyCounter) GetCount() (uint64, error) {
	var r C.uint64_t
	if cerr := C.PhidgetFrequencyCounter_getCount(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return uint64(r), nil
}

// GetTimeElapsed returns the time elapsed since the last reset in seconds
func (p *PhidgetFrequencyCounter) GetTimeElapsed() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetFrequencyCounter_getTimeElapsed(p.handle, r) })
}

// Reset resets the count and elapsed time to zero
func (p *PhidgetFrequencyCounter) Reset() error {
	return p.phidgetError(C.PhidgetFrequencyCounter_reset(p.handle))
}

// SetEnabled enables or disables counting
func (p *PhidgetFrequencyCounter) SetEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetFrequencyCounter_setEnabled(p.handle, boolToCInt(enabled)))
}

// GetEnabled returns whether counting is enabled
func (p *PhidgetFrequencyCounter) GetEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetFrequencyCounter_getEnabled(p.handle, r) })
}

// SetFrequencyCutoff sets the frequency below which the output is reported as zero
func (p *PhidgetFrequencyCounter) SetFrequencyCutoff(hz float64) error {
	return p.phidgetError(C.PhidgetFrequencyCounter_setFrequencyCutoff(p.handle, C.double(hz)))
}

// GetFrequencyCutoff returns the current frequency cutoff
func (p *PhidgetFrequencyCounter) GetFrequencyCutoff() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode {
		return C.PhidgetFrequencyCounter_getFrequencyCutoff(p.handle, r)
	})
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetFrequencyCounter) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetFrequencyCounter_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetFrequencyCounter) GetDataInterval() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetFrequencyCounter_getDataInterval(p.handle, r) })
}

// SetOnFrequencyChangeHandler sets a callback that fires when the measured frequency changes
func (p *PhidgetFrequencyCounter) SetOnFrequencyChangeHandler(f func(float64)) error {
	ctx := gopointer.Save(Passthrough{f})
	return p.phidgetError(C.PhidgetFrequencyCounter_setOnFrequencyChangeHandler(
		p.handle, (C.phidget_double_fcn)(unsafe.Pointer(C.ccallback)), ctx))
}

// SetOnCountChangeHandler sets a callback that fires when new pulses are counted.
// The callback receives: counts (new pulses in this interval), timeChange (seconds).
func (p *PhidgetFrequencyCounter) SetOnCountChangeHandler(f func(uint64, float64)) error {
	ctx := gopointer.Save(f)
	return p.phidgetError(C.PhidgetFrequencyCounter_setOnCountChangeHandler(
		p.handle, (C.phidget_count_fcn)(unsafe.Pointer(C.ccountcallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetFrequencyCounter) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetFrequencyCounter_delete(&p.handle))
}

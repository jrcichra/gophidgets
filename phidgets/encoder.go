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

// PhidgetEncoder wraps a Phidget encoder (quadrature or simple)
type PhidgetEncoder struct {
	phidget
	handle C.PhidgetEncoderHandle
}

// Create creates a PhidgetEncoder handle
func (p *PhidgetEncoder) Create() {
	C.PhidgetEncoder_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetPosition returns the current encoder position in pulses
func (p *PhidgetEncoder) GetPosition() (int64, error) {
	var r C.int64_t
	if cerr := C.PhidgetEncoder_getPosition(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int64(r), nil
}

// SetPosition sets the current position (useful to zero the encoder)
func (p *PhidgetEncoder) SetPosition(pos int64) error {
	return p.phidgetError(C.PhidgetEncoder_setPosition(p.handle, C.int64_t(pos)))
}

// GetIndexPosition returns the position at the most recent index pulse
func (p *PhidgetEncoder) GetIndexPosition() (int64, error) {
	var r C.int64_t
	if cerr := C.PhidgetEncoder_getIndexPosition(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int64(r), nil
}

// SetEnabled enables or disables the encoder
func (p *PhidgetEncoder) SetEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetEncoder_setEnabled(p.handle, boolToCInt(enabled)))
}

// GetEnabled returns whether the encoder is enabled
func (p *PhidgetEncoder) GetEnabled() (bool, error) {
	var r C.int
	if cerr := C.PhidgetEncoder_getEnabled(p.handle, &r); cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

// SetPositionChangeTrigger sets the number of position ticks that must change before the callback fires
func (p *PhidgetEncoder) SetPositionChangeTrigger(trigger uint32) error {
	return p.phidgetError(C.PhidgetEncoder_setPositionChangeTrigger(p.handle, C.uint32_t(trigger)))
}

// GetPositionChangeTrigger returns the current position change trigger
func (p *PhidgetEncoder) GetPositionChangeTrigger() (uint32, error) {
	var r C.uint32_t
	if cerr := C.PhidgetEncoder_getPositionChangeTrigger(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return uint32(r), nil
}

// SetDataInterval sets the data interval in milliseconds
func (p *PhidgetEncoder) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.PhidgetEncoder_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetDataInterval returns the current data interval in milliseconds
func (p *PhidgetEncoder) GetDataInterval() (uint32, error) {
	var r C.uint32_t
	if cerr := C.PhidgetEncoder_getDataInterval(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return uint32(r), nil
}

// SetOnPositionChangeHandler sets a callback that fires on position change events.
// The callback receives: positionChange (ticks since last event), timeChange (seconds), indexTriggered.
func (p *PhidgetEncoder) SetOnPositionChangeHandler(f func(int, float64, bool)) error {
	var pt EncoderPassthrough
	pt.f = f
	ctx := gopointer.Save(pt)
	return p.phidgetError(C.PhidgetEncoder_setOnPositionChangeHandler(
		p.handle, (C.phidget_encoder_fcn)(unsafe.Pointer(C.cencodercallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetEncoder) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetEncoder_delete(&p.handle))
}

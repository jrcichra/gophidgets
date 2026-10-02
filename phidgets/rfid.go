package phidgets

/*
#include <phidget22.h>
#include <stdlib.h>
#include "phidgets.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// PhidgetRFID wraps a Phidget RFID reader
type PhidgetRFID struct {
	phidget
	handle C.PhidgetRFIDHandle
}

// Create creates a PhidgetRFID handle
func (p *PhidgetRFID) Create() {
	C.PhidgetRFID_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetLastTag returns the tag data most recently read by the reader and the
// protocol the tag was written in.
//
// length is the expected tag data length in bytes, which depends on the tag
// protocol; the returned slice has exactly that length.
func (p *PhidgetRFID) GetLastTag(length int) ([]byte, int, error) {
	if length <= 0 {
		return nil, 0, errors.New("gophidgets: tag length must be positive")
	}
	buf := C.malloc(C.size_t(length))
	defer C.free(buf)
	var protocol C.PhidgetRFID_Protocol
	if cerr := C.PhidgetRFID_getLastTag(p.handle, (*C.char)(buf), C.size_t(length), &protocol); cerr != C.EPHIDGET_OK {
		return nil, 0, p.phidgetError(cerr)
	}
	tag := make([]byte, length)
	copy(tag, C.GoBytes(buf, C.int(length)))
	return tag, int(protocol), nil
}

// Write writes data to the tag currently being read by the reader.
func (p *PhidgetRFID) Write(tag string, protocol int, lock bool) error {
	cstr := C.CString(tag)
	defer C.free(unsafe.Pointer(cstr))
	return p.phidgetError(C.PhidgetRFID_write(p.handle, cstr, C.PhidgetRFID_Protocol(protocol), boolToCInt(lock)))
}

// SetAntennaEnabled enables or disables the RFID reader antenna
func (p *PhidgetRFID) SetAntennaEnabled(enabled bool) error {
	return p.phidgetError(C.PhidgetRFID_setAntennaEnabled(p.handle, boolToCInt(enabled)))
}

// GetAntennaEnabled returns whether the antenna is enabled
func (p *PhidgetRFID) GetAntennaEnabled() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetRFID_getAntennaEnabled(p.handle, r) })
}

// GetTagPresent returns whether a tag is currently being read
func (p *PhidgetRFID) GetTagPresent() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetRFID_getTagPresent(p.handle, r) })
}

// SetOnTagHandler sets a callback that fires when a tag is detected.
// The callback receives the tag data and the protocol it was written in.
func (p *PhidgetRFID) SetOnTagHandler(f func(tag string, protocol int)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetRFID_setOnTagHandler(
		p.handle, (C.phidget_rfid_fcn)(unsafe.Pointer(C.rfidtagcallback)), ctx))
}

// SetOnTagLostHandler sets a callback that fires when a detected tag is lost.
// The callback receives the tag data and the protocol it was written in.
func (p *PhidgetRFID) SetOnTagLostHandler(f func(tag string, protocol int)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetRFID_setOnTagLostHandler(
		p.handle, (C.phidget_rfid_fcn)(unsafe.Pointer(C.rfidtagcallback)), ctx))
}

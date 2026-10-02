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

// GetLastTag returns the most recently read tag's data and protocol, even if
// the tag is no longer in range.
func (p *PhidgetRFID) GetLastTag() (string, RFIDProtocol, error) {
	var protocol C.PhidgetRFID_Protocol
	tag, err := readString(&p.phidget, 256, func(buf *C.char, n C.size_t) C.PhidgetReturnCode {
		return C.PhidgetRFID_getLastTag(p.handle, buf, n, &protocol)
	})
	return tag, RFIDProtocol(protocol), err
}

// Write writes data to the tag currently being read by the reader.
func (p *PhidgetRFID) Write(tag string, protocol RFIDProtocol, lock bool) error {
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
func (p *PhidgetRFID) SetOnTagHandler(f func(tag string, protocol RFIDProtocol)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetRFID_setOnTagHandler(
		p.handle, (C.phidget_rfid_fcn)(unsafe.Pointer(C.rfidtagcallback)), ctx))
}

// SetOnTagLostHandler sets a callback that fires when a detected tag is lost.
// The callback receives the tag data and the protocol it was written in.
func (p *PhidgetRFID) SetOnTagLostHandler(f func(tag string, protocol RFIDProtocol)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetRFID_setOnTagLostHandler(
		p.handle, (C.phidget_rfid_fcn)(unsafe.Pointer(C.rfidtagcallback)), ctx))
}

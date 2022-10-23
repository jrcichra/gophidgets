package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"time"
	"unsafe"
)

// Phidget - general phidget interface that all phidgets are derived from (for ease of type management)
type phidget struct {
	handle C.PhidgetHandle
}

type Phidget interface {
	Create()
	Close() error
	String() string

	OpenWaitForAttachment(timeout time.Duration) error

	SetIsRemote(b bool) error
	GetIsRemote() (bool, error)
	SetDeviceSerialNumber(serial int) error
	GetDeviceSerialNumber() (int, error)
	SetHubPort(port int) error
	GetHubPort() (int, error)
	SetChannel(port int) error
	GetChannel() (int, error)

	SetIsHubPortDevice(b bool) error
	GetChannelClassName() (string, error)
	GetChannelName() (string, error)

	// Unexported function for internal management
	getRawHandle() *C.PhidgetHandle
}

// rawHandle updates the base Phidget object with a handle from another object
// This must be called when a new Phidget object is created of a differing class
func (p *phidget) rawHandle(handle unsafe.Pointer) {
	p.handle = (*C.struct__Phidget)(handle)
}

func (p *phidget) getRawHandle() *C.PhidgetHandle {
	return &p.handle
}

// OpenWaitForAttachment opens a phidget and waits for it to be available on the bus
func (p *phidget) OpenWaitForAttachment(timeout time.Duration) error {
	if cerr := C.Phidget_openWaitForAttachment(p.handle, C.uint(timeout.Milliseconds())); cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

func (p *phidget) phidgetError(cerr C.PhidgetReturnCode) error {
	if cerr == C.EPHIDGET_OK {
		return nil
	}
	var errorString *C.char
	var className *C.char
	C.Phidget_getErrorDescription(cerr, &errorString)
	if C.Phidget_getChannelClassName(p.handle, &className) == C.EPHIDGET_OK {
		return fmt.Errorf("%s: %s", C.GoString(className), C.GoString(errorString))
	}
	return errors.New(C.GoString(errorString))
}

// SetIsRemote sets a phidget sensor as a remote device
func (p *phidget) SetIsRemote(b bool) error {
	return p.phidgetError(C.Phidget_setIsRemote(p.handle, boolToCInt(b)))
}

// SetDeviceSerialNumber sets the serial number to use.
// This must be called before calling OpenWaitForAttachment
func (p *phidget) SetDeviceSerialNumber(serial int) error {
	return p.phidgetError(C.Phidget_setDeviceSerialNumber(p.handle, C.int(serial)))
}

// SetHubPort sets a phidget's hub port
func (p *phidget) SetHubPort(port int) error {
	return p.phidgetError(C.Phidget_setHubPort(p.handle, C.int(port)))
}

// GetHubPort gets a phidget's hub port
func (p *phidget) GetHubPort() (int, error) {
	var r C.int
	if cerr := C.Phidget_getHubPort(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// SetChannel sets a phidget motion sensor's channel port
func (p *phidget) SetChannel(port int) error {
	return p.phidgetError(C.Phidget_setChannel(p.handle, C.int(port)))
}

// GetIsRemote gets a phidget's remote status
func (p *phidget) GetIsRemote() (bool, error) {
	var r C.int
	if cerr := C.Phidget_getIsRemote(p.handle, &r); cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

// GetDeviceSerialNumber gets a phidget motion sensor's serial number
func (p *phidget) GetDeviceSerialNumber() (int, error) {
	var r C.int
	cerr := C.Phidget_getDeviceSerialNumber(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// Close - close the handle and delete it
func (p *phidget) Close() error {
	return p.phidgetError(C.Phidget_close(p.handle))
}

// GetChannelClassName gets the name of the channel class the channel belongs to.
func (p *phidget) GetChannelClassName() (string, error) {
	var name *C.char
	cerr := C.Phidget_getChannelClassName(p.handle, &name)
	if cerr != C.EPHIDGET_OK {
		return "", p.phidgetError(cerr)
	}
	return C.GoString(name), nil
}

// SetIsHubPortDevice sets a phidget sensor as a remote device
func (p *phidget) SetIsHubPortDevice(b bool) error {
	return p.phidgetError(C.Phidget_setIsHubPortDevice(p.handle, boolToCInt(b)))
}

// GetDeviceLabel retrieves the label for the device of this handle
func (p *phidget) GetDeviceLabel() (string, error) {
	var cstr *C.char
	if cerr := C.Phidget_getDeviceLabel(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		return "", p.phidgetError(cerr)
	}
	return C.GoString(cstr), nil
}

// GetChannelName retrieves the channel name of this handle
func (p *phidget) GetChannelName() (string, error) {
	var cstr *C.char
	if cerr := C.Phidget_getChannelName(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		return "", p.phidgetError(cerr)
	}
	return C.GoString(cstr), nil
}

// GetChannel retrives which channel this handle is attached to
func (p *phidget) GetChannel() (int, error) {
	var ch C.int
	if cerr := C.Phidget_getChannel(p.handle, &ch); cerr != C.EPHIDGET_OK {
		return -1, p.phidgetError(cerr)
	}
	return int(ch), nil
}

// String returns a string description of the handle
// implements Gos stringer interface to ease printing
func (p *phidget) String() string {
	name, _ := p.GetChannelName()
	class, _ := p.GetChannelClassName()
	channel, _ := p.GetChannel()
	serial, _ := p.GetDeviceSerialNumber()
	label, _ := p.GetDeviceLabel()
	port, _ := p.GetHubPort()

	return fmt.Sprintf("%s: %s channel %d [ser=%x] [label=%s] [port=%d]",
		name, class, channel, serial, label, port)
}

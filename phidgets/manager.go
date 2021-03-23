package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*attach_fcn)(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel);
void cattach_callback(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel);
*/
import "C"
import (
	"errors"
	"fmt"
	"time"
	"sync"
	"unsafe"
)

var (
	handles []*PhidgetHandle
	mutex   sync.Mutex
)

//PhidgetManager is the struct that is a phidget manager handle
type PhidgetManager struct {
	handle C.PhidgetManagerHandle
}

//PhidgetHandle is a generic Phidget handle
type PhidgetHandle struct {
	handle C.PhidgetHandle
}

//export attach_handler
func attach_handler(man C.PhidgetManagerHandle, ctx unsafe.Pointer, channel C.PhidgetHandle) {
	mutex.Lock()
	C.Phidget_retain(channel)
	handles = append(handles, &PhidgetHandle{handle: channel})
	mutex.Unlock()
}

//NewPhidgetManager Create creates a phidget manager
func NewPhidgetManager() (*PhidgetManager, error) {
	p := &PhidgetManager{}
	C.PhidgetManager_create(&p.handle)

	cerr := C.PhidgetManager_setOnAttachHandler(p.handle, (C.attach_fcn)(unsafe.Pointer(C.cattach_callback)), nil)
	if cerr != C.EPHIDGET_OK {
		C.PhidgetManager_delete(&p.handle)
		return nil, errors.New(getErrorDescription(p, cerr))
	}

	if cerr := C.PhidgetManager_open(p.handle); cerr != C.EPHIDGET_OK {
		C.PhidgetManager_delete(&p.handle)
		return nil, errors.New(getErrorDescription(p, cerr))
	}
	// Give the handles time to attach
	time.Sleep(time.Millisecond * 10)

	return p, nil
}

//ListPhidgets returns a list of phidgets that have been discovered
func (m *PhidgetManager) ListPhidgets() []*PhidgetHandle {
	mutex.Lock()
	l := append([]*PhidgetHandle{}, handles...)
	mutex.Unlock()
	return l
}

//Close - close the handle and delete it
func (p *PhidgetManager) Close() error {
	if cerr := C.PhidgetManager_close(p.handle); cerr != C.EPHIDGET_OK {
		return errors.New(getErrorDescription(p, cerr))
	}
	if cerr := C.PhidgetManager_delete(&p.handle); cerr != C.EPHIDGET_OK {
		return errors.New(getErrorDescription(p, cerr))
	}
	return nil
}

//GetChannelClass retrieves the class of channel for this handle
func (p *PhidgetHandle) GetChannelClass() string {
	var cstr *C.char
	if cerr := C.Phidget_getChannelClassName(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		panic(errors.New(getErrorDescription(p, cerr)))
		return ""
	}
	name := C.GoString(cstr)
	return name
}

//GetDeviceSerialNumber retrieves the serial number of the devices this handle is on
func (p *PhidgetHandle) GetDeviceSerialNumber() int {
	var ser C.int32_t
	C.Phidget_getDeviceSerialNumber(p.handle, &ser)
	return int(ser)
}

//GetDeviceLabel retrieves the label for the device of this handle
func (p *PhidgetHandle) GetDeviceLabel() string {
	var cstr *C.char
	if cerr := C.Phidget_getDeviceLabel(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		return ""
	}
	name := C.GoString(cstr)
	return name
}

//GetName retrieves the channel name of this handle
func (p *PhidgetHandle) GetName() string {
	var cstr *C.char
	if cerr := C.Phidget_getChannelName(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		return ""
	}
	name := C.GoString(cstr)
	return name
}

//GetHubPort retrieves which port on the hub this handle is attached to
func (p *PhidgetHandle) GetHubPort() int {
	var port C.int
	C.Phidget_getHubPort(p.handle, &port)
	return int(port)
}

//GetChannel retrives which channel this handle is attached to
func (p *PhidgetHandle) GetChannel() int {
	var ch C.int
	C.Phidget_getChannel(p.handle, &ch)
	return int(ch)
}

//String returns a string description of the handle
func (p *PhidgetHandle) String() string {
	return fmt.Sprintf("%s: %s channel %d [ser=%x] [label=%s] [port=%d]",
		p.GetName(), p.GetChannelClass(), p.GetChannel(), p.GetDeviceSerialNumber(), p.GetDeviceLabel(), p.GetHubPort())
}

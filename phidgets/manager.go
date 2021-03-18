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
	//"reflect"
	"sync"
	"unsafe"
)

var (
	handles []*PhidgetHandle
	mutex sync.Mutex
)

//PhidgetManager is the struct that is a phidget manager handle
type PhidgetManager struct {
	handle C.PhidgetManagerHandle
}

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

//Create creates a phidget manager
func NewPhidgetManager() (*PhidgetManager, error) {
	p := &PhidgetManager{}
	C.PhidgetManager_create(&p.handle)

	cerr := C.PhidgetManager_setOnAttachHandler(p.handle, (C.attach_fcn)(unsafe.Pointer(C.cattach_callback)), nil)
	//cerr := C.PhidgetManager_setOnAttachHandler(p.handle, (C.attach_fcn)(attach_handler), nil)
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

func (p *PhidgetHandle) ChannelClass() string {
	var cstr *C.char
	if cerr := C.Phidget_getChannelClassName(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		panic(errors.New(getErrorDescription(p, cerr)))
		return ""
	}
	name := C.GoString(cstr)
	return name
}

func (p *PhidgetHandle) DeviceSerialNumber() int32 {
	var ser C.int32_t
	C.Phidget_getDeviceSerialNumber(p.handle, &ser)
	return int32(ser)
}

func (p *PhidgetHandle) DeviceLabel() string {
	var cstr *C.char
	if cerr := C.Phidget_getDeviceLabel(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		return ""
	}
	name := C.GoString(cstr)
	return name	
}

func (p *PhidgetHandle) Name() string {
	var cstr *C.char
	if cerr := C.Phidget_getChannelName(p.handle, &cstr); cerr != C.EPHIDGET_OK {
		return ""
	}
	name := C.GoString(cstr)
	return name
}

func (p *PhidgetHandle) HubPort() int {
	var port C.int	
	C.Phidget_getHubPort(p.handle, &port)
	return int(port)
}

func (p *PhidgetHandle) Channel() int {
	var ch C.int	
	C.Phidget_getChannel(p.handle, &ch)
	return int(ch)
}

func (p *PhidgetHandle) String() string {
	return fmt.Sprintf("%s: %s channel %d [ser=%x] [label=%s] [port=%d]", 
		p.Name(), p.ChannelClass(), p.Channel(), p.DeviceSerialNumber(), p.DeviceLabel(), p.HubPort())
}

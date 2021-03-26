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
	"sync"
	"unsafe"
)

var (
	handles []Phidget
	mutex   sync.Mutex
)

//PhidgetManager is the struct that is a phidget manager handle
type PhidgetManager struct {
	handle C.PhidgetManagerHandle
}

//export attach_handler
func attach_handler(man C.PhidgetManagerHandle, ctx unsafe.Pointer, channel C.PhidgetHandle) {
	mutex.Lock()
	defer mutex.Unlock()

	var class C.Phidget_ChannelClass
	if cerr := C.Phidget_getChannelClass(channel, &class); cerr != C.EPHIDGET_OK {
		fmt.Printf("unable to determine class of attached phidget: %d\n", cerr)
		return
	}

	// TODO: All supported phidgets should be here
	switch class {
	case C.PHIDCHCLASS_ACCELEROMETER:
		p := &PhidgetAccelerometer{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetAccelerometerHandle)(unsafe.Pointer(channel))
		handles = append(handles, p)
	case C.PHIDCHCLASS_CURRENTINPUT:
		p := &PhidgetCurrentInput{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetCurrentInputHandle)(unsafe.Pointer(channel))
		handles = append(handles, p)
	case C.PHIDCHCLASS_LCD:
		p := &PhidgetLCD{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetLCDHandle)(unsafe.Pointer(channel))
		handles = append(handles, p)
	case C.PHIDCHCLASS_LIGHTSENSOR:
		p := &PhidgetLightSensor{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetLightSensorHandle)(unsafe.Pointer(channel))
		handles = append(handles, p)
	case C.PHIDCHCLASS_SOUNDSENSOR:
		p := &PhidgetSoundSensor{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetSoundSensorHandle)(unsafe.Pointer(channel))
		handles = append(handles, p)
	case C.PHIDCHCLASS_TEMPERATURESENSOR:
		p := &PhidgetTemperatureSensor{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetTemperatureSensorHandle)(unsafe.Pointer(channel))
		handles = append(handles, p)
	case C.PHIDCHCLASS_VOLTAGEINPUT:
		p := &PhidgetVoltageInput{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetVoltageInputHandle)(unsafe.Pointer(channel))
		handles = append(handles, p)
	case C.PHIDCHCLASS_VOLTAGERATIOINPUT:
		p := &PhidgetVoltageRatioInput{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetVoltageRatioInputHandle)(unsafe.Pointer(channel))
		handles = append(handles, p)
	default:
		fmt.Printf("unsupported phidget discovered: 0x%x\n", class)
		return
	}

	// TODO: We are not doing a Phidget_release at any point to get rid of these
	C.Phidget_retain(channel)
}

//NewPhidgetManager Create creates a phidget manager
func NewPhidgetManager() (*PhidgetManager, error) {
	p := &PhidgetManager{}
	C.PhidgetManager_create(&p.handle)

	cerr := C.PhidgetManager_setOnAttachHandler(p.handle, (C.attach_fcn)(unsafe.Pointer(C.cattach_callback)), nil)
	if cerr != C.EPHIDGET_OK {
		C.PhidgetManager_delete(&p.handle)
		return nil, managerError(cerr)
	}
	// TODO: Should install a PhidgetManager_OnDetachCallback as well so we can
	// support removing the devices

	if cerr := C.PhidgetManager_open(p.handle); cerr != C.EPHIDGET_OK {
		C.PhidgetManager_delete(&p.handle)
		return nil, managerError(cerr)
	}

	return p, nil
}

func managerError(cerr C.PhidgetReturnCode) error {
	if cerr == C.EPHIDGET_OK {
		return nil
	}
	var errorString *C.char
	C.Phidget_getErrorDescription(cerr, &errorString)
	return errors.New(C.GoString(errorString))
}

//ListPhidgets returns a list of phidgets that have been discovered
func (m *PhidgetManager) ListPhidgets() []Phidget {
	mutex.Lock()
	l := append([]Phidget{}, handles...)
	mutex.Unlock()
	return l
}

//Close - close the handle and delete it
func (p *PhidgetManager) Close() error {
	if cerr := C.PhidgetManager_close(p.handle); cerr != C.EPHIDGET_OK {
		return managerError(cerr)
	}
	if cerr := C.PhidgetManager_delete(&p.handle); cerr != C.EPHIDGET_OK {
		return managerError(cerr)
	}
	return nil
}

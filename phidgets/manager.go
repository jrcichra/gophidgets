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

//PhidgetManager is the struct that is a phidget manager handle
type PhidgetManager struct {
	sync.Mutex
	handle  C.PhidgetManagerHandle
	handles []Phidget
}

//export attach_handler
func attach_handler(man C.PhidgetManagerHandle, ctx unsafe.Pointer, channel C.PhidgetHandle) {
	m := (*PhidgetManager)(ctx)

	var class C.Phidget_ChannelClass
	if cerr := C.Phidget_getChannelClass(channel, &class); cerr != C.EPHIDGET_OK {
		fmt.Printf("unable to determine class of attached phidget: %d\n", cerr)
		return
	}

	m.Lock()
	defer m.Unlock()

	// TODO: All supported phidgets should be here
	switch class {
	case C.PHIDCHCLASS_ACCELEROMETER:
		p := &PhidgetAccelerometer{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetAccelerometerHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_CURRENTINPUT:
		p := &PhidgetCurrentInput{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetCurrentInputHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_DIGITALINPUT:
		p := &PhidgetDigitalInput{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetDigitalInputHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_DIGITALOUTPUT:
		p := &PhidgetDigitalOutput{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetDigitalOutputHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_LCD:
		p := &PhidgetLCD{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetLCDHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_LIGHTSENSOR:
		p := &PhidgetLightSensor{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetLightSensorHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_SOUNDSENSOR:
		p := &PhidgetSoundSensor{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetSoundSensorHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_TEMPERATURESENSOR:
		p := &PhidgetTemperatureSensor{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetTemperatureSensorHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_VOLTAGEINPUT:
		p := &PhidgetVoltageInput{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetVoltageInputHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	case C.PHIDCHCLASS_VOLTAGERATIOINPUT:
		p := &PhidgetVoltageRatioInput{}
		p.phidget.handle = channel
		p.handle = (C.PhidgetVoltageRatioInputHandle)(unsafe.Pointer(channel))
		m.handles = append(m.handles, p)
	default:
		fmt.Printf("unsupported phidget discovered: 0x%x\n", class)
		return
	}

	// TODO: We are not doing a Phidget_release at any point to get rid of these
	C.Phidget_retain(channel)
}

//NewPhidgetManager Create creates a phidget manager
func NewPhidgetManager() (*PhidgetManager, error) {
	m := &PhidgetManager{}
	C.PhidgetManager_create(&m.handle)

	cerr := C.PhidgetManager_setOnAttachHandler(m.handle, (C.attach_fcn)(unsafe.Pointer(C.cattach_callback)), unsafe.Pointer(m))
	if cerr != C.EPHIDGET_OK {
		C.PhidgetManager_delete(&m.handle)
		return nil, managerError(cerr)
	}
	// TODO: Should install a PhidgetManager_OnDetachCallback as well so we can
	// support removing the devices

	if cerr := C.PhidgetManager_open(m.handle); cerr != C.EPHIDGET_OK {
		C.PhidgetManager_delete(&m.handle)
		return nil, managerError(cerr)
	}

	return m, nil
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
	m.Lock()
	l := append([]Phidget{}, m.handles...)
	m.Unlock()
	return l
}

//Close - close the handle and delete it
func (m *PhidgetManager) Close() error {
	if cerr := C.PhidgetManager_close(m.handle); cerr != C.EPHIDGET_OK {
		return managerError(cerr)
	}
	if cerr := C.PhidgetManager_delete(&m.handle); cerr != C.EPHIDGET_OK {
		return managerError(cerr)
	}
	m.Lock()
	defer m.Unlock()
	for _, p := range m.handles {
		C.Phidget_release(p.getRawHandle())
	}
	m.handles = []Phidget{}

	return nil
}

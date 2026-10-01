package phidgets

/*
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

// PhidgetManager is the struct that is a phidget manager handle
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

	switch class {
	case C.PHIDCHCLASS_ACCELEROMETER:
		m.handles = append(m.handles, &PhidgetAccelerometer{phidget{handle: channel}, C.PhidgetAccelerometerHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_CURRENTINPUT:
		m.handles = append(m.handles, &PhidgetCurrentInput{phidget{handle: channel}, C.PhidgetCurrentInputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DIGITALINPUT:
		m.handles = append(m.handles, &PhidgetDigitalInput{phidget{handle: channel}, C.PhidgetDigitalInputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DIGITALOUTPUT:
		m.handles = append(m.handles, &PhidgetDigitalOutput{phidget{handle: channel}, C.PhidgetDigitalOutputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DISTANCESENSOR:
		m.handles = append(m.handles, &PhidgetDistanceSensor{phidget{handle: channel}, C.PhidgetDistanceSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_HUMIDITYSENSOR:
		m.handles = append(m.handles, &PhidgetHumiditySensor{phidget{handle: channel}, C.PhidgetHumiditySensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_LCD:
		m.handles = append(m.handles, &PhidgetLCD{phidget{handle: channel}, C.PhidgetLCDHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_LIGHTSENSOR:
		m.handles = append(m.handles, &PhidgetLightSensor{phidget{handle: channel}, C.PhidgetLightSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_SOUNDSENSOR:
		m.handles = append(m.handles, &PhidgetSoundSensor{phidget{handle: channel}, C.PhidgetSoundSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_TEMPERATURESENSOR:
		m.handles = append(m.handles, &PhidgetTemperatureSensor{phidget{handle: channel}, C.PhidgetTemperatureSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_VOLTAGEINPUT:
		m.handles = append(m.handles, &PhidgetVoltageInput{phidget{handle: channel}, C.PhidgetVoltageInputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_VOLTAGERATIOINPUT:
		m.handles = append(m.handles, &PhidgetVoltageRatioInput{phidget{handle: channel}, C.PhidgetVoltageRatioInputHandle(unsafe.Pointer(channel))})
	default:
		p := newGenerated(class, channel)
		if p == nil {
			fmt.Printf("unsupported phidget discovered: 0x%x\n", class)
			return
		}
		m.handles = append(m.handles, p)
	}

	// TODO: We are not doing a Phidget_release at any point to get rid of these
	C.Phidget_retain(channel)
}

// NewPhidgetManager Create creates a phidget manager
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

// ListPhidgets returns a list of phidgets that have been discovered
func (m *PhidgetManager) ListPhidgets() []Phidget {
	m.Lock()
	l := append([]Phidget{}, m.handles...)
	m.Unlock()
	return l
}

// Close - close the handle and delete it
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

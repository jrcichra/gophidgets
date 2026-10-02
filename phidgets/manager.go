package phidgets

/*
#include "phidgets.h"
*/
import "C"
import (
	"fmt"
	"runtime/cgo"
	"sync"
	"unsafe"
)

// PhidgetManager is the struct that is a phidget manager handle
type PhidgetManager struct {
	sync.Mutex
	handle   C.PhidgetManagerHandle
	handles  []Phidget
	onDetach func()     // user callback invoked when a channel detaches
	ctx      cgo.Handle // identifies this manager to C handlers
}

// channelClasses wraps an attached channel in its class-specific type
var channelClasses = map[C.Phidget_ChannelClass]func(C.PhidgetHandle) Phidget{
	C.PHIDCHCLASS_ACCELEROMETER: func(h C.PhidgetHandle) Phidget {
		return &PhidgetAccelerometer{phidget{handle: h}, C.PhidgetAccelerometerHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_BLDCMOTOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetBLDCMotor{phidget{handle: h}, C.PhidgetBLDCMotorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_CAPACITIVETOUCH: func(h C.PhidgetHandle) Phidget {
		return &PhidgetCapacitiveTouch{phidget{handle: h}, C.PhidgetCapacitiveTouchHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_DATAADAPTER: func(h C.PhidgetHandle) Phidget {
		return &PhidgetDataAdapter{phidget{handle: h}, C.PhidgetDataAdapterHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_CURRENTINPUT: func(h C.PhidgetHandle) Phidget {
		return &PhidgetCurrentInput{phidget{handle: h}, C.PhidgetCurrentInputHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_CURRENTOUTPUT: func(h C.PhidgetHandle) Phidget {
		return &PhidgetCurrentOutput{phidget{handle: h}, C.PhidgetCurrentOutputHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_DCMOTOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetDCMotor{phidget{handle: h}, C.PhidgetDCMotorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_DIGITALINPUT: func(h C.PhidgetHandle) Phidget {
		return &PhidgetDigitalInput{phidget{handle: h}, C.PhidgetDigitalInputHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_DICTIONARY: func(h C.PhidgetHandle) Phidget {
		return &PhidgetDictionary{phidget{handle: h}, C.PhidgetDictionaryHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_DIGITALOUTPUT: func(h C.PhidgetHandle) Phidget {
		return &PhidgetDigitalOutput{phidget{handle: h}, C.PhidgetDigitalOutputHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_DISTANCESENSOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetDistanceSensor{phidget{handle: h}, C.PhidgetDistanceSensorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_ENCODER: func(h C.PhidgetHandle) Phidget {
		return &PhidgetEncoder{phidget{handle: h}, C.PhidgetEncoderHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_FREQUENCYCOUNTER: func(h C.PhidgetHandle) Phidget {
		return &PhidgetFrequencyCounter{phidget{handle: h}, C.PhidgetFrequencyCounterHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_FIRMWAREUPGRADE: func(h C.PhidgetHandle) Phidget {
		return &PhidgetFirmwareUpgrade{phidget{handle: h}, C.PhidgetFirmwareUpgradeHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_GENERIC: func(h C.PhidgetHandle) Phidget {
		return &PhidgetGeneric{phidget{handle: h}, C.PhidgetGenericHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_GPS: func(h C.PhidgetHandle) Phidget {
		return &PhidgetGPS{phidget{handle: h}, C.PhidgetGPSHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_GYROSCOPE: func(h C.PhidgetHandle) Phidget {
		return &PhidgetGyroscope{phidget{handle: h}, C.PhidgetGyroscopeHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_HUMIDITYSENSOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetHumiditySensor{phidget{handle: h}, C.PhidgetHumiditySensorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_HUB: func(h C.PhidgetHandle) Phidget {
		return &PhidgetHub{phidget{handle: h}, C.PhidgetHubHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_IR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetIR{phidget{handle: h}, C.PhidgetIRHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_LCD: func(h C.PhidgetHandle) Phidget {
		return &PhidgetLCD{phidget{handle: h}, C.PhidgetLCDHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_LEDARRAY: func(h C.PhidgetHandle) Phidget {
		return &PhidgetLEDArray{phidget{handle: h}, C.PhidgetLEDArrayHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_LIGHTSENSOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetLightSensor{phidget{handle: h}, C.PhidgetLightSensorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_MAGNETOMETER: func(h C.PhidgetHandle) Phidget {
		return &PhidgetMagnetometer{phidget{handle: h}, C.PhidgetMagnetometerHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_MOTORPOSITIONCONTROLLER: func(h C.PhidgetHandle) Phidget {
		return &PhidgetMotorPositionController{phidget{handle: h}, C.PhidgetMotorPositionControllerHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_MOTORVELOCITYCONTROLLER: func(h C.PhidgetHandle) Phidget {
		return &PhidgetMotorVelocityController{phidget{handle: h}, C.PhidgetMotorVelocityControllerHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_PHSENSOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetPHSensor{phidget{handle: h}, C.PhidgetPHSensorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_POWERGUARD: func(h C.PhidgetHandle) Phidget {
		return &PhidgetPowerGuard{phidget{handle: h}, C.PhidgetPowerGuardHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_PRESSURESENSOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetPressureSensor{phidget{handle: h}, C.PhidgetPressureSensorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_RCSERVO: func(h C.PhidgetHandle) Phidget {
		return &PhidgetRCServo{phidget{handle: h}, C.PhidgetRCServoHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_RFID: func(h C.PhidgetHandle) Phidget {
		return &PhidgetRFID{phidget{handle: h}, C.PhidgetRFIDHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_RESISTANCEINPUT: func(h C.PhidgetHandle) Phidget {
		return &PhidgetResistanceInput{phidget{handle: h}, C.PhidgetResistanceInputHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_SOUNDSENSOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetSoundSensor{phidget{handle: h}, C.PhidgetSoundSensorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_SPATIAL: func(h C.PhidgetHandle) Phidget {
		return &PhidgetSpatial{phidget{handle: h}, C.PhidgetSpatialHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_STEPPER: func(h C.PhidgetHandle) Phidget {
		return &PhidgetStepper{phidget{handle: h}, C.PhidgetStepperHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_TEMPERATURESENSOR: func(h C.PhidgetHandle) Phidget {
		return &PhidgetTemperatureSensor{phidget{handle: h}, C.PhidgetTemperatureSensorHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_VOLTAGEINPUT: func(h C.PhidgetHandle) Phidget {
		return &PhidgetVoltageInput{phidget{handle: h}, C.PhidgetVoltageInputHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_VOLTAGEOUTPUT: func(h C.PhidgetHandle) Phidget {
		return &PhidgetVoltageOutput{phidget{handle: h}, C.PhidgetVoltageOutputHandle(unsafe.Pointer(h))}
	},
	C.PHIDCHCLASS_VOLTAGERATIOINPUT: func(h C.PhidgetHandle) Phidget {
		return &PhidgetVoltageRatioInput{phidget{handle: h}, C.PhidgetVoltageRatioInputHandle(unsafe.Pointer(h))}
	},
}

//export attach_handler
func attach_handler(man C.PhidgetManagerHandle, ctx unsafe.Pointer, channel C.PhidgetHandle) {
	m, _ := cgo.Handle(uintptr(ctx)).Value().(*PhidgetManager)
	if m == nil {
		return
	}

	var class C.Phidget_ChannelClass
	if cerr := C.Phidget_getChannelClass(channel, &class); cerr != C.EPHIDGET_OK {
		fmt.Printf("unable to determine class of attached phidget: %d\n", cerr)
		return
	}

	m.Lock()
	defer m.Unlock()

	newPhidget, ok := channelClasses[class]
	if !ok {
		fmt.Printf("unsupported phidget discovered: 0x%x\n", class)
		return
	}
	m.handles = append(m.handles, newPhidget(channel))

	// Keep a reference to the channel; it is released when the channel detaches
	// (manager_detach_handler) or when the manager is closed.
	C.Phidget_retain(channel)
}

//export manager_detach_handler
func manager_detach_handler(man C.PhidgetManagerHandle, ctx unsafe.Pointer, channel C.PhidgetHandle) {
	m, _ := cgo.Handle(uintptr(ctx)).Value().(*PhidgetManager)
	if m == nil {
		return
	}

	m.Lock()
	var removed Phidget
	for i, p := range m.handles {
		if p.getRawHandle() != nil && *p.getRawHandle() == channel {
			removed = p
			m.handles = append(m.handles[:i], m.handles[i+1:]...)
			break
		}
	}
	onDetach := m.onDetach
	m.Unlock()

	if removed != nil {
		C.Phidget_release(&channel)
		if onDetach != nil {
			onDetach()
		}
	}
}

// SetOnDetachHandler sets a callback that is invoked when a channel
// detaches from the hub. Detached handles are removed from ListPhidgets()
// regardless of whether a callback is set.
func (m *PhidgetManager) SetOnDetachHandler(f func()) {
	m.Lock()
	m.onDetach = f
	m.Unlock()
}

// NewPhidgetManager Create creates a phidget manager
func NewPhidgetManager() (*PhidgetManager, error) {
	m := &PhidgetManager{}
	C.PhidgetManager_create(&m.handle)

	m.ctx = cgo.NewHandle(m)
	ctx := C.handlectx(C.uintptr_t(m.ctx))

	cerr := C.PhidgetManager_setOnAttachHandler(m.handle, (C.phidget_manager_fcn)(unsafe.Pointer(C.attach_handler)), ctx)
	if cerr != C.EPHIDGET_OK {
		m.ctx.Delete()
		C.PhidgetManager_delete(&m.handle)
		return nil, newPhidgetError(cerr)
	}
	// Install the detach handler so we can support removing the devices
	if cerr := C.PhidgetManager_setOnDetachHandler(m.handle, (C.phidget_manager_fcn)(unsafe.Pointer(C.manager_detach_handler)), ctx); cerr != C.EPHIDGET_OK {
		m.ctx.Delete()
		C.PhidgetManager_delete(&m.handle)
		return nil, newPhidgetError(cerr)
	}

	if cerr := C.PhidgetManager_open(m.handle); cerr != C.EPHIDGET_OK {
		C.PhidgetManager_delete(&m.handle)
		return nil, newPhidgetError(cerr)
	}

	return m, nil
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
		return newPhidgetError(cerr)
	}
	if cerr := C.PhidgetManager_delete(&m.handle); cerr != C.EPHIDGET_OK {
		return newPhidgetError(cerr)
	}
	m.Lock()
	defer m.Unlock()
	for _, p := range m.handles {
		C.Phidget_release(p.getRawHandle())
	}
	m.handles = []Phidget{}
	// Manager_close has stopped handler threads, so the token is safe to release
	m.ctx.Delete()

	return nil
}

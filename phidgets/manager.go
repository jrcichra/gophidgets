package phidgets

/*
#include <phidget22.h>
typedef void (*attach_fcn)(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel);
typedef void (*detach_fcn)(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel);
void cattach_callback(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel);
void cmanager_detach_callback(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel);
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

// PhidgetManager is the struct that is a phidget manager handle
type PhidgetManager struct {
	sync.Mutex
	handle   C.PhidgetManagerHandle
	handles  []Phidget
	onDetach func()         // user callback invoked when a channel detaches
	ctx      unsafe.Pointer // gopointer token identifying this manager to C handlers
}

//export attach_handler
func attach_handler(man C.PhidgetManagerHandle, ctx unsafe.Pointer, channel C.PhidgetHandle) {
	m, _ := gopointer.Restore(ctx).(*PhidgetManager)
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

	switch class {
	case C.PHIDCHCLASS_ACCELEROMETER:
		m.handles = append(m.handles, &PhidgetAccelerometer{phidget{handle: channel}, C.PhidgetAccelerometerHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_BLDCMOTOR:
		m.handles = append(m.handles, &PhidgetBLDCMotor{phidget{handle: channel}, C.PhidgetBLDCMotorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_CAPACITIVETOUCH:
		m.handles = append(m.handles, &PhidgetCapacitiveTouch{phidget{handle: channel}, C.PhidgetCapacitiveTouchHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DATAADAPTER:
		m.handles = append(m.handles, &PhidgetDataAdapter{phidget{handle: channel}, C.PhidgetDataAdapterHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_CURRENTINPUT:
		m.handles = append(m.handles, &PhidgetCurrentInput{phidget{handle: channel}, C.PhidgetCurrentInputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_CURRENTOUTPUT:
		m.handles = append(m.handles, &PhidgetCurrentOutput{phidget{handle: channel}, C.PhidgetCurrentOutputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DCMOTOR:
		m.handles = append(m.handles, &PhidgetDCMotor{phidget{handle: channel}, C.PhidgetDCMotorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DIGITALINPUT:
		m.handles = append(m.handles, &PhidgetDigitalInput{phidget{handle: channel}, C.PhidgetDigitalInputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DICTIONARY:
		m.handles = append(m.handles, &PhidgetDictionary{phidget{handle: channel}, C.PhidgetDictionaryHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DIGITALOUTPUT:
		m.handles = append(m.handles, &PhidgetDigitalOutput{phidget{handle: channel}, C.PhidgetDigitalOutputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_DISTANCESENSOR:
		m.handles = append(m.handles, &PhidgetDistanceSensor{phidget{handle: channel}, C.PhidgetDistanceSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_ENCODER:
		m.handles = append(m.handles, &PhidgetEncoder{phidget{handle: channel}, C.PhidgetEncoderHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_FREQUENCYCOUNTER:
		m.handles = append(m.handles, &PhidgetFrequencyCounter{phidget{handle: channel}, C.PhidgetFrequencyCounterHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_FIRMWAREUPGRADE:
		m.handles = append(m.handles, &PhidgetFirmwareUpgrade{phidget{handle: channel}, C.PhidgetFirmwareUpgradeHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_GENERIC:
		m.handles = append(m.handles, &PhidgetGeneric{phidget{handle: channel}, C.PhidgetGenericHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_GPS:
		m.handles = append(m.handles, &PhidgetGPS{phidget{handle: channel}, C.PhidgetGPSHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_GYROSCOPE:
		m.handles = append(m.handles, &PhidgetGyroscope{phidget{handle: channel}, C.PhidgetGyroscopeHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_HUMIDITYSENSOR:
		m.handles = append(m.handles, &PhidgetHumiditySensor{phidget{handle: channel}, C.PhidgetHumiditySensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_HUB:
		m.handles = append(m.handles, &PhidgetHub{phidget{handle: channel}, C.PhidgetHubHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_IR:
		m.handles = append(m.handles, &PhidgetIR{phidget{handle: channel}, C.PhidgetIRHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_LCD:
		m.handles = append(m.handles, &PhidgetLCD{phidget{handle: channel}, C.PhidgetLCDHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_LEDARRAY:
		m.handles = append(m.handles, &PhidgetLEDArray{phidget{handle: channel}, C.PhidgetLEDArrayHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_LIGHTSENSOR:
		m.handles = append(m.handles, &PhidgetLightSensor{phidget{handle: channel}, C.PhidgetLightSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_MAGNETOMETER:
		m.handles = append(m.handles, &PhidgetMagnetometer{phidget{handle: channel}, C.PhidgetMagnetometerHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_MOTORPOSITIONCONTROLLER:
		m.handles = append(m.handles, &PhidgetMotorPositionController{phidget{handle: channel}, C.PhidgetMotorPositionControllerHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_MOTORVELOCITYCONTROLLER:
		m.handles = append(m.handles, &PhidgetMotorVelocityController{phidget{handle: channel}, C.PhidgetMotorVelocityControllerHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_PHSENSOR:
		m.handles = append(m.handles, &PhidgetPHSensor{phidget{handle: channel}, C.PhidgetPHSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_POWERGUARD:
		m.handles = append(m.handles, &PhidgetPowerGuard{phidget{handle: channel}, C.PhidgetPowerGuardHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_PRESSURESENSOR:
		m.handles = append(m.handles, &PhidgetPressureSensor{phidget{handle: channel}, C.PhidgetPressureSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_RCSERVO:
		m.handles = append(m.handles, &PhidgetRCServo{phidget{handle: channel}, C.PhidgetRCServoHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_RFID:
		m.handles = append(m.handles, &PhidgetRFID{phidget{handle: channel}, C.PhidgetRFIDHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_RESISTANCEINPUT:
		m.handles = append(m.handles, &PhidgetResistanceInput{phidget{handle: channel}, C.PhidgetResistanceInputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_SOUNDSENSOR:
		m.handles = append(m.handles, &PhidgetSoundSensor{phidget{handle: channel}, C.PhidgetSoundSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_SPATIAL:
		m.handles = append(m.handles, &PhidgetSpatial{phidget{handle: channel}, C.PhidgetSpatialHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_STEPPER:
		m.handles = append(m.handles, &PhidgetStepper{phidget{handle: channel}, C.PhidgetStepperHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_TEMPERATURESENSOR:
		m.handles = append(m.handles, &PhidgetTemperatureSensor{phidget{handle: channel}, C.PhidgetTemperatureSensorHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_VOLTAGEINPUT:
		m.handles = append(m.handles, &PhidgetVoltageInput{phidget{handle: channel}, C.PhidgetVoltageInputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_VOLTAGEOUTPUT:
		m.handles = append(m.handles, &PhidgetVoltageOutput{phidget{handle: channel}, C.PhidgetVoltageOutputHandle(unsafe.Pointer(channel))})
	case C.PHIDCHCLASS_VOLTAGERATIOINPUT:
		m.handles = append(m.handles, &PhidgetVoltageRatioInput{phidget{handle: channel}, C.PhidgetVoltageRatioInputHandle(unsafe.Pointer(channel))})
	default:
		fmt.Printf("unsupported phidget discovered: 0x%x\n", class)
		return
	}

	// Keep a reference to the channel; it is released when the channel detaches
	// (manager_detach_handler) or when the manager is closed.
	C.Phidget_retain(channel)
}

//export manager_detach_handler
func manager_detach_handler(man C.PhidgetManagerHandle, ctx unsafe.Pointer, channel C.PhidgetHandle) {
	m, _ := gopointer.Restore(ctx).(*PhidgetManager)
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

	m.ctx = gopointer.Save(m)

	cerr := C.PhidgetManager_setOnAttachHandler(m.handle, (C.attach_fcn)(unsafe.Pointer(C.cattach_callback)), m.ctx)
	if cerr != C.EPHIDGET_OK {
		gopointer.Unref(m.ctx)
		C.PhidgetManager_delete(&m.handle)
		return nil, managerError(cerr)
	}
	// Install the detach handler so we can support removing the devices
	if cerr := C.PhidgetManager_setOnDetachHandler(m.handle, (C.detach_fcn)(unsafe.Pointer(C.cmanager_detach_callback)), m.ctx); cerr != C.EPHIDGET_OK {
		gopointer.Unref(m.ctx)
		C.PhidgetManager_delete(&m.handle)
		return nil, managerError(cerr)
	}

	if cerr := C.PhidgetManager_open(m.handle); cerr != C.EPHIDGET_OK {
		C.PhidgetManager_delete(&m.handle)
		return nil, managerError(cerr)
	}

	return m, nil
}

func managerError(cerr C.PhidgetReturnCode) error {
	return newPhidgetError(cerr)
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
	// Manager_close has stopped handler threads, so the token is safe to release
	gopointer.Unref(m.ctx)
	m.ctx = nil

	return nil
}

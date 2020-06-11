package phidgets

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

//PhidgetVoltageRatioInput is the struct that is a phidget voltageinputratio sensor
type PhidgetVoltageRatioInput struct {
	handle C.PhidgetVoltageRatioInputHandle
}

//Create creates a phidget voltageinputratio sensor
func (t *PhidgetVoltageRatioInput) Create() {
	C.PhidgetVoltageRatioInput_create(&t.handle)
}

//GetVoltage gets the voltageinputratio from a phidget voltageinputratio sensor
func (t *PhidgetVoltageRatioInput) GetVoltage() float32 {
	var r C.double
	C.PhidgetVoltageRatioInput_getVoltageRatio(t.handle, &r)
	return cDoubleTofloat32(r)
}

//Common to all derived phidgets

func (p *PhidgetVoltageRatioInput) getErrorDescription(cerr C.PhidgetReturnCode) string {
	var errorString *C.char
	C.Phidget_getErrorDescription(cerr, &errorString)
	//Get the name of our class
	t := reflect.TypeOf(p)
	return t.Elem().Name() + ": " + C.GoString(errorString)
}

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetVoltageRatioInput) SetIsRemote(b bool) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setIsRemote(h, boolToCInt(b))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil

}

//SetDeviceSerialNumber sets a phidget lcd sensor's serial number
func (p *PhidgetVoltageRatioInput) SetDeviceSerialNumber(serial int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//SetHubPort sets a phidget lcd sensor's hub port
func (p *PhidgetVoltageRatioInput) SetHubPort(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setHubPort(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetIsRemote gets a phidget lcd sensor's remote status
func (p *PhidgetVoltageRatioInput) GetIsRemote() (bool, error) {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getIsRemote(h, &r)
	if cerr != C.EPHIDGET_OK {
		return false, errors.New(p.getErrorDescription(cerr))
	}
	return cIntTobool(r), nil
}

//GetDeviceSerialNumber gets a phidget lcd sensor's serial number
func (p *PhidgetVoltageRatioInput) GetDeviceSerialNumber() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getDeviceSerialNumber(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//GetHubPort gets a phidget lcd sensor's hub port
func (p *PhidgetVoltageRatioInput) GetHubPort() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getHubPort(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//OpenWaitForAttachment opens a phidget lcd sensor for attachment
func (p *PhidgetVoltageRatioInput) OpenWaitForAttachment(timeout uint) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//Specific to a voltageinputratio - setting the proper sensor type
func (p *PhidgetVoltageRatioInput) SetSensorType(sensorType string) {
	//TODO: need a better way to select a voltage ratio input sensor type by bringing the enum out to go world
	var cSensor C.PhidgetVoltageRatioInput_SensorType
	switch sensorType {
	case "SENSOR_TYPE_1122_DC":
		cSensor = C.SENSOR_TYPE_1122_DC
	default:
		panic("Unknown sensorType: " + sensorType + ". Please add it to the mapping switch in voltageinputratio.go")
	}
	cSensorType := C.PhidgetVoltageRatioInput_SensorType(cSensor)
	C.PhidgetVoltageRatioInput_setSensorType(p.handle, cSensorType)
}

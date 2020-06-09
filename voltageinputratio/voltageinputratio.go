package voltageinputratio

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"
import (
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

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetVoltageRatioInput) SetIsRemote(b bool) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setIsRemote(h, boolToCInt(b))
}

//SetDeviceSerialNumber sets a phidget voltageinputratio sensor's serial number
func (p *PhidgetVoltageRatioInput) SetDeviceSerialNumber(serial int) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
}

//SetHubPort sets a phidget voltageinputratio sensor's hub port
func (p *PhidgetVoltageRatioInput) SetHubPort(port int) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setHubPort(h, intToCInt(port))
}

//GetIsRemote gets a phidget voltageinputratio sensor's remote status
func (p *PhidgetVoltageRatioInput) GetIsRemote() bool {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getIsRemote(h, &r)
	return cIntTobool(r)
}

//GetDeviceSerialNumber gets a phidget voltageinputratio sensor's serial number
func (p *PhidgetVoltageRatioInput) GetDeviceSerialNumber() int {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getDeviceSerialNumber(h, &r)
	return cIntToint(r)
}

//GetHubPort gets a phidget voltageinputratio sensor's hub port
func (p *PhidgetVoltageRatioInput) GetHubPort() int {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getHubPort(h, &r)
	return cIntToint(r)
}

//OpenWaitForAttachment opens a phidget voltageinputratio sensor for attachment
func (p *PhidgetVoltageRatioInput) OpenWaitForAttachment(timeout uint) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
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

//Can't put these in a common module because their type is associated with the module

func boolToCInt(b bool) C.int {
	var r C.int
	if b {
		r = 1
	} else {
		r = 0
	}
	return r
}

func intToBool(i int) bool {
	var b bool
	if i > 0 {
		b = true
	} else {
		b = false
	}
	return b
}

func intToCInt(i int) C.int {
	var c C.int
	c = (C.int)(i)
	return c
}

func cIntToint(c C.int) int {
	var i int
	i = (int)(c)
	return i
}

func cIntTobool(c C.int) bool {
	i := cIntToint(c)
	return intToBool(i)
}

func uintToCUInt(i uint) C.uint {
	var c C.uint
	c = (C.uint)(i)
	return c
}

func cDoubleTofloat32(d C.double) float32 {
	var f float32
	f = (float32)(d)
	return f
}

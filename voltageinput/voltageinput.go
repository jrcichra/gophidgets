package voltageinput

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"
import "unsafe"

//PhidgetVoltageInputHandle is the struct that is a phidget voltageinput sensor
type PhidgetVoltageInputHandle struct {
	handle C.PhidgetVoltageInputHandle
}

//Create creates a phidget voltageinput sensor
func (t *PhidgetVoltageInputHandle) Create() {
	C.PhidgetVoltageInput_create(&t.handle)
}

//GetVoltage gets the voltageinput from a phidget voltageinput sensor
func (t *PhidgetVoltageInputHandle) GetVoltage() float32 {
	var r C.double
	C.PhidgetVoltageInput_getVoltage(t.handle, &r)
	return cDoubleTofloat32(r)
}

//Common to all derived phidgets

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetVoltageInputHandle) SetIsRemote(b bool) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setIsRemote(h, boolToCInt(b))
}

//SetDeviceSerialNumber sets a phidget voltageinput sensor's serial number
func (p *PhidgetVoltageInputHandle) SetDeviceSerialNumber(serial int) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
}

//SetHubPort sets a phidget voltageinput sensor's hub port
func (p *PhidgetVoltageInputHandle) SetHubPort(port int) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setHubPort(h, intToCInt(port))
}

//GetIsRemote gets a phidget voltageinput sensor's remote status
func (p *PhidgetVoltageInputHandle) GetIsRemote() bool {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getIsRemote(h, &r)
	return cIntTobool(r)
}

//GetDeviceSerialNumber gets a phidget voltageinput sensor's serial number
func (p *PhidgetVoltageInputHandle) GetDeviceSerialNumber() int {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getDeviceSerialNumber(h, &r)
	return cIntToint(r)
}

//GetHubPort gets a phidget voltageinput sensor's hub port
func (p *PhidgetVoltageInputHandle) GetHubPort() int {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getHubPort(h, &r)
	return cIntToint(r)
}

//OpenWaitForAttachment opens a phidget voltageinput sensor for attachment
func (p *PhidgetVoltageInputHandle) OpenWaitForAttachment(timeout uint) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
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

package current

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"
import "unsafe"

//PhidgetCurrentInput is the struct that is a phidget current sensor
type PhidgetCurrentInput struct {
	handle C.PhidgetCurrentInputHandle
}

//Create creates a phidget current sensor
func (t *PhidgetCurrentInput) Create() {
	C.PhidgetCurrentInput_create(&t.handle)
}

//GetTemperature gets the current from a phidget current sensor
func (t *PhidgetCurrentInput) GetTemperature() float32 {
	var r C.double
	C.PhidgetCurrentInput_getCurrent(t.handle, &r)
	return cDoubleTofloat32(r)
}

//Common to all derived phidgets

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetCurrentInput) SetIsRemote(b bool) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setIsRemote(h, boolToCInt(b))
}

//SetDeviceSerialNumber sets a phidget current sensor's serial number
func (p *PhidgetCurrentInput) SetDeviceSerialNumber(serial int) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
}

//SetHubPort sets a phidget current sensor's hub port
func (p *PhidgetCurrentInput) SetHubPort(port int) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	C.Phidget_setHubPort(h, intToCInt(port))
}

//GetIsRemote gets a phidget current sensor's remote status
func (p *PhidgetCurrentInput) GetIsRemote() bool {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getIsRemote(h, &r)
	return cIntTobool(r)
}

//GetDeviceSerialNumber gets a phidget current sensor's serial number
func (p *PhidgetCurrentInput) GetDeviceSerialNumber() int {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getDeviceSerialNumber(h, &r)
	return cIntToint(r)
}

//GetHubPort gets a phidget current sensor's hub port
func (p *PhidgetCurrentInput) GetHubPort() int {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	C.Phidget_getHubPort(h, &r)
	return cIntToint(r)
}

//OpenWaitForAttachment opens a phidget current sensor for attachment
func (p *PhidgetCurrentInput) OpenWaitForAttachment(timeout uint) {
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
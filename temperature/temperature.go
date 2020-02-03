package temperature

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"
import (
	"unsafe"
)

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

type TemperatureSensor struct {
	handle C.PhidgetTemperatureSensorHandle
}

func (t *TemperatureSensor) Create() {
	C.PhidgetTemperatureSensor_create(&t.handle)
}

func (t *TemperatureSensor) SetIsRemote(b bool) {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(t.handle))
	C.Phidget_setIsRemote(h, boolToCInt(b))
}

func (t *TemperatureSensor) SetDeviceSerialNumber(serial int) {
	h := (*C.struct__Phidget)(unsafe.Pointer(t.handle))
	C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
}

func (t *TemperatureSensor) SetHubPort(port int) {
	h := (*C.struct__Phidget)(unsafe.Pointer(t.handle))
	C.Phidget_setHubPort(h, intToCInt(port))
}

func (t *TemperatureSensor) GetIsRemote() bool {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(t.handle))
	var r C.int
	C.Phidget_getIsRemote(h, &r)
	return cIntTobool(r)
}

func (t *TemperatureSensor) GetDeviceSerialNumber() int {
	h := (*C.struct__Phidget)(unsafe.Pointer(t.handle))
	var r C.int
	C.Phidget_getDeviceSerialNumber(h, &r)
	return cIntToint(r)
}

func (t *TemperatureSensor) GetHubPort() int {
	h := (*C.struct__Phidget)(unsafe.Pointer(t.handle))
	var r C.int
	C.Phidget_getHubPort(h, &r)
	return cIntToint(r)
}

func (t *TemperatureSensor) OpenWaitForAttachment(timeout uint) {
	h := (*C.struct__Phidget)(unsafe.Pointer(t.handle))
	C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
}

func (t *TemperatureSensor) GetTemperature() float32 {
	var r C.double
	C.PhidgetTemperatureSensor_getTemperature(t.handle, &r)
	return cDoubleTofloat32(r)
}

func AddServer(serverName string, address string, port int, password string, flags int) {
	C.PhidgetNet_addServer(C.CString(serverName), C.CString(address), intToCInt(port), C.CString(password), intToCInt(flags))
}

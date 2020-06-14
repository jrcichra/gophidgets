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

//PhidgetHumiditySensor is the struct that is a phidget humidity sensor
type PhidgetHumiditySensor struct {
	handle C.PhidgetHumiditySensorHandle
}

//Create creates a phidget humidity sensor
func (p *PhidgetHumiditySensor) Create() {
	C.PhidgetHumiditySensor_create(&p.handle)
}

//GetValue gets the humidity from a phidget humidity sensor
func (p *PhidgetHumiditySensor) GetValue() (float32, error) {
	var r C.double
	cerr := C.PhidgetHumiditySensor_getHumidity(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cDoubleTofloat32(r), nil
}

//Common to all derived phidgets

func (p *PhidgetHumiditySensor) getErrorDescription(cerr C.PhidgetReturnCode) string {
	var errorString *C.char
	C.Phidget_getErrorDescription(cerr, &errorString)
	//Get the name of our class
	t := reflect.TypeOf(p)
	return t.Elem().Name() + ": " + C.GoString(errorString)
}

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetHumiditySensor) SetIsRemote(b bool) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setIsRemote(h, boolToCInt(b))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil

}

//SetDeviceSerialNumber sets a phidget lcd sensor's serial number
func (p *PhidgetHumiditySensor) SetDeviceSerialNumber(serial int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//SetHubPort sets a phidget lcd sensor's hub port
func (p *PhidgetHumiditySensor) SetHubPort(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setHubPort(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetIsRemote gets a phidget lcd sensor's remote status
func (p *PhidgetHumiditySensor) GetIsRemote() (bool, error) {
	//Cast humidityHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getIsRemote(h, &r)
	if cerr != C.EPHIDGET_OK {
		return false, errors.New(p.getErrorDescription(cerr))
	}
	return cIntTobool(r), nil
}

//GetDeviceSerialNumber gets a phidget lcd sensor's serial number
func (p *PhidgetHumiditySensor) GetDeviceSerialNumber() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getDeviceSerialNumber(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//GetHubPort gets a phidget lcd sensor's hub port
func (p *PhidgetHumiditySensor) GetHubPort() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getHubPort(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//OpenWaitForAttachment opens a phidget humidity sensor for attachment
func (p *PhidgetHumiditySensor) OpenWaitForAttachment(timeout uint) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//SetChannel sets a phidget humidity sensor's channel port
func (p *PhidgetHumiditySensor) SetChannel(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setChannel(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetChannel gets a phidget humidity sensor's channel port
func (p *PhidgetHumiditySensor) GetChannel() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getChannel(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//Close - close the handle and delete it
func (p *PhidgetHumiditySensor) Close() error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_close(h)
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	cerr = C.PhidgetHumiditySensor_delete((*C.PhidgetHumiditySensorHandle)(&p.handle))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

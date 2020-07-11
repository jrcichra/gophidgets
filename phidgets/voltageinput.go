package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*callback_fcn)(void* handle, void* ctx, double b);
void ccallback(void* handle, void* ctx, double b);  // Forward declaration.
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

//PhidgetVoltageInput is the struct that is a phidget voltageinput sensor
type PhidgetVoltageInput struct {
	handle C.PhidgetVoltageInputHandle
}

//Create creates a phidget voltageinput sensor
func (p *PhidgetVoltageInput) Create() {
	C.PhidgetVoltageInput_create(&p.handle)
}

//GetValue gets the voltageinput from a phidget voltageinput sensor
func (p *PhidgetVoltageInput) GetValue() (float32, error) {
	var r C.double
	cerr := C.PhidgetVoltageInput_getVoltage(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cDoubleTofloat32(r), nil
}

//SetOnVoltageChangeHandler - voltage input for temperature changes calls a function
func (p *PhidgetVoltageInput) SetOnVoltageChangeHandler(f func(Phidget, float32)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	passthrough.handle = p
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetVoltageInput_setOnVoltageChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//Common to all derived phidgets

func (p *PhidgetVoltageInput) getErrorDescription(cerr C.PhidgetReturnCode) string {
	var errorString *C.char
	C.Phidget_getErrorDescription(cerr, &errorString)
	//Get the name of our class
	t := reflect.TypeOf(p)
	return t.Elem().Name() + ": " + C.GoString(errorString)
}

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetVoltageInput) SetIsRemote(b bool) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setIsRemote(h, boolToCInt(b))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil

}

//SetDeviceSerialNumber sets a phidget voltageinput sensor's serial number
func (p *PhidgetVoltageInput) SetDeviceSerialNumber(serial int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//SetHubPort sets a phidget voltageinput sensor's hub port
func (p *PhidgetVoltageInput) SetHubPort(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setHubPort(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetIsRemote gets a phidget voltageinput sensor's remote status
func (p *PhidgetVoltageInput) GetIsRemote() (bool, error) {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getIsRemote(h, &r)
	if cerr != C.EPHIDGET_OK {
		return false, errors.New(p.getErrorDescription(cerr))
	}
	return cIntTobool(r), nil
}

//GetDeviceSerialNumber gets a phidget voltageinput sensor's serial number
func (p *PhidgetVoltageInput) GetDeviceSerialNumber() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getDeviceSerialNumber(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//GetHubPort gets a phidget voltageinput sensor's hub port
func (p *PhidgetVoltageInput) GetHubPort() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getHubPort(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//SetChannel sets a phidget voltageinput sensor's channel port
func (p *PhidgetVoltageInput) SetChannel(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setChannel(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetChannel gets a phidget voltageinput sensor's channel port
func (p *PhidgetVoltageInput) GetChannel() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getChannel(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//OpenWaitForAttachment opens a phidget voltageinput sensor for attachment
func (p *PhidgetVoltageInput) OpenWaitForAttachment(timeout uint) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//Close - close the handle and delete it
func (p *PhidgetVoltageInput) Close() error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_close(h)
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	cerr = C.PhidgetVoltageInput_delete((*C.PhidgetVoltageInputHandle)(&p.handle))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

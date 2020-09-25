package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*callback_fcn)(void* handle, void* ctx, const double acceleration[3], double timestamp);
void ccallback(void* handle, void* ctx, const double acceleration[3], double timestamp);  // Forward declaration.
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

//PhidgetAccelerometer is the struct that is a phidget motion sensor
type PhidgetAccelerometer struct {
	handle C.PhidgetAccelerometerHandle
}

//Create creates a phidget motion sensor
func (p *PhidgetAccelerometer) Create() {
	C.PhidgetAccelerometer_create(&p.handle)
}

//GetAcceleration gets the acceleration from a phidget motion sensor
func (p *PhidgetAccelerometer) GetAcceleration() ([]float32, error) {
	var r [3]C.double
	cerr := C.PhidgetAccelerometer_getAcceleration(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return nil, errors.New(p.getErrorDescription(cerr))
	}
	var ret []float32 = []float32{ (float32)(r[0]), (float32)(r[1]), (float32)(r[2]) }
	return ret, nil
}

//GetMinAcceleration gets the min acceleration value from a phidget motion sensor
func (p *PhidgetAccelerometer) GetMinAcceleration() ([]float32, error) {
	var r [3]C.double
	cerr := C.PhidgetAccelerometer_getMinAcceleration(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return nil, errors.New(p.getErrorDescription(cerr))
	}
	var ret []float32 = []float32{ (float32)(r[0]), (float32)(r[1]), (float32)(r[2]) }
	return ret, nil
}

//GetMaxAcceleration gets the max acceleration value from a phidget motion sensor
func (p *PhidgetAccelerometer) GetMaxAcceleration() ([]float32, error) {
	var r [3]C.double
	cerr := C.PhidgetAccelerometer_getMaxAcceleration(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return nil, errors.New(p.getErrorDescription(cerr))
	}
	var ret []float32 = []float32{ (float32)(r[0]), (float32)(r[1]), (float32)(r[2]) }
	return ret, nil
}

//GetAccelerationChangeTrigger gets the acceleration from a phidget temperature sensor
func (p *PhidgetAccelerometer) GetAccelerationChangeTrigger() (float32, error) {
	var r C.double
	cerr := C.PhidgetAccelerometer_getAccelerationChangeTrigger(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cDoubleTofloat32(r), nil
}

//SetAccelerationChangeTrigger sets the acceleration trigger in the phidget temperature sensor
func (p *PhidgetAccelerometer) SetAccelerationChangeTrigger(value float32) error {
	cerr := C.PhidgetAccelerometer_setAccelerationChangeTrigger(p.handle, float32ToCdouble(value))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetMinAccelerationChangeTrigger sets the min acceleration trigger in the phidget temperature sensor
func (p *PhidgetAccelerometer) GetMinAccelerationChangeTrigger() (float32, error) {
	var r C.double
	cerr := C.PhidgetAccelerometer_getMinAccelerationChangeTrigger(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cDoubleTofloat32(r), nil
}

//GetMaxAccelerationChangeTrigger sets the min acceleration trigger in the phidget temperature sensor
func (p *PhidgetAccelerometer) GetMaxAccelerationChangeTrigger() (float32, error) {
	var r C.double
	cerr := C.PhidgetAccelerometer_getMaxAccelerationChangeTrigger(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cDoubleTofloat32(r), nil
}

//GetAxisCount return the number of axis of the motion sensor
func (p *PhidgetAccelerometer) GetAxisCount() (int, error) {
	var r C.int
	cerr := C.PhidgetAccelerometer_getAxisCount(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//GetDataInterval return the number of axis of the motion sensor
func (p *PhidgetAccelerometer) GetDataInterval() (uint32, error) {
	var r C.uint32_t
	cerr := C.PhidgetAccelerometer_getDataInterval(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return (uint32)(r), nil
}

//SetDataInterval sets the interval between OnAccelerationChange callback calls
func (p *PhidgetAccelerometer) SetDataInterval(value uint32) error {
	cerr := C.PhidgetAccelerometer_setDataInterval(p.handle, (C.uint32_t)(value))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetMinDataInterval return the number of axis of the motion sensor
func (p *PhidgetAccelerometer) GetMinDataInterval() (uint32, error) {
	var r C.uint32_t
	cerr := C.PhidgetAccelerometer_getMinDataInterval(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return (uint32)(r), nil
}

//GetMaxDataInterval return the number of axis of the motion sensor
func (p *PhidgetAccelerometer) GetMaxDataInterval() (uint32, error) {
	var r C.uint32_t
	cerr := C.PhidgetAccelerometer_getMaxDataInterval(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return (uint32)(r), nil
}

//SetOnAccelerationChangeHandler - interrupt for motion changes calls a function
func (p *PhidgetAccelerometer) SetOnAccelerationChangeHandler(f func(Phidget, interface{}, []float32, float32), ctx interface{}) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough MotionPassthrough
	passthrough.f = f
	passthrough.ctx = ctx
	passthrough.handle = p
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetAccelerometer_setOnAccelerationChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//Common to all derived phidgets

func (p *PhidgetAccelerometer) getErrorDescription(cerr C.PhidgetReturnCode) string {
	var errorString *C.char
	C.Phidget_getErrorDescription(cerr, &errorString)
	//Get the name of our class
	t := reflect.TypeOf(p)
	return t.Elem().Name() + ": " + C.GoString(errorString)
}

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetAccelerometer) SetIsRemote(b bool) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setIsRemote(h, boolToCInt(b))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil

}

//SetDeviceSerialNumber sets a phidget motion sensor's serial number
func (p *PhidgetAccelerometer) SetDeviceSerialNumber(serial int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//SetHubPort sets a phidget motion sensor's hub port
func (p *PhidgetAccelerometer) SetHubPort(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setHubPort(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetIsRemote gets a phidget motion sensor's remote status
func (p *PhidgetAccelerometer) GetIsRemote() (bool, error) {
	//Cast MotionHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getIsRemote(h, &r)
	if cerr != C.EPHIDGET_OK {
		return false, errors.New(p.getErrorDescription(cerr))
	}
	return cIntTobool(r), nil
}

//GetDeviceSerialNumber gets a phidget motion sensor's serial number
func (p *PhidgetAccelerometer) GetDeviceSerialNumber() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getDeviceSerialNumber(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//GetHubPort gets a phidget motion sensor's hub port
func (p *PhidgetAccelerometer) GetHubPort() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getHubPort(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//OpenWaitForAttachment opens a phidget motion sensor for attachment
func (p *PhidgetAccelerometer) OpenWaitForAttachment(timeout uint) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//SetChannel sets a phidget motion sensor's channel port
func (p *PhidgetAccelerometer) SetChannel(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setChannel(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetChannel gets a phidget motion sensor's channel port
func (p *PhidgetAccelerometer) GetChannel() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getChannel(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//Close - close the handle and delete it
func (p *PhidgetAccelerometer) Close() error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_close(h)
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	cerr = C.PhidgetAccelerometer_delete((*C.PhidgetAccelerometerHandle)(&p.handle))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

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
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

//PhidgetVoltageRatioInput is the struct that is a phidget voltageinputratio sensor
type PhidgetVoltageRatioInput struct {
	phidget
	handle C.PhidgetVoltageRatioInputHandle
}

//Create creates a phidget voltageinputratio sensor
func (p *PhidgetVoltageRatioInput) Create() {
	C.PhidgetVoltageRatioInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//GetValue gets the voltageinputratio from a phidget voltageinputratio sensor
func (p *PhidgetVoltageRatioInput) GetValue() (float64, error) {
	var r C.double
	if cerr := C.PhidgetVoltageRatioInput_getVoltageRatio(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

//SetOnVoltageRatioChangeHandler - voltage input changes calls a function
func (p *PhidgetVoltageRatioInput) SetOnVoltageRatioChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetVoltageRatioInput_setOnVoltageRatioChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

//SetSensorType - Specific to a voltageinputratio - setting the proper sensor type
func (p *PhidgetVoltageRatioInput) SetSensorType(sensorType string) error {
	//TODO: need a better way to select a voltage ratio input sensor type by bringing the enum out to go world
	var cSensor C.PhidgetVoltageRatioInput_SensorType
	switch sensorType {
	case "SENSOR_TYPE_1122_DC":
		cSensor = C.SENSOR_TYPE_1122_DC
	default:
		return errors.New("Unknown sensorType: " + sensorType + ". Please add it to the mapping switch in voltageinputratio.go")
	}
	return p.phidgetError(C.PhidgetVoltageRatioInput_setSensorType(p.handle, cSensor))
}

//Close - close the handle and delete it
func (p *PhidgetVoltageRatioInput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetVoltageRatioInput_delete(&p.handle))
}

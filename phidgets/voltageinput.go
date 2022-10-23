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

var voltageInputSensorTypeMap map[string]C.PhidgetVoltageInput_SensorType = map[string]C.PhidgetVoltageInput_SensorType{
	"SENSOR_TYPE_VOLTAGE":  0x0,
	"SENSOR_TYPE_1114":     0x2b84,
	"SENSOR_TYPE_1117":     0x2ba2,
	"SENSOR_TYPE_1123":     0x2bde,
	"SENSOR_TYPE_1127":     0x2c06,
	"SENSOR_TYPE_1130_PH":  0x2c25,
	"SENSOR_TYPE_1130_ORP": 0x2c26,
	"SENSOR_TYPE_1132":     0x2c38,
	"SENSOR_TYPE_1133":     0x2c42,
	"SENSOR_TYPE_1135":     0x2c56,
	"SENSOR_TYPE_1142":     0x2c9c,
	"SENSOR_TYPE_1143":     0x2ca6,
	"SENSOR_TYPE_3500":     0x88b8,
	"SENSOR_TYPE_3501":     0x88c2,
	"SENSOR_TYPE_3502":     0x88cc,
	"SENSOR_TYPE_3503":     0x88d6,
	"SENSOR_TYPE_3507":     0x88fe,
	"SENSOR_TYPE_3508":     0x8908,
	"SENSOR_TYPE_3509":     0x8912,
	"SENSOR_TYPE_3510":     0x891c,
	"SENSOR_TYPE_3511":     0x8926,
	"SENSOR_TYPE_3512":     0x8930,
	"SENSOR_TYPE_3513":     0x893a,
	"SENSOR_TYPE_3514":     0x8944,
	"SENSOR_TYPE_3515":     0x894e,
	"SENSOR_TYPE_3516":     0x8958,
	"SENSOR_TYPE_3517":     0x8962,
	"SENSOR_TYPE_3518":     0x896c,
	"SENSOR_TYPE_3519":     0x8976,
	"SENSOR_TYPE_3584":     0x8c00,
	"SENSOR_TYPE_3585":     0x8c0a,
	"SENSOR_TYPE_3586":     0x8c14,
	"SENSOR_TYPE_3587":     0x8c1e,
	"SENSOR_TYPE_3588":     0x8c28,
	"SENSOR_TYPE_3589":     0x8c32,
}

// PhidgetVoltageInput is the struct that is a phidget voltageinput sensor
type PhidgetVoltageInput struct {
	phidget
	handle C.PhidgetVoltageInputHandle
}

// Create creates a phidget voltageinput sensor
func (p *PhidgetVoltageInput) Create() {
	C.PhidgetVoltageInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the voltageinput from a phidget voltageinput sensor
func (p *PhidgetVoltageInput) GetValue() (float64, error) {
	var r C.double
	cerr := C.PhidgetVoltageInput_getVoltage(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetOnVoltageChangeHandler - voltage input for temperature changes calls a function
func (p *PhidgetVoltageInput) SetOnVoltageChangeHandler(f func(float64)) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetVoltageInput_setOnVoltageChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

// SetSensorType - Specific to a voltageinputratio - setting the proper sensor type
func (p *PhidgetVoltageInput) SetSensorType(sensorType string) error {
	if sensorCode, ok := voltageInputSensorTypeMap[sensorType]; ok {
		return p.phidgetError(C.PhidgetVoltageInput_setSensorType(p.handle, sensorCode))
	}
	return errors.New("Unknown sensorType: " + sensorType + ". Please add it to the mapping switch in gophidgets voltageinput.go")
}

// Close - close the handle and delete it
func (p *PhidgetVoltageInput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetVoltageInput_delete(&p.handle))
}

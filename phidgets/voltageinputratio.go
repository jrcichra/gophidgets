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

var voltageRatioInputSensorTypeMap map[string]C.PhidgetVoltageRatioInput_SensorType = map[string]C.PhidgetVoltageRatioInput_SensorType{
	"SENSOR_TYPE_VOLTAGERATIO":      0x0,
	"SENSOR_TYPE_1101_SHARP_2D120X": 0x2b03,
	"SENSOR_TYPE_1101_SHARP_2Y0A21": 0x2b04,
	"SENSOR_TYPE_1101_SHARP_2Y0A02": 0x2b05,
	"SENSOR_TYPE_1102":              0x2b0c,
	"SENSOR_TYPE_1103":              0x2b16,
	"SENSOR_TYPE_1104":              0x2b20,
	"SENSOR_TYPE_1105":              0x2b2a,
	"SENSOR_TYPE_1106":              0x2b34,
	"SENSOR_TYPE_1107":              0x2b3e,
	"SENSOR_TYPE_1108":              0x2b48,
	"SENSOR_TYPE_1109":              0x2b52,
	"SENSOR_TYPE_1110":              0x2b5c,
	"SENSOR_TYPE_1111":              0x2b66,
	"SENSOR_TYPE_1112":              0x2b70,
	"SENSOR_TYPE_1113":              0x2b7a,
	"SENSOR_TYPE_1115":              0x2b8e,
	"SENSOR_TYPE_1116":              0x2b98,
	"SENSOR_TYPE_1118_AC":           0x2bad,
	"SENSOR_TYPE_1118_DC":           0x2bae,
	"SENSOR_TYPE_1119_AC":           0x2bb7,
	"SENSOR_TYPE_1119_DC":           0x2bb8,
	"SENSOR_TYPE_1120":              0x2bc0,
	"SENSOR_TYPE_1121":              0x2bca,
	"SENSOR_TYPE_1122_AC":           0x2bd5,
	"SENSOR_TYPE_1122_DC":           0x2bd6,
	"SENSOR_TYPE_1124":              0x2be8,
	"SENSOR_TYPE_1125_HUMIDITY":     0x2bf3,
	"SENSOR_TYPE_1125_TEMPERATURE":  0x2bf4,
	"SENSOR_TYPE_1126":              0x2bfc,
	"SENSOR_TYPE_1128":              0x2c10,
	"SENSOR_TYPE_1129":              0x2c1a,
	"SENSOR_TYPE_1131":              0x2c2e,
	"SENSOR_TYPE_1134":              0x2c4c,
	"SENSOR_TYPE_1136":              0x2c60,
	"SENSOR_TYPE_1137":              0x2c6a,
	"SENSOR_TYPE_1138":              0x2c74,
	"SENSOR_TYPE_1139":              0x2c7e,
	"SENSOR_TYPE_1140":              0x2c88,
	"SENSOR_TYPE_1141":              0x2c92,
	"SENSOR_TYPE_1146":              0x2cc4,
	"SENSOR_TYPE_3120":              0x79e0,
	"SENSOR_TYPE_3121":              0x79ea,
	"SENSOR_TYPE_3122":              0x79f4,
	"SENSOR_TYPE_3123":              0x79fe,
	"SENSOR_TYPE_3130":              0x7a44,
	"SENSOR_TYPE_3520":              0x8980,
	"SENSOR_TYPE_3521":              0x898a,
	"SENSOR_TYPE_3522":              0x8994,
}

// PhidgetVoltageRatioInput is the struct that is a phidget voltageinputratio sensor
type PhidgetVoltageRatioInput struct {
	phidget
	handle C.PhidgetVoltageRatioInputHandle
}

// Create creates a phidget voltageinputratio sensor
func (p *PhidgetVoltageRatioInput) Create() {
	C.PhidgetVoltageRatioInput_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetValue gets the voltageinputratio from a phidget voltageinputratio sensor
func (p *PhidgetVoltageRatioInput) GetValue() (float64, error) {
	var r C.double
	if cerr := C.PhidgetVoltageRatioInput_getSensorValue(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetOnVoltageRatioChangeHandler - voltage input changes calls a function
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

// SetSensorType - Specific to a voltageinputratio - setting the proper sensor type
func (p *PhidgetVoltageRatioInput) SetSensorType(sensorType string) error {
	if sensorCode, ok := voltageRatioInputSensorTypeMap[sensorType]; ok {
		return p.phidgetError(C.PhidgetVoltageRatioInput_setSensorType(p.handle, sensorCode))
	}
	return errors.New("Unknown sensorType: " + sensorType + ". Please add it to the mapping switch in gophidgets voltageinputratio.go")
}

// Close - close the handle and delete it
func (p *PhidgetVoltageRatioInput) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetVoltageRatioInput_delete(&p.handle))
}

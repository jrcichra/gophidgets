# gophidgets [![Go Report Card](https://goreportcard.com/badge/github.com/jrcichra/gophidgets)](https://goreportcard.com/report/github.com/jrcichra/gophidgets)

Golang bindings for the Phidgets C library

# Changelog

- 2026-10-01 - Wrapped every remaining libphidget22 channel class: `MotorPositionController`, `MotorVelocityController`, `CurrentOutput`, `RFID`, `Hub`, `IR`, `PowerGuard`, `LEDArray`, `Dictionary`, `FirmwareUpgrade`, `DataAdapter`, and `Generic`, each registerable by the `PhidgetManager`. Completed the `PhidgetLCD` API (framebuffers, font sizes, bitmaps, backlight/contrast limits) and filled in the rest of the base `Phidget` API: device identity and firmware info, channel class/persistence, hub identity and port-speed/VINT-speed settings, data rate and data interval, server info, local/channel state, property-change handlers, reboot, and label/flash writes. Added state handlers to `DigitalInput`, distance handlers to `DistanceSensor`, and a fix-state handler to `GPS` that now reports a true/false fix. `DigitalOutput` gained duty-cycle and LED-current-limit controls. The base `Phidget` now exposes `Open()`, attach/detach, and error handlers, and `PhidgetManager` removes detached channels from `ListPhidgets()` (with an optional `SetOnDetachHandler`). Errors are now a typed `PhidgetError` carrying the `EPHIDGET_*` code, so `errors.Is` can match sentinels like `phidgets.ErrTimeout`, and the `EEPHIDGET_*` event codes delivered to error handlers are named.
- 2025/02/10 - Changed `GetSonarReflections()` to only return `distance` and `amplitude` slices. The count is not required as Go makes it simple to determine. A check was added to validate the lengths are the same.
- 2022/11/24 - VoltageInput and VoltageInputRatio `GetValue()` always called `getVoltage()`, not `getSensorValue()`. I broke out the functions to match the Phidget's library names since `VoltageInput` and `VoltageRatioInput` are used in different ways based on the hardware.

## Install

`go get "github.com/jrcichra/gophidgets/phidgets"`

## Supported devices

- Accelerometer, BLDC Motor, Capacitive Touch, Current Input, Current
  Output, Data Adapter, DC Motor, Dictionary, Digital Input, Digital
  Output, Distance Sensor, Encoder, Firmware Upgrade, Frequency
  Counter, Generic, GPS, Gyroscope, Hub, Humidity Sensor, IR, LCD,
  LED Array, Light Sensor, Magnetometer, Motor Position Controller,
  Motor Velocity Controller, PH Sensor, Power Guard, Pressure Sensor,
  RC Servo, RFID, Resistance Input, Sound Sensor, Spatial, Stepper,
  Temperature Sensor, Voltage Input, Voltage Ratio Input, Voltage Output

## Requirements

Link against a current libphidget22 build — the Motor Velocity
Controller, Current Output, and a few other APIs are only present in the
2026 1.26.x generation that CI builds from
`https://www.phidgets.com/downloads/phidget22/libraries/linux/libphidget22.tar.gz`.
Older installations will fail to compile the new classes.

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/jrcichra/gophidgets/phidgets"
)

func main() {
	t := phidgets.PhidgetTemperatureSensor{}
	t.Create()
	t.SetIsRemote(true)
	t.SetDeviceSerialNumber(11111)
	t.SetHubPort(0)
	if err := t.OpenWaitForAttachment(2 * time.Second); err != nil {
		panic(err)
	}
	//Loop forever
	for {
		temp, err := t.GetValue()
		if err != nil {
			panic(err)
		}
		fmt.Println("Temperature is", temp*9.0/5.0+32)
		time.Sleep(5 * time.Second)
	}
}
```

package main

import (
	"fmt"
	"time"

	"github.com/jrcichra/gophidgets/lcd"

	"github.com/jrcichra/gophidgets/voltageinputratio"

	"github.com/jrcichra/gophidgets/humidity"
	"github.com/jrcichra/gophidgets/phidgetnet"
	"github.com/jrcichra/gophidgets/temperature"
)

func main() {

	var err error

	phidgetnet.AddServer("Justin", "10.0.0.176", 5661, "", 0)

	t := temperature.PhidgetTemperatureSensor{}
	t.Create()
	t.SetIsRemote(true)
	t.SetDeviceSerialNumber(597101)
	t.SetHubPort(0)
	err = t.OpenWaitForAttachment(2000)
	if err != nil {
		panic(err)
	}

	h := humidity.PhidgetHumiditySensor{}
	h.Create()
	h.SetIsRemote(true)
	h.SetDeviceSerialNumber(597101)
	h.SetHubPort(0)
	err = h.OpenWaitForAttachment(2000)
	if err != nil {
		panic(err)
	}

	vr := voltageinputratio.PhidgetVoltageRatioInput{}
	vr.Create()
	vr.SetSensorType("SENSOR_TYPE_1122_DC")

	lcd := lcd.PhidgetLCD{}
	lcd.Create()
	lcd.SetDeviceSerialNumber(597102)
	lcd.SetHubPort(5)
	lcd.SetIsRemote(true)
	lcd.SetBacklight(.55)
	err = lcd.OpenWaitForAttachment(2000)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		temperature := t.GetTemperature()*9.0/5.0 + 32
		fmt.Println("Temperature is", temperature)
		fmt.Println("Humidity is", h.GetHumidity())
		lcd.SetText(fmt.Sprintf("Justin: %f", temperature))
		time.Sleep(time.Duration(5) * time.Second)
	}
}

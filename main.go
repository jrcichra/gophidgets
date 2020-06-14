package main

import (
	"fmt"
	"time"

	"github.com/jrcichra/gophidgets/phidgets"
)

func main() {

	var err error

	//Array of generic phidget sensors
	sensors := make([]phidgets.Phidget, 0)

	phidgets.AddServer("Justin", "10.0.0.176", 5661, "", 0)

	t := phidgets.PhidgetTemperatureSensor{}
	t.Create()
	t.SetIsRemote(true)
	t.SetDeviceSerialNumber(597101)
	t.SetHubPort(0)
	err = t.OpenWaitForAttachment(2000)
	if err != nil {
		panic(err)
	}
	sensors = append(sensors, &t)

	h := phidgets.PhidgetHumiditySensor{}
	h.Create()
	h.SetIsRemote(true)
	h.SetDeviceSerialNumber(597101)
	h.SetHubPort(0)
	err = h.OpenWaitForAttachment(2000)
	if err != nil {
		panic(err)
	}
	sensors = append(sensors, &h)

	vr := phidgets.PhidgetVoltageRatioInput{}
	vr.Create()
	vr.SetSensorType("SENSOR_TYPE_1122_DC")

	lcd := phidgets.PhidgetLCD{}
	lcd.Create()
	lcd.SetDeviceSerialNumber(597101)
	lcd.SetHubPort(5)
	lcd.SetIsRemote(true)
	lcd.SetBacklight(.55)
	err = lcd.OpenWaitForAttachment(2000)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		for _, sensor := range sensors {
			switch s := sensor.(type) {
			case *phidgets.PhidgetTemperatureSensor:
				fmt.Println("Temperature is", s.GetValue()*9.0/5.0+32)
				lcd.SetText(fmt.Sprintf("Justin: %f", s.GetValue()*9.0/5.0+32))
			case *phidgets.PhidgetHumiditySensor:
				fmt.Println("Humidity is", s.GetValue())
			}
		}
		time.Sleep(time.Duration(5) * time.Second)
	}

	//Close the sensors
	for _, sensor := range sensors {
		sensor.Close()
	}

	lcd.Close()
}

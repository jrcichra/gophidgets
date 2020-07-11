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
	h.SetOnHumidityChangeHandler(func(p phidgets.Phidget, ctx interface{}, value float32) {
		fmt.Println("I got a humidity of", value)
		serial, _ := p.GetDeviceSerialNumber()
		fmt.Println("My phidget serial is", serial)
	}, nil)
	// sensors = append(sensorqs, &h)

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
				val, _ := s.GetValue()
				val = val*9.0/5.0 + 32
				fmt.Println("Temperature is")
				lcd.SetText(fmt.Sprintf("Justin: %f", val))
			case *phidgets.PhidgetHumiditySensor:
				hum, _ := s.GetValue()
				fmt.Println("Humidity is", hum)
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

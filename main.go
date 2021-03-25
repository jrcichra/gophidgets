package main

import (
	"fmt"
	"io"
	"time"

	"github.com/jrcichra/gophidgets/phidgets"
)

func main() {
	//Array of generic phidget sensors
	sensors := make([]phidgets.Phidget, 0)

	phidgets.AddServer("Justin", "10.0.0.176", 5661, "", 0)

	t := phidgets.PhidgetTemperatureSensor{}
	t.Create()
	t.SetIsRemote(true)
	t.SetDeviceSerialNumber(597101)
	t.SetHubPort(0)
	if err := t.OpenWaitForAttachment(time.Second * 2); err != nil {
		panic(err)
	}
	if err := t.SetOnTemperatureChangeHandler(func(val float64) {
		fmt.Printf("temp change: %f\n", val)
	}); err != nil {
		panic(err)
	}
	sensors = append(sensors, &t)

	h := phidgets.PhidgetHumiditySensor{}
	h.Create()
	h.SetIsRemote(true)
	h.SetDeviceSerialNumber(597101)
	h.SetHubPort(0)
	if err := h.OpenWaitForAttachment(time.Second * 2); err != nil {
		panic(err)
	}
	h.SetOnHumidityChangeHandler(func(value float64) {
		fmt.Println("I got a humidity of", value)
		serial, _ := h.GetDeviceSerialNumber()
		fmt.Println("My phidget serial is", serial)
	})
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
	if err := lcd.OpenWaitForAttachment(time.Second * 2); err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		for _, sensor := range sensors {
			switch s := sensor.(type) {
			case *phidgets.PhidgetTemperatureSensor:
				val, _ := s.GetValue()
				val = val*9.0/5.0 + 32
				fmt.Printf("Temperature is %f Fahrenheit\n", val)
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
		s := sensor.(io.Closer)
		s.Close()
	}

	lcd.Close()
}

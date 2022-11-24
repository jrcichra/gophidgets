package main

import (
	"fmt"
	"time"

	"github.com/jrcichra/gophidgets/phidgets"
)

func main() {
	m, err := phidgets.NewPhidgetManager()
	if err != nil {
		panic(err)
	}

	// Sometimes the phidgets take a while to attach
	fmt.Printf("Starting phidget discovery...\n")
	time.Sleep(1000 * time.Millisecond)

	available := m.ListPhidgets()
	fmt.Printf("Found %d phidgets\n", len(available))
	for i, p := range available {
		fmt.Printf("  %d: %s\n", i, p)
		if err := p.OpenWaitForAttachment(time.Second); err != nil {
			fmt.Printf("Failed to open: %s\n", err)
			continue
		}
		switch s := p.(type) {
		case *phidgets.PhidgetCurrentInput:
			val, _ := s.GetValue()
			fmt.Printf("Current is %f\n", val)
		case *phidgets.PhidgetDigitalInput:
			val, _ := s.GetState()
			fmt.Printf("Input state: %t\n", val)
		case *phidgets.PhidgetDigitalOutput:
			val, _ := s.GetState()
			fmt.Printf("Output state: %t\n", val)
		case *phidgets.PhidgetTemperatureSensor:
			val, _ := s.GetValue()
			fmt.Printf("Temperature is %f\n", val)
		case *phidgets.PhidgetHumiditySensor:
			hum, _ := s.GetValue()
			fmt.Printf("Humidity is %f\n", hum)
		case *phidgets.PhidgetVoltageInput:
			val, _ := s.GetVoltage()
			fmt.Printf("Voltage is %f\n", val)
		case *phidgets.PhidgetVoltageRatioInput:
			val, _ := s.GetVoltageRatio()
			fmt.Printf("Voltage Ratio: %f\n", val)
		}
		p.Close()
	}
	m.Close()

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
		sensor.Close()
	}

	lcd.Close()
}

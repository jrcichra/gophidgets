package main

import (
	"fmt"
	"time"

	"github.com/jrcichra/gophidgets/temperature"
)

func main() {
	t := temperature.TemperatureSensor{}
	temperature.AddServer("Justin", "10.0.0.176", 5661, "", 0)
	t.Create()
	t.SetIsRemote(true)
	t.SetDeviceSerialNumber(597101)
	t.SetHubPort(0)
	t.OpenWaitForAttachment(20000)
	for i := 0; i < 5; i++ {
		fmt.Println("temperature is", t.GetTemperature()*9.0/5.0+32)
		time.Sleep(time.Duration(5) * time.Second)
	}
}

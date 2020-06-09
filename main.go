package main

import (
	"fmt"
	"time"

	"github.com/jrcichra/gophidgets/voltageinputratio"

	"github.com/jrcichra/gophidgets/humidity"
	"github.com/jrcichra/gophidgets/phidgetnet"
	"github.com/jrcichra/gophidgets/temperature"
)

func main() {
	phidgetnet.AddServer("Justin", "10.0.0.176", 5661, "", 0)

	t := temperature.PhidgetTemperatureSensor{}
	t.Create()
	t.SetIsRemote(true)
	t.SetDeviceSerialNumber(597101)
	t.SetHubPort(0)
	t.OpenWaitForAttachment(20000)

	h := humidity.PhidgetHumiditySensor{}
	h.Create()
	h.SetIsRemote(true)
	h.SetDeviceSerialNumber(597101)
	h.SetHubPort(0)
	h.OpenWaitForAttachment(20000)

	vr := voltageinputratio.PhidgetVoltageRatioInput{}
	vr.Create()
	vr.SetSensorType("SENSOR_TYPE_1122_DC")

	for i := 0; i < 5; i++ {
		fmt.Println("Temperature is", t.GetTemperature()*9.0/5.0+32)
		fmt.Println("Humidity is", h.GetHumidity())
		time.Sleep(time.Duration(5) * time.Second)
	}
}

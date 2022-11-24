# gophidgets [![Go Report Card](https://goreportcard.com/badge/github.com/jrcichra/gophidgets)](https://goreportcard.com/report/github.com/jrcichra/gophidgets)

Golang bindings for the Phidgets C library

# Release Notes

- 11/24/2022 - VoltageInput and VoltageInputRatio `GetValue()` mitakenly called `getVoltage()`, not `getSensorValue()`. It will now pay attention to the sensor type provided.

## Install

`go get "github.com/jrcichra/gophidgets/phidgets"`

## Example

```go
t := phidgets.PhidgetTemperatureSensor{}
t.Create()
t.SetIsRemote(true)
t.SetDeviceSerialNumber(11111)
t.SetHubPort(0)
err = t.OpenWaitForAttachment(2000)
if err != nil {
    panic(err)
}
//Loop forever
for {
    fmt.Println("Temperature is", t.GetValue()*9.0/5.0+32)
    time.Sleep(time.Duration(5) * time.Second)
}
```

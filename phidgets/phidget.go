package phidgets

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"

//Phidget - general phidget interface that all phidgets are derived from (for ease of type management)
type Phidget interface {
	Create()
	getErrorDescription(C.PhidgetReturnCode) string
	SetIsRemote(bool) error
	SetDeviceSerialNumber(int) error
	SetHubPort(int) error
	GetIsRemote() (bool, error)
	GetDeviceSerialNumber() (int, error)
	GetHubPort() (int, error)
	OpenWaitForAttachment(uint) error
}

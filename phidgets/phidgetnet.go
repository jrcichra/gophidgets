package phidgets

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"

//AddServer adds a
func AddServer(serverName string, address string, port int, password string, flags int) {
	C.PhidgetNet_addServer(C.CString(serverName), C.CString(address), intToCInt(port), C.CString(password), intToCInt(flags))
}

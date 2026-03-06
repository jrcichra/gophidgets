package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
*/
import "C"
import "unsafe"

// AddServer adds a network server for the phidget library to connect to.
func AddServer(serverName string, address string, port int, password string, flags int) {
	serverC := C.CString(serverName)
	addressC := C.CString(address)
	passwordC := C.CString(password)

	C.PhidgetNet_addServer(serverC, addressC, C.int(port), passwordC, C.int(flags))

	C.free(unsafe.Pointer(serverC))
	C.free(unsafe.Pointer(addressC))
	C.free(unsafe.Pointer(passwordC))
}

// RemoveServer removes a previously added network server entry, killing its reconnect thread.
func RemoveServer(serverName string) {
	serverC := C.CString(serverName)
	C.PhidgetNet_removeServer(serverC)
	C.free(unsafe.Pointer(serverC))
}

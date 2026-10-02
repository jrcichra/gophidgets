package phidgets

/*
#include <phidget22.h>
*/
import "C"
import (
	"unsafe"
)

// PhidgetGeneric wraps a Generic Phidget channel — a channel with no
// class-specific API, usable through the base Phidget interface.
type PhidgetGeneric struct {
	phidget
	handle C.PhidgetGenericHandle
}

// Create creates a PhidgetGeneric handle
func (p *PhidgetGeneric) Create() {
	C.PhidgetGeneric_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// Close closes the handle and deletes it
func (p *PhidgetGeneric) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetGeneric_delete(&p.handle))
}

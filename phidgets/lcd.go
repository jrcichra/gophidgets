package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*callback_fcn)(void* handle, void* ctx, double b);
void ccallback(void* handle, void* ctx, double b);  // Forward declaration.
*/
import "C"
import (
	"unsafe"
)

//PhidgetLCD is the struct that is a phidget lcd sensor
type PhidgetLCD struct {
	Phidget
	handle C.PhidgetLCDHandle
}

//Create creates a phidget lcd sensor
func (p *PhidgetLCD) Create() {
	C.PhidgetLCD_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

//SetText sets the lcd text
func (p *PhidgetLCD) SetText(text string) error {
	if cerr := C.PhidgetLCD_writeText(p.handle, C.FONT_6x12, 40, 25, C.CString(text)); cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return p.phidgetError(C.PhidgetLCD_flush(p.handle))
}

//SetBacklight - sets the backlight value
func (p *PhidgetLCD) SetBacklight(brightness float32) error {
	return p.phidgetError(C.PhidgetLCD_setBacklight(p.handle, C.double(brightness)))
}

//Close - close the handle and delete it
func (p *PhidgetLCD) Close() error {
	if err := p.Phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetLCD_delete(&p.handle))
}

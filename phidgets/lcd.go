package phidgets

/*
#include <stdlib.h>
#include <phidget22.h>
*/
import "C"
import (
	"unsafe"
)

// PhidgetLCD is the struct that wraps a Phidget LCD display
type PhidgetLCD struct {
	phidget
	handle C.PhidgetLCDHandle
}

// Create creates a PhidgetLCD handle
func (p *PhidgetLCD) Create() {
	C.PhidgetLCD_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// WriteText writes text to the display at the given font, x, and y position.
// Call Flush afterward (or enable AutoFlush) to push the framebuffer to the screen.
func (p *PhidgetLCD) WriteText(font C.PhidgetLCD_Font, x, y int, text string) error {
	str := C.CString(text)
	defer C.free(unsafe.Pointer(str))
	return p.phidgetError(C.PhidgetLCD_writeText(p.handle, font, C.int(x), C.int(y), str))
}

// SetText is a convenience wrapper that writes text at position (0, 0) using FONT_6x12 and flushes.
func (p *PhidgetLCD) SetText(text string) error {
	if err := p.WriteText(C.FONT_6x12, 0, 0, text); err != nil {
		return err
	}
	return p.Flush()
}

// Flush pushes the current framebuffer contents to the screen.
func (p *PhidgetLCD) Flush() error {
	return p.phidgetError(C.PhidgetLCD_flush(p.handle))
}

// Clear clears the display framebuffer. Call Flush to apply.
func (p *PhidgetLCD) Clear() error {
	return p.phidgetError(C.PhidgetLCD_clear(p.handle))
}

// SetBacklight sets the backlight brightness (0.0–1.0).
func (p *PhidgetLCD) SetBacklight(brightness float64) error {
	return p.phidgetError(C.PhidgetLCD_setBacklight(p.handle, C.double(brightness)))
}

// GetBacklight returns the current backlight brightness.
func (p *PhidgetLCD) GetBacklight() (float64, error) {
	var r C.double
	if cerr := C.PhidgetLCD_getBacklight(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetContrast sets the display contrast (0.0–1.0).
func (p *PhidgetLCD) SetContrast(contrast float64) error {
	return p.phidgetError(C.PhidgetLCD_setContrast(p.handle, C.double(contrast)))
}

// GetContrast returns the current display contrast.
func (p *PhidgetLCD) GetContrast() (float64, error) {
	var r C.double
	if cerr := C.PhidgetLCD_getContrast(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return float64(r), nil
}

// SetAutoFlush enables or disables automatic flushing after each draw/write call.
func (p *PhidgetLCD) SetAutoFlush(autoFlush bool) error {
	return p.phidgetError(C.PhidgetLCD_setAutoFlush(p.handle, boolToCInt(autoFlush)))
}

// GetAutoFlush returns whether automatic flushing is enabled.
func (p *PhidgetLCD) GetAutoFlush() (bool, error) {
	var r C.int
	if cerr := C.PhidgetLCD_getAutoFlush(p.handle, &r); cerr != C.EPHIDGET_OK {
		return false, p.phidgetError(cerr)
	}
	return r != 0, nil
}

// SetCursorOn shows or hides the cursor.
func (p *PhidgetLCD) SetCursorOn(on bool) error {
	return p.phidgetError(C.PhidgetLCD_setCursorOn(p.handle, boolToCInt(on)))
}

// SetCursorBlink enables or disables cursor blinking.
func (p *PhidgetLCD) SetCursorBlink(blink bool) error {
	return p.phidgetError(C.PhidgetLCD_setCursorBlink(p.handle, boolToCInt(blink)))
}

// SetSleeping puts the display to sleep or wakes it.
func (p *PhidgetLCD) SetSleeping(sleeping bool) error {
	return p.phidgetError(C.PhidgetLCD_setSleeping(p.handle, boolToCInt(sleeping)))
}

// GetWidth returns the display width in pixels.
func (p *PhidgetLCD) GetWidth() (int, error) {
	var r C.int
	if cerr := C.PhidgetLCD_getWidth(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// GetHeight returns the display height in pixels.
func (p *PhidgetLCD) GetHeight() (int, error) {
	var r C.int
	if cerr := C.PhidgetLCD_getHeight(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// DrawPixel draws a single pixel at (x, y). Use PhidgetLCD_Pixel_ON or PhidgetLCD_Pixel_OFF.
func (p *PhidgetLCD) DrawPixel(x, y int, pixelState C.PhidgetLCD_PixelState) error {
	return p.phidgetError(C.PhidgetLCD_drawPixel(p.handle, C.int(x), C.int(y), pixelState))
}

// DrawLine draws a line from (x1, y1) to (x2, y2).
func (p *PhidgetLCD) DrawLine(x1, y1, x2, y2 int) error {
	return p.phidgetError(C.PhidgetLCD_drawLine(p.handle, C.int(x1), C.int(y1), C.int(x2), C.int(y2)))
}

// DrawRect draws a rectangle. Set filled=true to fill, eraseInside=true to erase the interior.
func (p *PhidgetLCD) DrawRect(x1, y1, x2, y2 int, filled, eraseInside bool) error {
	return p.phidgetError(C.PhidgetLCD_drawRect(p.handle,
		C.int(x1), C.int(y1), C.int(x2), C.int(y2),
		boolToCInt(filled), boolToCInt(eraseInside)))
}

// Initialize resets the display to its default state.
func (p *PhidgetLCD) Initialize() error {
	return p.phidgetError(C.PhidgetLCD_initialize(p.handle))
}

// SetScreenSize sets the screen size (use PhidgetLCD_ScreenSize constants).
func (p *PhidgetLCD) SetScreenSize(size C.PhidgetLCD_ScreenSize) error {
	return p.phidgetError(C.PhidgetLCD_setScreenSize(p.handle, size))
}

// GetScreenSize returns the current screen size.
func (p *PhidgetLCD) GetScreenSize() (C.PhidgetLCD_ScreenSize, error) {
	var r C.PhidgetLCD_ScreenSize
	if cerr := C.PhidgetLCD_getScreenSize(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return r, nil
}

// Close closes the handle and deletes it.
func (p *PhidgetLCD) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetLCD_delete(&p.handle))
}

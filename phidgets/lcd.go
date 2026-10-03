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
func (p *PhidgetLCD) WriteText(font LCDFont, x, y int, text string) error {
	str := C.CString(text)
	defer C.free(unsafe.Pointer(str))
	return p.phidgetError(C.PhidgetLCD_writeText(p.handle, C.PhidgetLCD_Font(font), C.int(x), C.int(y), str))
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
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLCD_getBacklight(p.handle, r) })
}

// SetContrast sets the display contrast (0.0–1.0).
func (p *PhidgetLCD) SetContrast(contrast float64) error {
	return p.phidgetError(C.PhidgetLCD_setContrast(p.handle, C.double(contrast)))
}

// GetContrast returns the current display contrast.
func (p *PhidgetLCD) GetContrast() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLCD_getContrast(p.handle, r) })
}

// SetAutoFlush enables or disables automatic flushing after each draw/write call.
func (p *PhidgetLCD) SetAutoFlush(autoFlush bool) error {
	return p.phidgetError(C.PhidgetLCD_setAutoFlush(p.handle, boolToCInt(autoFlush)))
}

// GetAutoFlush returns whether automatic flushing is enabled.
func (p *PhidgetLCD) GetAutoFlush() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetLCD_getAutoFlush(p.handle, r) })
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
	return getInt(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetLCD_getWidth(p.handle, r) })
}

// GetHeight returns the display height in pixels.
func (p *PhidgetLCD) GetHeight() (int, error) {
	return getInt(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetLCD_getHeight(p.handle, r) })
}

// DrawPixel draws a single pixel at (x, y). Use LCDPixelOn, LCDPixelOff or LCDPixelInvert.
func (p *PhidgetLCD) DrawPixel(x, y int, pixelState LCDPixelState) error {
	return p.phidgetError(C.PhidgetLCD_drawPixel(p.handle, C.int(x), C.int(y), C.PhidgetLCD_PixelState(pixelState)))
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

// SetScreenSize sets the screen size (use the LCDScreen constants).
func (p *PhidgetLCD) SetScreenSize(size LCDScreenSize) error {
	return p.phidgetError(C.PhidgetLCD_setScreenSize(p.handle, C.PhidgetLCD_ScreenSize(size)))
}

// GetScreenSize returns the current screen size.
func (p *PhidgetLCD) GetScreenSize() (LCDScreenSize, error) {
	return get(&p.phidget, func(r *C.PhidgetLCD_ScreenSize) C.PhidgetReturnCode { return C.PhidgetLCD_getScreenSize(p.handle, r) }, func(r C.PhidgetLCD_ScreenSize) LCDScreenSize { return LCDScreenSize(r) })
}

// GetMinBacklight returns the minimum backlight brightness.
func (p *PhidgetLCD) GetMinBacklight() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLCD_getMinBacklight(p.handle, r) })
}

// GetMaxBacklight returns the maximum backlight brightness.
func (p *PhidgetLCD) GetMaxBacklight() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLCD_getMaxBacklight(p.handle, r) })
}

// GetMinContrast returns the minimum contrast.
func (p *PhidgetLCD) GetMinContrast() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLCD_getMinContrast(p.handle, r) })
}

// GetMaxContrast returns the maximum contrast.
func (p *PhidgetLCD) GetMaxContrast() (float64, error) {
	return getDouble(&p.phidget, func(r *C.double) C.PhidgetReturnCode { return C.PhidgetLCD_getMaxContrast(p.handle, r) })
}

// GetCursorOn returns whether the cursor is visible.
func (p *PhidgetLCD) GetCursorOn() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetLCD_getCursorOn(p.handle, r) })
}

// GetCursorBlink returns whether cursor blinking is enabled.
func (p *PhidgetLCD) GetCursorBlink() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetLCD_getCursorBlink(p.handle, r) })
}

// GetSleeping returns whether the display is sleeping.
func (p *PhidgetLCD) GetSleeping() (bool, error) {
	return getBool(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetLCD_getSleeping(p.handle, r) })
}

// GetFrameBuffer returns the frame buffer that is currently used for drawing.
func (p *PhidgetLCD) GetFrameBuffer() (int, error) {
	return getInt(&p.phidget, func(r *C.int) C.PhidgetReturnCode { return C.PhidgetLCD_getFrameBuffer(p.handle, r) })
}

// SetFrameBuffer selects the frame buffer to use for drawing.
func (p *PhidgetLCD) SetFrameBuffer(frameBuffer int) error {
	return p.phidgetError(C.PhidgetLCD_setFrameBuffer(p.handle, C.int(frameBuffer)))
}

// SaveFrameBuffer saves the specified frame buffer to flash memory.
func (p *PhidgetLCD) SaveFrameBuffer(frameBuffer int) error {
	return p.phidgetError(C.PhidgetLCD_saveFrameBuffer(p.handle, C.int(frameBuffer)))
}

// Copy copies a rectangular region from one frame buffer to another using
// sourceFrameBuffer as the source, destFrameBuffer as the destination.
func (p *PhidgetLCD) Copy(sourceFrameBuffer, destFrameBuffer, sourceX1, sourceY1, sourceX2, sourceY2, destX, destY int, inverted bool) error {
	return p.phidgetError(C.PhidgetLCD_copy(p.handle,
		C.int(sourceFrameBuffer), C.int(destFrameBuffer),
		C.int(sourceX1), C.int(sourceY1), C.int(sourceX2), C.int(sourceY2),
		C.int(destX), C.int(destY), boolToCInt(inverted)))
}

// GetFontSize returns the current character width and height for the given font
// (an LCDFont constant).
func (p *PhidgetLCD) GetFontSize(font LCDFont) (width, height int, err error) {
	var cw, ch C.int
	if cerr := C.PhidgetLCD_getFontSize(p.handle, C.PhidgetLCD_Font(font), &cw, &ch); cerr != C.EPHIDGET_OK {
		return 0, 0, p.phidgetError(cerr)
	}
	return int(cw), int(ch), nil
}

// SetFontSize sets the character width and height to use with the given font
// (an LCDFont constant).
func (p *PhidgetLCD) SetFontSize(font LCDFont, width, height int) error {
	return p.phidgetError(C.PhidgetLCD_setFontSize(p.handle, C.PhidgetLCD_Font(font), C.int(width), C.int(height)))
}

// WriteBitmap writes a bitmap (bitmap) to the screen at the given position
// with the specified dimensions.
func (p *PhidgetLCD) WriteBitmap(x, y, xSize, ySize int, bitmap []byte) error {
	if len(bitmap) == 0 {
		return p.phidgetError(C.EPHIDGET_INVALIDARG)
	}
	ptr := (*C.uint8_t)(unsafe.Pointer(&bitmap[0]))
	return p.phidgetError(C.PhidgetLCD_writeBitmap(p.handle, C.int(x), C.int(y), C.int(xSize), C.int(ySize), ptr))
}

// SetCharacterBitmap sets the bitmap for a single ASCII character in the
// specified font (an LCDFont constant).
func (p *PhidgetLCD) SetCharacterBitmap(font LCDFont, character byte, bitmap []byte) error {
	if len(bitmap) == 0 {
		return p.phidgetError(C.EPHIDGET_INVALIDARG)
	}
	cchar := C.CString(string(rune(character)))
	defer C.free(unsafe.Pointer(cchar))
	ptr := (*C.uint8_t)(unsafe.Pointer(&bitmap[0]))
	return p.phidgetError(C.PhidgetLCD_setCharacterBitmap(p.handle, C.PhidgetLCD_Font(font), cchar, ptr))
}

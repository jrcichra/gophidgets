package phidgets

/*
#include <phidget22.h>
*/
import "C"

// LCDFont values, for use with the methods that take them as int.
const (
	LCDFontUser1 = C.FONT_User1
	LCDFontUser2 = C.FONT_User2
	LCDFont6x10  = C.FONT_6x10
	LCDFont5x8   = C.FONT_5x8
	LCDFont6x12  = C.FONT_6x12
)

// LCDPixel values, for use with the methods that take them as int.
const (
	LCDPixelOff    = C.PIXEL_STATE_OFF
	LCDPixelOn     = C.PIXEL_STATE_ON
	LCDPixelInvert = C.PIXEL_STATE_INVERT
)

// LCDScreen values, for use with the methods that take them as int.
const (
	LCDScreenNone   = C.SCREEN_SIZE_NONE
	LCDScreen1x8    = C.SCREEN_SIZE_1x8
	LCDScreen2x8    = C.SCREEN_SIZE_2x8
	LCDScreen1x16   = C.SCREEN_SIZE_1x16
	LCDScreen2x16   = C.SCREEN_SIZE_2x16
	LCDScreen4x16   = C.SCREEN_SIZE_4x16
	LCDScreen2x20   = C.SCREEN_SIZE_2x20
	LCDScreen4x20   = C.SCREEN_SIZE_4x20
	LCDScreen2x24   = C.SCREEN_SIZE_2x24
	LCDScreen1x40   = C.SCREEN_SIZE_1x40
	LCDScreen2x40   = C.SCREEN_SIZE_2x40
	LCDScreen4x40   = C.SCREEN_SIZE_4x40
	LCDScreen64x128 = C.SCREEN_SIZE_64x128
)

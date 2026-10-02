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

// SpatialAlgorithm values, for use with the methods that take them as int.
const (
	SpatialAlgorithmNone = C.SPATIAL_ALGORITHM_NONE
	SpatialAlgorithmAHRS = C.SPATIAL_ALGORITHM_AHRS
	SpatialAlgorithmIMU  = C.SPATIAL_ALGORITHM_IMU
)

// VoltageOutputRange values, for use with the methods that take them as int.
const (
	VoltageOutputRange10V = C.VOLTAGE_OUTPUT_RANGE_10V
	VoltageOutputRange5V  = C.VOLTAGE_OUTPUT_RANGE_5V
)

// PositionType values, for use with the motor velocity controller
// SetPositionType/GetPositionType.
const (
	PositionTypeEncoder    = C.POSITION_TYPE_ENCODER
	PositionTypeHallSensor = C.POSITION_TYPE_HALL_SENSOR
)

// RFIDProtocol values, for use with the methods that take them as int.
const (
	RFIDProtocolEM4100      = C.PROTOCOL_EM4100
	RFIDProtocolISO11785FDX = C.PROTOCOL_ISO11785_FDX_B
	RFIDProtocolPhidget     = C.PROTOCOL_PHIDGETS
)

// ErrorEvent values, for matching the code delivered to the
// SetOnErrorHandler callback (EEPHIDGET_*).
const (
	ErrorEventOK             = C.EEPHIDGET_OK
	ErrorEventOverrun        = C.EEPHIDGET_OVERRUN
	ErrorEventPacketLost     = C.EEPHIDGET_PACKETLOST
	ErrorEventWrap           = C.EEPHIDGET_WRAP
	ErrorEventOverTemp       = C.EEPHIDGET_OVERTEMP
	ErrorEventOverCurrent    = C.EEPHIDGET_OVERCURRENT
	ErrorEventOutOfRange     = C.EEPHIDGET_OUTOFRANGE
	ErrorEventBadPower       = C.EEPHIDGET_BADPOWER
	ErrorEventSaturation     = C.EEPHIDGET_SATURATION
	ErrorEventOverVoltage    = C.EEPHIDGET_OVERVOLTAGE
	ErrorEventFailsafe       = C.EEPHIDGET_FAILSAFE
	ErrorEventVoltageError   = C.EEPHIDGET_VOLTAGEERROR
	ErrorEventEnergyDump     = C.EEPHIDGET_ENERGYDUMP
	ErrorEventMotorStall     = C.EEPHIDGET_MOTORSTALL
	ErrorEventInvalidState   = C.EEPHIDGET_INVALIDSTATE
	ErrorEventBadConnection  = C.EEPHIDGET_BADCONNECTION
	ErrorEventOutOfRangeHigh = C.EEPHIDGET_OUTOFRANGEHIGH
	ErrorEventOutOfRangeLow  = C.EEPHIDGET_OUTOFRANGELOW
	ErrorEventFault          = C.EEPHIDGET_FAULT
	ErrorEventEStop          = C.EEPHIDGET_ESTOP
	ErrorEventBadCurrent     = C.EEPHIDGET_BADCURRENT
)

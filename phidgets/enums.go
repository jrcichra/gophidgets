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

// FanMode values, for use with the methods that take them as int.
const (
	FanModeOff  = C.FAN_MODE_OFF
	FanModeOn   = C.FAN_MODE_ON
	FanModeAuto = C.FAN_MODE_AUTO
)

// HubPortSpeed values for PhidgetHubPortSpeed / VINT port-speed methods (bps).
const (
	HubPortSpeedAuto = C.PHIDGET_HUBPORTSPEED_AUTO
)

// HubPortMode values, for use with the PhidgetHub port-mode methods.
const (
	HubPortModeVintPort          = C.PORT_MODE_VINT_PORT
	HubPortModeDigitalInput      = C.PORT_MODE_DIGITAL_INPUT
	HubPortModeDigitalOutput     = C.PORT_MODE_DIGITAL_OUTPUT
	HubPortModeVoltageInput      = C.PORT_MODE_VOLTAGE_INPUT
	HubPortModeVoltageRatioInput = C.PORT_MODE_VOLTAGE_RATIO_INPUT
)

// IREncoding values, for use with the IRCodeInfo struct.
const (
	IREncodingUnknown = C.IR_ENCODING_UNKNOWN
	IREncodingSpace   = C.IR_ENCODING_SPACE
	IREncodingPulse   = C.IR_ENCODING_PULSE
	IREncodingBiPhase = C.IR_ENCODING_BIPHASE
	IREncodingRC5     = C.IR_ENCODING_RC5
	IREncodingRC6     = C.IR_ENCODING_RC6
)

// IRLength values, for use with the IRCodeInfo struct.
const (
	IRLengthUnknown  = C.IR_LENGTH_UNKNOWN
	IRLengthConstant = C.IR_LENGTH_CONSTANT
	IRLengthVariable = C.IR_LENGTH_VARIABLE
)

// LEDColorOrder values, for use with the PhidgetLEDArray color-order methods.
const (
	LEDColorOrderRGB  = C.LED_COLOR_ORDER_RGB
	LEDColorOrderGRB  = C.LED_COLOR_ORDER_GRB
	LEDColorOrderRGBW = C.LED_COLOR_ORDER_RGBW
	LEDColorOrderGRBW = C.LED_COLOR_ORDER_GRBW
)

// LEDAnimationType values, for use with the LEDAnimation struct.
const (
	LEDAnimationForwardScroll       = C.ANIMATION_TYPE_FORWARD_SCROLL
	LEDAnimationReverseScroll       = C.ANIMATION_TYPE_REVERSE_SCROLL
	LEDAnimationRandomize           = C.ANIMATION_TYPE_RANDOMIZE
	LEDAnimationForwardScrollMirror = C.ANIMATION_TYPE_FORWARD_SCROLL_MIRROR
	LEDAnimationReverseScrollMirror = C.ANIMATION_TYPE_REVERSE_SCROLL_MIRROR
)

// DataAdapterVoltage values, for use with the PhidgetDataAdapter voltage methods.
const (
	DataAdapterVoltageExtern = C.DATAADAPTER_VOLTAGE_EXTERN
	DataAdapterVoltage2_5V   = C.DATAADAPTER_VOLTAGE_2_5V
	DataAdapterVoltage3_3V   = C.DATAADAPTER_VOLTAGE_3_3V
	DataAdapterVoltage5_0V   = C.DATAADAPTER_VOLTAGE_5_0V
)

// DataAdapterEndianness values, for use with the PhidgetDataAdapter endianness methods.
const (
	DataAdapterEndiannessMSBFirst = C.ENDIANNESS_MSB_FIRST
	DataAdapterEndiannessLSBFirst = C.ENDIANNESS_LSB_FIRST
)

// DataAdapterFrequency values, for use with the PhidgetDataAdapter frequency methods (Hz).
const (
	DataAdapterFrequency10kHz   = C.FREQUENCY_10kHz
	DataAdapterFrequency100kHz  = C.FREQUENCY_100kHz
	DataAdapterFrequency400kHz  = C.FREQUENCY_400kHz
	DataAdapterFrequency188kHz  = C.FREQUENCY_188kHz
	DataAdapterFrequency375kHz  = C.FREQUENCY_375kHz
	DataAdapterFrequency750kHz  = C.FREQUENCY_750kHz
	DataAdapterFrequency1500kHz = C.FREQUENCY_1500kHz
	DataAdapterFrequency3MHz    = C.FREQUENCY_3MHz
	DataAdapterFrequency6MHz    = C.FREQUENCY_6MHz
)

// DataAdapterSPIChipSelect values, for use with the PhidgetDataAdapter chip-select methods.
const (
	DataAdapterSPIChipSelectActiveLow  = C.SPI_CHIP_SELECT_ACTIVE_LOW
	DataAdapterSPIChipSelectActiveHigh = C.SPI_CHIP_SELECT_ACTIVE_HIGH
	DataAdapterSPIChipSelectLow        = C.SPI_CHIP_SELECT_LOW
	DataAdapterSPIChipSelectHigh       = C.SPI_CHIP_SELECT_HIGH
)

// DataAdapterSPIMode values, for use with the PhidgetDataAdapter SPI-mode methods.
const (
	DataAdapterSPIMode0 = C.SPI_MODE_0
	DataAdapterSPIMode1 = C.SPI_MODE_1
	DataAdapterSPIMode2 = C.SPI_MODE_2
	DataAdapterSPIMode3 = C.SPI_MODE_3
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

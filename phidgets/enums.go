package phidgets

/*
#include <phidget22.h>
*/
import "C"

// LCDFont is a libphidget22 enum.
type LCDFont int

// LCDFont values.
const (
	LCDFontUser1 LCDFont = C.FONT_User1
	LCDFontUser2 LCDFont = C.FONT_User2
	LCDFont6x10  LCDFont = C.FONT_6x10
	LCDFont5x8   LCDFont = C.FONT_5x8
	LCDFont6x12  LCDFont = C.FONT_6x12
)

// LCDPixelState is a libphidget22 enum.
type LCDPixelState int

// LCDPixelState values.
const (
	LCDPixelOff    LCDPixelState = C.PIXEL_STATE_OFF
	LCDPixelOn     LCDPixelState = C.PIXEL_STATE_ON
	LCDPixelInvert LCDPixelState = C.PIXEL_STATE_INVERT
)

// LCDScreenSize is a libphidget22 enum.
type LCDScreenSize int

// LCDScreenSize values.
const (
	LCDScreenNone   LCDScreenSize = C.SCREEN_SIZE_NONE
	LCDScreen1x8    LCDScreenSize = C.SCREEN_SIZE_1x8
	LCDScreen2x8    LCDScreenSize = C.SCREEN_SIZE_2x8
	LCDScreen1x16   LCDScreenSize = C.SCREEN_SIZE_1x16
	LCDScreen2x16   LCDScreenSize = C.SCREEN_SIZE_2x16
	LCDScreen4x16   LCDScreenSize = C.SCREEN_SIZE_4x16
	LCDScreen2x20   LCDScreenSize = C.SCREEN_SIZE_2x20
	LCDScreen4x20   LCDScreenSize = C.SCREEN_SIZE_4x20
	LCDScreen2x24   LCDScreenSize = C.SCREEN_SIZE_2x24
	LCDScreen1x40   LCDScreenSize = C.SCREEN_SIZE_1x40
	LCDScreen2x40   LCDScreenSize = C.SCREEN_SIZE_2x40
	LCDScreen4x40   LCDScreenSize = C.SCREEN_SIZE_4x40
	LCDScreen64x128 LCDScreenSize = C.SCREEN_SIZE_64x128
)

// SpatialAlgorithm is a libphidget22 enum.
type SpatialAlgorithm int

// SpatialAlgorithm values.
const (
	SpatialAlgorithmNone SpatialAlgorithm = C.SPATIAL_ALGORITHM_NONE
	SpatialAlgorithmAHRS SpatialAlgorithm = C.SPATIAL_ALGORITHM_AHRS
	SpatialAlgorithmIMU  SpatialAlgorithm = C.SPATIAL_ALGORITHM_IMU
)

// VoltageOutputRange is a libphidget22 enum.
type VoltageOutputRange int

// VoltageOutputRange values.
const (
	VoltageOutputRange10V VoltageOutputRange = C.VOLTAGE_OUTPUT_RANGE_10V
	VoltageOutputRange5V  VoltageOutputRange = C.VOLTAGE_OUTPUT_RANGE_5V
)

// PositionType is a libphidget22 enum.
type PositionType int

// PositionType values, for use with the motor velocity controller
// SetPositionType/GetPositionType.
const (
	PositionTypeEncoder    PositionType = C.POSITION_TYPE_ENCODER
	PositionTypeHallSensor PositionType = C.POSITION_TYPE_HALL_SENSOR
)

// RFIDProtocol is a libphidget22 enum.
type RFIDProtocol int

// RFIDProtocol values.
const (
	RFIDProtocolEM4100      RFIDProtocol = C.PROTOCOL_EM4100
	RFIDProtocolISO11785FDX RFIDProtocol = C.PROTOCOL_ISO11785_FDX_B
	RFIDProtocolPhidget     RFIDProtocol = C.PROTOCOL_PHIDGETS
)

// FanMode is a libphidget22 enum.
type FanMode int

// FanMode values.
const (
	FanModeOff  FanMode = C.FAN_MODE_OFF
	FanModeOn   FanMode = C.FAN_MODE_ON
	FanModeAuto FanMode = C.FAN_MODE_AUTO
)

// HubPortSpeed values for PhidgetHubPortSpeed / VINT port-speed methods (bps).
const (
	HubPortSpeedAuto = C.PHIDGET_HUBPORTSPEED_AUTO
)

// HubPortMode is a libphidget22 enum.
type HubPortMode int

// HubPortMode values, for use with the PhidgetHub port-mode methods.
const (
	HubPortModeVintPort          HubPortMode = C.PORT_MODE_VINT_PORT
	HubPortModeDigitalInput      HubPortMode = C.PORT_MODE_DIGITAL_INPUT
	HubPortModeDigitalOutput     HubPortMode = C.PORT_MODE_DIGITAL_OUTPUT
	HubPortModeVoltageInput      HubPortMode = C.PORT_MODE_VOLTAGE_INPUT
	HubPortModeVoltageRatioInput HubPortMode = C.PORT_MODE_VOLTAGE_RATIO_INPUT
)

// IREncoding is a libphidget22 enum.
type IREncoding int

// IREncoding values, for use with the IRCodeInfo struct.
const (
	IREncodingUnknown IREncoding = C.IR_ENCODING_UNKNOWN
	IREncodingSpace   IREncoding = C.IR_ENCODING_SPACE
	IREncodingPulse   IREncoding = C.IR_ENCODING_PULSE
	IREncodingBiPhase IREncoding = C.IR_ENCODING_BIPHASE
	IREncodingRC5     IREncoding = C.IR_ENCODING_RC5
	IREncodingRC6     IREncoding = C.IR_ENCODING_RC6
)

// IRLength is a libphidget22 enum.
type IRLength int

// IRLength values, for use with the IRCodeInfo struct.
const (
	IRLengthUnknown  IRLength = C.IR_LENGTH_UNKNOWN
	IRLengthConstant IRLength = C.IR_LENGTH_CONSTANT
	IRLengthVariable IRLength = C.IR_LENGTH_VARIABLE
)

// LEDColorOrder is a libphidget22 enum.
type LEDColorOrder int

// LEDColorOrder values, for use with the PhidgetLEDArray color-order methods.
const (
	LEDColorOrderRGB  LEDColorOrder = C.LED_COLOR_ORDER_RGB
	LEDColorOrderGRB  LEDColorOrder = C.LED_COLOR_ORDER_GRB
	LEDColorOrderRGBW LEDColorOrder = C.LED_COLOR_ORDER_RGBW
	LEDColorOrderGRBW LEDColorOrder = C.LED_COLOR_ORDER_GRBW
)

// LEDAnimationType is a libphidget22 enum.
type LEDAnimationType int

// LEDAnimationType values, for use with the LEDAnimation struct.
const (
	LEDAnimationForwardScroll       LEDAnimationType = C.ANIMATION_TYPE_FORWARD_SCROLL
	LEDAnimationReverseScroll       LEDAnimationType = C.ANIMATION_TYPE_REVERSE_SCROLL
	LEDAnimationRandomize           LEDAnimationType = C.ANIMATION_TYPE_RANDOMIZE
	LEDAnimationForwardScrollMirror LEDAnimationType = C.ANIMATION_TYPE_FORWARD_SCROLL_MIRROR
	LEDAnimationReverseScrollMirror LEDAnimationType = C.ANIMATION_TYPE_REVERSE_SCROLL_MIRROR
)

// DataAdapterVoltage is a libphidget22 enum.
type DataAdapterVoltage int

// DataAdapterVoltage values, for use with the PhidgetDataAdapter voltage methods.
const (
	DataAdapterVoltageExtern DataAdapterVoltage = C.DATAADAPTER_VOLTAGE_EXTERN
	DataAdapterVoltage2_5V   DataAdapterVoltage = C.DATAADAPTER_VOLTAGE_2_5V
	DataAdapterVoltage3_3V   DataAdapterVoltage = C.DATAADAPTER_VOLTAGE_3_3V
	DataAdapterVoltage5_0V   DataAdapterVoltage = C.DATAADAPTER_VOLTAGE_5_0V
)

// DataAdapterEndianness is a libphidget22 enum.
type DataAdapterEndianness int

// DataAdapterEndianness values, for use with the PhidgetDataAdapter endianness methods.
const (
	DataAdapterEndiannessMSBFirst DataAdapterEndianness = C.ENDIANNESS_MSB_FIRST
	DataAdapterEndiannessLSBFirst DataAdapterEndianness = C.ENDIANNESS_LSB_FIRST
)

// DataAdapterFrequency is a libphidget22 enum.
type DataAdapterFrequency int

// DataAdapterFrequency values, for use with the PhidgetDataAdapter frequency methods (Hz).
const (
	DataAdapterFrequency10kHz   DataAdapterFrequency = C.FREQUENCY_10kHz
	DataAdapterFrequency100kHz  DataAdapterFrequency = C.FREQUENCY_100kHz
	DataAdapterFrequency400kHz  DataAdapterFrequency = C.FREQUENCY_400kHz
	DataAdapterFrequency188kHz  DataAdapterFrequency = C.FREQUENCY_188kHz
	DataAdapterFrequency375kHz  DataAdapterFrequency = C.FREQUENCY_375kHz
	DataAdapterFrequency750kHz  DataAdapterFrequency = C.FREQUENCY_750kHz
	DataAdapterFrequency1500kHz DataAdapterFrequency = C.FREQUENCY_1500kHz
	DataAdapterFrequency3MHz    DataAdapterFrequency = C.FREQUENCY_3MHz
	DataAdapterFrequency6MHz    DataAdapterFrequency = C.FREQUENCY_6MHz
)

// DataAdapterSPIChipSelect is a libphidget22 enum.
type DataAdapterSPIChipSelect int

// DataAdapterSPIChipSelect values, for use with the PhidgetDataAdapter chip-select methods.
const (
	DataAdapterSPIChipSelectActiveLow  DataAdapterSPIChipSelect = C.SPI_CHIP_SELECT_ACTIVE_LOW
	DataAdapterSPIChipSelectActiveHigh DataAdapterSPIChipSelect = C.SPI_CHIP_SELECT_ACTIVE_HIGH
	DataAdapterSPIChipSelectLow        DataAdapterSPIChipSelect = C.SPI_CHIP_SELECT_LOW
	DataAdapterSPIChipSelectHigh       DataAdapterSPIChipSelect = C.SPI_CHIP_SELECT_HIGH
)

// DataAdapterSPIMode is a libphidget22 enum.
type DataAdapterSPIMode int

// DataAdapterSPIMode values, for use with the PhidgetDataAdapter SPI-mode methods.
const (
	DataAdapterSPIMode0 DataAdapterSPIMode = C.SPI_MODE_0
	DataAdapterSPIMode1 DataAdapterSPIMode = C.SPI_MODE_1
	DataAdapterSPIMode2 DataAdapterSPIMode = C.SPI_MODE_2
	DataAdapterSPIMode3 DataAdapterSPIMode = C.SPI_MODE_3
)

// ErrorEvent is a libphidget22 enum.
type ErrorEvent int

// ErrorEvent values, for matching the code delivered to the
// SetOnErrorHandler callback (EEPHIDGET_*).
const (
	ErrorEventOK             ErrorEvent = C.EEPHIDGET_OK
	ErrorEventOverrun        ErrorEvent = C.EEPHIDGET_OVERRUN
	ErrorEventPacketLost     ErrorEvent = C.EEPHIDGET_PACKETLOST
	ErrorEventWrap           ErrorEvent = C.EEPHIDGET_WRAP
	ErrorEventOverTemp       ErrorEvent = C.EEPHIDGET_OVERTEMP
	ErrorEventOverCurrent    ErrorEvent = C.EEPHIDGET_OVERCURRENT
	ErrorEventOutOfRange     ErrorEvent = C.EEPHIDGET_OUTOFRANGE
	ErrorEventBadPower       ErrorEvent = C.EEPHIDGET_BADPOWER
	ErrorEventSaturation     ErrorEvent = C.EEPHIDGET_SATURATION
	ErrorEventOverVoltage    ErrorEvent = C.EEPHIDGET_OVERVOLTAGE
	ErrorEventFailsafe       ErrorEvent = C.EEPHIDGET_FAILSAFE
	ErrorEventVoltageError   ErrorEvent = C.EEPHIDGET_VOLTAGEERROR
	ErrorEventEnergyDump     ErrorEvent = C.EEPHIDGET_ENERGYDUMP
	ErrorEventMotorStall     ErrorEvent = C.EEPHIDGET_MOTORSTALL
	ErrorEventInvalidState   ErrorEvent = C.EEPHIDGET_INVALIDSTATE
	ErrorEventBadConnection  ErrorEvent = C.EEPHIDGET_BADCONNECTION
	ErrorEventOutOfRangeHigh ErrorEvent = C.EEPHIDGET_OUTOFRANGEHIGH
	ErrorEventOutOfRangeLow  ErrorEvent = C.EEPHIDGET_OUTOFRANGELOW
	ErrorEventFault          ErrorEvent = C.EEPHIDGET_FAULT
	ErrorEventEStop          ErrorEvent = C.EEPHIDGET_ESTOP
	ErrorEventBadCurrent     ErrorEvent = C.EEPHIDGET_BADCURRENT
)

// DeviceClass is the class of a Phidget device (PHIDCLASS_*).
type DeviceClass int

// DeviceClass values.
const (
	DeviceClassNone              DeviceClass = C.PHIDCLASS_NOTHING
	DeviceClassAccelerometer     DeviceClass = C.PHIDCLASS_ACCELEROMETER
	DeviceClassAdvancedServo     DeviceClass = C.PHIDCLASS_ADVANCEDSERVO
	DeviceClassAnalog            DeviceClass = C.PHIDCLASS_ANALOG
	DeviceClassBridge            DeviceClass = C.PHIDCLASS_BRIDGE
	DeviceClassCurrentInput      DeviceClass = C.PHIDCLASS_CURRENTINPUT
	DeviceClassCurrentOutput     DeviceClass = C.PHIDCLASS_CURRENTOUTPUT
	DeviceClassDataAdapter       DeviceClass = C.PHIDCLASS_DATAADAPTER
	DeviceClassDictionary        DeviceClass = C.PHIDCLASS_DICTIONARY
	DeviceClassEncoder           DeviceClass = C.PHIDCLASS_ENCODER
	DeviceClassFirmwareUpgrade   DeviceClass = C.PHIDCLASS_FIRMWAREUPGRADE
	DeviceClassFrequencyCounter  DeviceClass = C.PHIDCLASS_FREQUENCYCOUNTER
	DeviceClassFrequencyOutput   DeviceClass = C.PHIDCLASS_FREQUENCYOUTPUT
	DeviceClassGeneric           DeviceClass = C.PHIDCLASS_GENERIC
	DeviceClassGPS               DeviceClass = C.PHIDCLASS_GPS
	DeviceClassHub               DeviceClass = C.PHIDCLASS_HUB
	DeviceClassInterfaceKit      DeviceClass = C.PHIDCLASS_INTERFACEKIT
	DeviceClassIR                DeviceClass = C.PHIDCLASS_IR
	DeviceClassLED               DeviceClass = C.PHIDCLASS_LED
	DeviceClassLEDArray          DeviceClass = C.PHIDCLASS_LEDARRAY
	DeviceClassMotorControl      DeviceClass = C.PHIDCLASS_MOTORCONTROL
	DeviceClassPHSensor          DeviceClass = C.PHIDCLASS_PHSENSOR
	DeviceClassRFID              DeviceClass = C.PHIDCLASS_RFID
	DeviceClassServo             DeviceClass = C.PHIDCLASS_SERVO
	DeviceClassSpatial           DeviceClass = C.PHIDCLASS_SPATIAL
	DeviceClassStepper           DeviceClass = C.PHIDCLASS_STEPPER
	DeviceClassTemperatureSensor DeviceClass = C.PHIDCLASS_TEMPERATURESENSOR
	DeviceClassTextLCD           DeviceClass = C.PHIDCLASS_TEXTLCD
	DeviceClassVINT              DeviceClass = C.PHIDCLASS_VINT
	DeviceClassVoltageInput      DeviceClass = C.PHIDCLASS_VOLTAGEINPUT
)

// ChannelClass is the class of a Phidget channel (PHIDCHCLASS_*).
type ChannelClass int

// ChannelClass values.
const (
	ChannelClassNone                    ChannelClass = C.PHIDCHCLASS_NOTHING
	ChannelClassAccelerometer           ChannelClass = C.PHIDCHCLASS_ACCELEROMETER
	ChannelClassBLDCMotor               ChannelClass = C.PHIDCHCLASS_BLDCMOTOR
	ChannelClassCapacitiveTouch         ChannelClass = C.PHIDCHCLASS_CAPACITIVETOUCH
	ChannelClassCurrentInput            ChannelClass = C.PHIDCHCLASS_CURRENTINPUT
	ChannelClassCurrentOutput           ChannelClass = C.PHIDCHCLASS_CURRENTOUTPUT
	ChannelClassDataAdapter             ChannelClass = C.PHIDCHCLASS_DATAADAPTER
	ChannelClassDCMotor                 ChannelClass = C.PHIDCHCLASS_DCMOTOR
	ChannelClassDictionary              ChannelClass = C.PHIDCHCLASS_DICTIONARY
	ChannelClassDigitalInput            ChannelClass = C.PHIDCHCLASS_DIGITALINPUT
	ChannelClassDigitalOutput           ChannelClass = C.PHIDCHCLASS_DIGITALOUTPUT
	ChannelClassDistanceSensor          ChannelClass = C.PHIDCHCLASS_DISTANCESENSOR
	ChannelClassEncoder                 ChannelClass = C.PHIDCHCLASS_ENCODER
	ChannelClassFirmwareUpgrade         ChannelClass = C.PHIDCHCLASS_FIRMWAREUPGRADE
	ChannelClassFrequencyCounter        ChannelClass = C.PHIDCHCLASS_FREQUENCYCOUNTER
	ChannelClassGeneric                 ChannelClass = C.PHIDCHCLASS_GENERIC
	ChannelClassGPS                     ChannelClass = C.PHIDCHCLASS_GPS
	ChannelClassGyroscope               ChannelClass = C.PHIDCHCLASS_GYROSCOPE
	ChannelClassHub                     ChannelClass = C.PHIDCHCLASS_HUB
	ChannelClassHumiditySensor          ChannelClass = C.PHIDCHCLASS_HUMIDITYSENSOR
	ChannelClassIR                      ChannelClass = C.PHIDCHCLASS_IR
	ChannelClassLCD                     ChannelClass = C.PHIDCHCLASS_LCD
	ChannelClassLEDArray                ChannelClass = C.PHIDCHCLASS_LEDARRAY
	ChannelClassLightSensor             ChannelClass = C.PHIDCHCLASS_LIGHTSENSOR
	ChannelClassMagnetometer            ChannelClass = C.PHIDCHCLASS_MAGNETOMETER
	ChannelClassMotorPositionController ChannelClass = C.PHIDCHCLASS_MOTORPOSITIONCONTROLLER
	ChannelClassMotorVelocityController ChannelClass = C.PHIDCHCLASS_MOTORVELOCITYCONTROLLER
	ChannelClassPHSensor                ChannelClass = C.PHIDCHCLASS_PHSENSOR
	ChannelClassPowerGuard              ChannelClass = C.PHIDCHCLASS_POWERGUARD
	ChannelClassPressureSensor          ChannelClass = C.PHIDCHCLASS_PRESSURESENSOR
	ChannelClassRCServo                 ChannelClass = C.PHIDCHCLASS_RCSERVO
	ChannelClassResistanceInput         ChannelClass = C.PHIDCHCLASS_RESISTANCEINPUT
	ChannelClassRFID                    ChannelClass = C.PHIDCHCLASS_RFID
	ChannelClassSoundSensor             ChannelClass = C.PHIDCHCLASS_SOUNDSENSOR
	ChannelClassSpatial                 ChannelClass = C.PHIDCHCLASS_SPATIAL
	ChannelClassStepper                 ChannelClass = C.PHIDCHCLASS_STEPPER
	ChannelClassTemperatureSensor       ChannelClass = C.PHIDCHCLASS_TEMPERATURESENSOR
	ChannelClassVoltageInput            ChannelClass = C.PHIDCHCLASS_VOLTAGEINPUT
	ChannelClassVoltageOutput           ChannelClass = C.PHIDCHCLASS_VOLTAGEOUTPUT
	ChannelClassVoltageRatioInput       ChannelClass = C.PHIDCHCLASS_VOLTAGERATIOINPUT
)

// ChannelSubclass is the subclass of a Phidget channel (PHIDCHSUBCLASS_*).
type ChannelSubclass int

// ChannelSubclass values.
const (
	ChannelSubclassNone                          ChannelSubclass = C.PHIDCHSUBCLASS_NONE
	ChannelSubclassDigitaloutputDutyCycle        ChannelSubclass = C.PHIDCHSUBCLASS_DIGITALOUTPUT_DUTY_CYCLE
	ChannelSubclassDigitaloutputFrequency        ChannelSubclass = C.PHIDCHSUBCLASS_DIGITALOUTPUT_FREQUENCY
	ChannelSubclassDigitaloutputLedDriver        ChannelSubclass = C.PHIDCHSUBCLASS_DIGITALOUTPUT_LED_DRIVER
	ChannelSubclassEncoderModeSettable           ChannelSubclass = C.PHIDCHSUBCLASS_ENCODER_MODE_SETTABLE
	ChannelSubclassLcdGraphic                    ChannelSubclass = C.PHIDCHSUBCLASS_LCD_GRAPHIC
	ChannelSubclassLcdText                       ChannelSubclass = C.PHIDCHSUBCLASS_LCD_TEXT
	ChannelSubclassRfidNfc                       ChannelSubclass = C.PHIDCHSUBCLASS_RFID_NFC
	ChannelSubclassSpatialAhrs                   ChannelSubclass = C.PHIDCHSUBCLASS_SPATIAL_AHRS
	ChannelSubclassTemperaturesensorRtd          ChannelSubclass = C.PHIDCHSUBCLASS_TEMPERATURESENSOR_RTD
	ChannelSubclassTemperaturesensorThermocouple ChannelSubclass = C.PHIDCHSUBCLASS_TEMPERATURESENSOR_THERMOCOUPLE
	ChannelSubclassVoltageinputSensorPort        ChannelSubclass = C.PHIDCHSUBCLASS_VOLTAGEINPUT_SENSOR_PORT
	ChannelSubclassVoltageratioinputBridge       ChannelSubclass = C.PHIDCHSUBCLASS_VOLTAGERATIOINPUT_BRIDGE
	ChannelSubclassVoltageratioinputSensorPort   ChannelSubclass = C.PHIDCHSUBCLASS_VOLTAGERATIOINPUT_SENSOR_PORT
)

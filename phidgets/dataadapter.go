package phidgets

/*
#include <phidget22.h>
#include <stdlib.h>
#include "phidgets.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// PhidgetDataAdapter wraps a Phidget Data Adapter (SPI/I2C)
type PhidgetDataAdapter struct {
	phidget
	handle C.PhidgetDataAdapterHandle
}

// Create creates a PhidgetDataAdapter handle
func (p *PhidgetDataAdapter) Create() {
	C.PhidgetDataAdapter_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetDataAdapterVoltage returns the adapter I/O voltage (Phidget_DataAdapterVoltage)
func (p *PhidgetDataAdapter) GetDataAdapterVoltage() (int, error) {
	var r C.Phidget_DataAdapterVoltage
	if cerr := C.PhidgetDataAdapter_getDataAdapterVoltage(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// SetDataAdapterVoltage sets the adapter I/O voltage (Phidget_DataAdapterVoltage)
func (p *PhidgetDataAdapter) SetDataAdapterVoltage(voltage int) error {
	return p.phidgetError(C.PhidgetDataAdapter_setDataAdapterVoltage(p.handle, C.Phidget_DataAdapterVoltage(voltage)))
}

// GetDataBits returns the number of data bits
func (p *PhidgetDataAdapter) GetDataBits() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetDataAdapter_getDataBits(p.handle, r) })
}

// SetDataBits sets the number of data bits
func (p *PhidgetDataAdapter) SetDataBits(bits uint32) error {
	return p.phidgetError(C.PhidgetDataAdapter_setDataBits(p.handle, C.uint32_t(bits)))
}

// GetMinDataBits returns the minimum number of data bits
func (p *PhidgetDataAdapter) GetMinDataBits() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetDataAdapter_getMinDataBits(p.handle, r) })
}

// GetMaxDataBits returns the maximum number of data bits
func (p *PhidgetDataAdapter) GetMaxDataBits() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode { return C.PhidgetDataAdapter_getMaxDataBits(p.handle, r) })
}

// GetEndianness returns the data endianness (PhidgetDataAdapter_Endianness)
func (p *PhidgetDataAdapter) GetEndianness() (int, error) {
	var r C.PhidgetDataAdapter_Endianness
	if cerr := C.PhidgetDataAdapter_getEndianness(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// SetEndianness sets the data endianness (PhidgetDataAdapter_Endianness)
func (p *PhidgetDataAdapter) SetEndianness(order int) error {
	return p.phidgetError(C.PhidgetDataAdapter_setEndianness(p.handle, C.PhidgetDataAdapter_Endianness(order)))
}

// GetFrequency returns the bus frequency (PhidgetDataAdapter_Frequency)
func (p *PhidgetDataAdapter) GetFrequency() (int, error) {
	var r C.PhidgetDataAdapter_Frequency
	if cerr := C.PhidgetDataAdapter_getFrequency(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// SetFrequency sets the bus frequency (PhidgetDataAdapter_Frequency)
func (p *PhidgetDataAdapter) SetFrequency(freq int) error {
	return p.phidgetError(C.PhidgetDataAdapter_setFrequency(p.handle, C.PhidgetDataAdapter_Frequency(freq)))
}

// GetSPIChipSelect returns the SPI chip select configuration (PhidgetDataAdapter_SPIChipSelect)
func (p *PhidgetDataAdapter) GetSPIChipSelect() (int, error) {
	var r C.PhidgetDataAdapter_SPIChipSelect
	if cerr := C.PhidgetDataAdapter_getSPIChipSelect(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// SetSPIChipSelect sets the SPI chip select configuration (PhidgetDataAdapter_SPIChipSelect)
func (p *PhidgetDataAdapter) SetSPIChipSelect(sel int) error {
	return p.phidgetError(C.PhidgetDataAdapter_setSPIChipSelect(p.handle, C.PhidgetDataAdapter_SPIChipSelect(sel)))
}

// GetSPIMode returns the SPI mode (PhidgetDataAdapter_SPIMode)
func (p *PhidgetDataAdapter) GetSPIMode() (int, error) {
	var r C.PhidgetDataAdapter_SPIMode
	if cerr := C.PhidgetDataAdapter_getSPIMode(p.handle, &r); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// SetSPIMode sets the SPI mode (PhidgetDataAdapter_SPIMode)
func (p *PhidgetDataAdapter) SetSPIMode(mode int) error {
	return p.phidgetError(C.PhidgetDataAdapter_setSPIMode(p.handle, C.PhidgetDataAdapter_SPIMode(mode)))
}

// GetMaxReceivePacketLength returns the maximum number of bytes that can be received per packet
func (p *PhidgetDataAdapter) GetMaxReceivePacketLength() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetDataAdapter_getMaxReceivePacketLength(p.handle, r)
	})
}

// GetMaxSendPacketLength returns the maximum number of bytes that can be sent per packet
func (p *PhidgetDataAdapter) GetMaxSendPacketLength() (uint32, error) {
	return getUint32(&p.phidget, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.PhidgetDataAdapter_getMaxSendPacketLength(p.handle, r)
	})
}

// SendPacket sends data to the adapter
func (p *PhidgetDataAdapter) SendPacket(data []byte) error {
	if len(data) == 0 {
		return errors.New("gophidgets: packet data must not be empty")
	}
	ptr := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	return p.phidgetError(C.PhidgetDataAdapter_sendPacket(p.handle, ptr, C.size_t(len(data))))
}

// SendPacketWaitResponse sends data and waits for a response. maxReceive is
// the size in bytes of the response buffer; the returned slice holds the
// number of bytes actually received.
func (p *PhidgetDataAdapter) SendPacketWaitResponse(data []byte, maxReceive int) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("gophidgets: packet data must not be empty")
	}
	if maxReceive <= 0 {
		return nil, errors.New("gophidgets: receive length must be positive")
	}
	dptr := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	buf := C.malloc(C.size_t(maxReceive))
	defer C.free(buf)
	recvLen := C.size_t(maxReceive)
	cerr := C.PhidgetDataAdapter_sendPacketWaitResponse(p.handle, dptr, C.size_t(len(data)), (*C.uint8_t)(buf), &recvLen)
	if cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return C.GoBytes(unsafe.Pointer(buf), C.int(recvLen)), nil
}

// I2CSendReceive sends data to an I2C device and receives up to maxReceive bytes
func (p *PhidgetDataAdapter) I2CSendReceive(address int32, data []byte, maxReceive int) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("gophidgets: I2C data must not be empty")
	}
	if maxReceive <= 0 {
		return nil, errors.New("gophidgets: receive length must be positive")
	}
	dptr := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	buf := C.malloc(C.size_t(maxReceive))
	defer C.free(buf)
	if cerr := C.PhidgetDataAdapter_i2cSendReceive(
		p.handle, C.int32_t(address), dptr, C.size_t(len(data)), (*C.uint8_t)(buf), C.size_t(maxReceive)); cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return C.GoBytes(unsafe.Pointer(buf), C.int(maxReceive)), nil
}

// I2CComplexTransaction performs a complex I2C transaction described by
// packetString ('s' = start, 'R' = receive, etc.). data holds the bytes to
// send, in order; the receive bytes are returned, up to maxReceive bytes.
func (p *PhidgetDataAdapter) I2CComplexTransaction(address int32, packetString string, data []byte, maxReceive int) ([]byte, error) {
	if maxReceive <= 0 {
		return nil, errors.New("gophidgets: receive length must be positive")
	}
	ckey := C.CString(packetString)
	defer C.free(unsafe.Pointer(ckey))
	var dptr *C.uint8_t
	if len(data) > 0 {
		dptr = (*C.uint8_t)(unsafe.Pointer(&data[0]))
	}
	buf := C.malloc(C.size_t(maxReceive))
	defer C.free(buf)
	recvLen := C.size_t(maxReceive)
	cerr := C.PhidgetDataAdapter_i2cComplexTransaction(
		p.handle, C.int32_t(address), ckey, dptr, C.size_t(len(data)), (*C.uint8_t)(buf), &recvLen)
	if cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return C.GoBytes(unsafe.Pointer(buf), C.int(recvLen)), nil
}

// Close closes the handle and deletes it
func (p *PhidgetDataAdapter) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetDataAdapter_delete(&p.handle))
}

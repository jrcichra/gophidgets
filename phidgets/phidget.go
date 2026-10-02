package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime/cgo"
	"time"
	"unsafe"
)

// Phidget - general phidget interface that all phidgets are derived from (for ease of type management)
type phidget struct {
	handle C.PhidgetHandle
	ctxs   []cgo.Handle // handler contexts handed to C; released on Close
}

// save registers f as a C handler context and remembers it so Close can release it
func (p *phidget) save(f interface{}) unsafe.Pointer {
	h := cgo.NewHandle(f)
	p.ctxs = append(p.ctxs, h)
	return C.handlectx(C.uintptr_t(h))
}

type Phidget interface {
	Create()
	Close() error
	String() string

	OpenWaitForAttachment(timeout time.Duration) error

	SetIsRemote(b bool) error
	GetIsRemote() (bool, error)
	SetDeviceSerialNumber(serial int) error
	GetDeviceSerialNumber() (int, error)
	SetHubPort(port int) error
	GetHubPort() (int, error)
	SetChannel(port int) error
	GetChannel() (int, error)

	SetIsHubPortDevice(b bool) error
	GetChannelClassName() (string, error)
	GetChannelName() (string, error)
	GetAttached() (bool, error)

	// Channel-level event handlers
	SetOnAttachHandler(f func()) error
	SetOnDetachHandler(f func()) error
	SetOnErrorHandler(f func(code int, message string)) error

	// Unexported function for internal management
	getRawHandle() *C.PhidgetHandle
}

// rawHandle updates the base Phidget object with a handle from another object
// This must be called when a new Phidget object is created of a differing class
func (p *phidget) rawHandle(handle unsafe.Pointer) {
	p.handle = (*C.struct__Phidget)(handle)
}

func (p *phidget) getRawHandle() *C.PhidgetHandle {
	return &p.handle
}

// Open opens the phidget immediately, without waiting for it to be attached.
func (p *phidget) Open() error {
	return p.phidgetError(C.Phidget_open(p.handle))
}

// OpenWaitForAttachment opens a phidget and waits for it to be available on the bus
func (p *phidget) OpenWaitForAttachment(timeout time.Duration) error {
	if cerr := C.Phidget_openWaitForAttachment(p.handle, C.uint(timeout.Milliseconds())); cerr != C.EPHIDGET_OK {
		return p.phidgetError(cerr)
	}
	return nil
}

// PhidgetError is the typed error returned when a libphidget22 call fails.
// Code() returns the underlying EPHIDGET_* return code and the error
// message is unchanged.
type PhidgetError struct {
	code    int32
	message string
}

func (e *PhidgetError) Error() string { return e.message }

// Code returns the underlying EPHIDGET_* return code
func (e *PhidgetError) Code() int32 { return e.code }

var (
	// ErrTimeout is matched by errors.Is when a libphidget22 call
	// (such as OpenWaitForAttachment) returns EPHIDGET_TIMEOUT.
	ErrTimeout = errors.New("gophidgets: timeout")
	// ErrUnknownValue is matched by errors.Is when a libphidget22 call
	// returns EPHIDGET_UNKNOWNVAL.
	ErrUnknownValue = errors.New("gophidgets: unknown value")
	// ErrNotAttached is matched by errors.Is when a libphidget22 call
	// returns EPHIDGET_NOTATTACHED.
	ErrNotAttached = errors.New("gophidgets: not attached")
	// ErrClosed is matched by errors.Is when a libphidget22 call
	// returns EPHIDGET_CLOSED.
	ErrClosed = errors.New("gophidgets: closed")
	// ErrUnsupported is matched by errors.Is when a libphidget22 call
	// returns EPHIDGET_UNSUPPORTED.
	ErrUnsupported = errors.New("gophidgets: unsupported")
	// ErrInvalidArg is matched by errors.Is when a libphidget22 call
	// returns EPHIDGET_INVALIDARG.
	ErrInvalidArg = errors.New("gophidgets: invalid argument")
	// ErrFailSafe is matched by errors.Is when a libphidget22 call
	// returns EPHIDGET_FAILSAFE.
	ErrFailSafe = errors.New("gophidgets: failsafe")
)

func (e *PhidgetError) Is(target error) bool {
	switch e.code {
	case int32(C.EPHIDGET_TIMEOUT):
		return target == ErrTimeout
	case int32(C.EPHIDGET_UNKNOWNVAL):
		return target == ErrUnknownValue
	case int32(C.EPHIDGET_NOTATTACHED):
		return target == ErrNotAttached
	case int32(C.EPHIDGET_CLOSED):
		return target == ErrClosed
	case int32(C.EPHIDGET_UNSUPPORTED):
		return target == ErrUnsupported
	case int32(C.EPHIDGET_INVALIDARG):
		return target == ErrInvalidArg
	case int32(C.EPHIDGET_FAILSAFE):
		return target == ErrFailSafe
	}
	return false
}

func newPhidgetError(code C.PhidgetReturnCode) error {
	if code == C.EPHIDGET_OK {
		return nil
	}
	var errString *C.char
	C.Phidget_getErrorDescription(code, &errString)
	return &PhidgetError{code: int32(code), message: C.GoString(errString)}
}

func (p *phidget) phidgetError(cerr C.PhidgetReturnCode) error {
	err := newPhidgetError(cerr)
	var className *C.char
	if err != nil && C.Phidget_getChannelClassName(p.handle, &className) == C.EPHIDGET_OK {
		err.(*PhidgetError).message = C.GoString(className) + ": " + err.Error()
	}
	return err
}

// SetOnAttachHandler sets a callback that is called when the channel attaches
func (p *phidget) SetOnAttachHandler(f func()) error {
	ctx := p.save(f)
	return p.phidgetError(C.Phidget_setOnAttachHandler(
		p.handle, (C.phidget_void_fcn)(unsafe.Pointer(C.voidcallback)), ctx))
}

// SetOnDetachHandler sets a callback that is called when the channel detaches
func (p *phidget) SetOnDetachHandler(f func()) error {
	ctx := p.save(f)
	return p.phidgetError(C.Phidget_setOnDetachHandler(
		p.handle, (C.phidget_void_fcn)(unsafe.Pointer(C.voidcallback)), ctx))
}

// SetOnErrorHandler sets a callback that is called when the channel reports an error.
// The callback receives the EEPHIDGET_* error event code and the error message.
func (p *phidget) SetOnErrorHandler(f func(code int, message string)) error {
	ctx := p.save(f)
	return p.phidgetError(C.Phidget_setOnErrorHandler(
		p.handle, (C.phidget_error_fcn)(unsafe.Pointer(C.errorcallback)), ctx))
}

// SetOnPropertyChangeHandler sets a callback that fires when a channel property
// is changed externally (e.g. from a network client)
func (p *phidget) SetOnPropertyChangeHandler(f func(propertyName string)) error {
	ctx := p.save(f)
	return p.phidgetError(C.Phidget_setOnPropertyChangeHandler(
		p.handle, (C.phidget_prop_fcn)(unsafe.Pointer(C.propcallback)), ctx))
}

// SetIsRemote sets a phidget sensor as a remote device
func (p *phidget) SetIsRemote(b bool) error {
	return p.phidgetError(C.Phidget_setIsRemote(p.handle, boolToCInt(b)))
}

// SetDeviceSerialNumber sets the serial number to use.
// This must be called before calling OpenWaitForAttachment
func (p *phidget) SetDeviceSerialNumber(serial int) error {
	return p.phidgetError(C.Phidget_setDeviceSerialNumber(p.handle, C.int(serial)))
}

// SetHubPort sets a phidget's hub port
func (p *phidget) SetHubPort(port int) error {
	return p.phidgetError(C.Phidget_setHubPort(p.handle, C.int(port)))
}

// GetHubPort gets a phidget's hub port
func (p *phidget) GetHubPort() (int, error) {
	return getInt(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getHubPort(p.handle, r) })
}

// SetChannel sets a phidget motion sensor's channel port
func (p *phidget) SetChannel(port int) error {
	return p.phidgetError(C.Phidget_setChannel(p.handle, C.int(port)))
}

// GetIsRemote gets a phidget's remote status
func (p *phidget) GetIsRemote() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getIsRemote(p.handle, r) })
}

// GetDeviceSerialNumber gets a phidget motion sensor's serial number
func (p *phidget) GetDeviceSerialNumber() (int, error) {
	return getInt(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getDeviceSerialNumber(p.handle, r) })
}

// GetDeviceID returns the library ID of the attached device
func (p *phidget) GetDeviceID() (int, error) {
	return get(p, func(r *C.Phidget_DeviceID) C.PhidgetReturnCode { return C.Phidget_getDeviceID(p.handle, r) }, func(r C.Phidget_DeviceID) int { return int(r) })
}

// GetDeviceName returns the name of the attached device
func (p *phidget) GetDeviceName() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getDeviceName(p.handle, r) })
}

// GetDeviceVersion returns the firmware version of the attached device
func (p *phidget) GetDeviceVersion() (int, error) {
	return getInt(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getDeviceVersion(p.handle, r) })
}

// GetDeviceSKU returns the SKU of the attached device
func (p *phidget) GetDeviceSKU() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getDeviceSKU(p.handle, r) })
}

// GetDeviceVINTID returns the VINT ID of the attached device
func (p *phidget) GetDeviceVINTID() (uint32, error) {
	return getUint32(p, func(r *C.uint32_t) C.PhidgetReturnCode { return C.Phidget_getDeviceVINTID(p.handle, r) })
}

// GetDeviceClass returns the device class of the attached device
func (p *phidget) GetDeviceClass() (int, error) {
	return get(p, func(r *C.Phidget_DeviceClass) C.PhidgetReturnCode { return C.Phidget_getDeviceClass(p.handle, r) }, func(r C.Phidget_DeviceClass) int { return int(r) })
}

// GetDeviceClassName returns the class name of the attached device
func (p *phidget) GetDeviceClassName() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getDeviceClassName(p.handle, r) })
}

// GetDeviceChannelCount returns the number of channels of the given class
// (a PHIDGETCHCLASS_* value) that the attached device has
func (p *phidget) GetDeviceChannelCount(cls int) (uint32, error) {
	return getUint32(p, func(r *C.uint32_t) C.PhidgetReturnCode {
		return C.Phidget_getDeviceChannelCount(p.handle, C.Phidget_ChannelClass(cls), r)
	})
}

// GetDeviceFirmwareUpgradeString returns the VINT-compatible firmware upgrade
// string of the attached device
func (p *phidget) GetDeviceFirmwareUpgradeString() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getDeviceFirmwareUpgradeString(p.handle, r) })
}

// GetChannelClass returns the class of this channel (a PHIDGETCHCLASS_* value)
func (p *phidget) GetChannelClass() (int, error) {
	return get(p, func(r *C.Phidget_ChannelClass) C.PhidgetReturnCode { return C.Phidget_getChannelClass(p.handle, r) }, func(r C.Phidget_ChannelClass) int { return int(r) })
}

// GetChannelSubclass returns the subclass of this channel
func (p *phidget) GetChannelSubclass() (int, error) {
	return get(p, func(r *C.Phidget_ChannelSubclass) C.PhidgetReturnCode {
		return C.Phidget_getChannelSubclass(p.handle, r)
	}, func(r C.Phidget_ChannelSubclass) int { return int(r) })
}

// GetChannelPersistence returns the hub-port channel persistence state
// of this channel (0 = off, 1 = on).
func (p *phidget) GetChannelPersistence() (int, error) {
	return getInt(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getChannelPersistence(p.handle, r) })
}

// SetChannelPersistence sets the hub-port channel persistence state of this
// channel (0 = off, 1 = on).
func (p *phidget) SetChannelPersistence(mode int) error {
	return p.phidgetError(C.Phidget_setChannelPersistence(p.handle, C.int(mode)))
}

// GetHub returns the serial number of the hub this channel is attached to
func (p *phidget) GetHub() (int, error) {
	var hub C.PhidgetHandle
	if cerr := C.Phidget_getHub(p.handle, &hub); cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	var serial C.int
	if cerr := C.Phidget_getDeviceSerialNumber(hub, &serial); cerr != C.EPHIDGET_OK {
		C.Phidget_release(&hub)
		return 0, p.phidgetError(cerr)
	}
	C.Phidget_release(&hub)
	return int(serial), nil
}

// GetIsOpen returns whether this channel is open
func (p *phidget) GetIsOpen() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getIsOpen(p.handle, r) })
}

// GetIsLocal returns whether this channel is attached to the local machine
func (p *phidget) GetIsLocal() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getIsLocal(p.handle, r) })
}

// SetIsLocal sets whether this channel is treated as local
func (p *phidget) SetIsLocal(local bool) error {
	return p.phidgetError(C.Phidget_setIsLocal(p.handle, boolToCInt(local)))
}

// GetIsChannel returns whether this object is a channel (as opposed to a device)
func (p *phidget) GetIsChannel() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getIsChannel(p.handle, r) })
}

// GetIsHubPortDevice returns whether this channel is a VINT device on a hub port.
// (SetIsHubPortDevice already exists on the base type.)
func (p *phidget) GetIsHubPortDevice() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getIsHubPortDevice(p.handle, r) })
}

// GetHubPortCount returns the number of hub ports on the attached device
func (p *phidget) GetHubPortCount() (int, error) {
	return getInt(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getHubPortCount(p.handle, r) })
}

// GetHubPortSpeed returns the hub port speed in bps of this channel
func (p *phidget) GetHubPortSpeed() (uint32, error) {
	return getUint32(p, func(r *C.uint32_t) C.PhidgetReturnCode { return C.Phidget_getHubPortSpeed(p.handle, r) })
}

// SetHubPortSpeed sets the hub port speed in bps of this channel. Pass
// HubPortSpeedAuto (0) to let the hub negotiate automatically when supported.
func (p *phidget) SetHubPortSpeed(speed uint32) error {
	return p.phidgetError(C.Phidget_setHubPortSpeed(p.handle, C.uint32_t(speed)))
}

// GetMaxHubPortSpeed returns the maximum supported hub port speed in bps
func (p *phidget) GetMaxHubPortSpeed() (uint32, error) {
	return getUint32(p, func(r *C.uint32_t) C.PhidgetReturnCode { return C.Phidget_getMaxHubPortSpeed(p.handle, r) })
}

// GetHubPortSupportsAutoSetSpeed returns whether this channel's hub port
// supports automatic speed setting
func (p *phidget) GetHubPortSupportsAutoSetSpeed() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getHubPortSupportsAutoSetSpeed(p.handle, r) })
}

// GetHubPortSupportsSetSpeed returns whether this channel's hub port
// supports manual speed setting
func (p *phidget) GetHubPortSupportsSetSpeed() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getHubPortSupportsSetSpeed(p.handle, r) })
}

// GetMaxVINTDeviceSpeed returns the maximum speed in bps supported by the VINT device
func (p *phidget) GetMaxVINTDeviceSpeed() (uint32, error) {
	return getUint32(p, func(r *C.uint32_t) C.PhidgetReturnCode { return C.Phidget_getMaxVINTDeviceSpeed(p.handle, r) })
}

// GetVINTDeviceSupportsAutoSetSpeed returns whether the VINT device supports
// automatic speed setting
func (p *phidget) GetVINTDeviceSupportsAutoSetSpeed() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getVINTDeviceSupportsAutoSetSpeed(p.handle, r) })
}

// GetVINTDeviceSupportsSetSpeed returns whether the VINT device supports
// manual speed setting
func (p *phidget) GetVINTDeviceSupportsSetSpeed() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getVINTDeviceSupportsSetSpeed(p.handle, r) })
}

// GetDataInterval returns the data interval in milliseconds
func (p *phidget) GetDataInterval() (uint32, error) {
	return getUint32(p, func(r *C.uint32_t) C.PhidgetReturnCode { return C.Phidget_getDataInterval(p.handle, r) })
}

// SetDataInterval sets the data interval in milliseconds
func (p *phidget) SetDataInterval(ms uint32) error {
	return p.phidgetError(C.Phidget_setDataInterval(p.handle, C.uint32_t(ms)))
}

// GetMinDataInterval returns the minimum data interval in milliseconds
func (p *phidget) GetMinDataInterval() (uint32, error) {
	return getUint32(p, func(r *C.uint32_t) C.PhidgetReturnCode { return C.Phidget_getMinDataInterval(p.handle, r) })
}

// GetMaxDataInterval returns the maximum data interval in milliseconds
func (p *phidget) GetMaxDataInterval() (uint32, error) {
	return getUint32(p, func(r *C.uint32_t) C.PhidgetReturnCode { return C.Phidget_getMaxDataInterval(p.handle, r) })
}

// GetDataRate returns the data rate in samples per second
func (p *phidget) GetDataRate() (float64, error) {
	return getDouble(p, func(r *C.double) C.PhidgetReturnCode { return C.Phidget_getDataRate(p.handle, r) })
}

// SetDataRate sets the data rate in samples per second
func (p *phidget) SetDataRate(rate float64) error {
	return p.phidgetError(C.Phidget_setDataRate(p.handle, C.double(rate)))
}

// GetMinDataRate returns the minimum data rate in samples per second
func (p *phidget) GetMinDataRate() (float64, error) {
	return getDouble(p, func(r *C.double) C.PhidgetReturnCode { return C.Phidget_getMinDataRate(p.handle, r) })
}

// GetMaxDataRate returns the maximum data rate in samples per second
func (p *phidget) GetMaxDataRate() (float64, error) {
	return getDouble(p, func(r *C.double) C.PhidgetReturnCode { return C.Phidget_getMaxDataRate(p.handle, r) })
}

// GetServerName returns the name of the phidget22 server in use
func (p *phidget) GetServerName() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getServerName(p.handle, r) })
}

// SetServerName sets the name of the phidget22 server to use
func (p *phidget) SetServerName(name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return p.phidgetError(C.Phidget_setServerName(p.handle, cname))
}

// GetServerHostname returns the hostname of the phidget22 server in use
func (p *phidget) GetServerHostname() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getServerHostname(p.handle, r) })
}

// GetServerPeerName returns the peer name of the phidget22 server in use
func (p *phidget) GetServerPeerName() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getServerPeerName(p.handle, r) })
}

// GetServerUniqueName returns the unique name of the phidget22 server in use
func (p *phidget) GetServerUniqueName() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getServerUniqueName(p.handle, r) })
}

// GetServerVersion returns the major and minor versions of the phidget22 server in use
func (p *phidget) GetServerVersion() (major, minor int, err error) {
	var cmaj, cmin C.int
	if cerr := C.Phidget_getServerVersion(p.handle, &cmaj, &cmin); cerr != C.EPHIDGET_OK {
		return 0, 0, p.phidgetError(cerr)
	}
	return int(cmaj), int(cmin), nil
}

// GetClientVersion returns the major and minor versions of the client library
func (p *phidget) GetClientVersion() (major, minor int, err error) {
	var cmaj, cmin C.int
	if cerr := C.Phidget_getClientVersion(p.handle, &cmaj, &cmin); cerr != C.EPHIDGET_OK {
		return 0, 0, p.phidgetError(cerr)
	}
	return int(cmaj), int(cmin), nil
}

// Reboot reboots the attached device.
func (p *phidget) Reboot() error {
	return p.phidgetError(C.Phidget_reboot(p.handle))
}

// WriteDeviceLabel writes the given label to the attached device's flash, then
// commits it. GetDeviceLabel already exists on the base type.
func (p *phidget) WriteDeviceLabel(label string) error {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return p.phidgetError(C.Phidget_writeDeviceLabel(p.handle, clabel))
}

// WriteFlash commits pending flash writes (e.g. a previously set label) to flash
func (p *phidget) WriteFlash() error {
	return p.phidgetError(C.Phidget_writeFlash(p.handle))
}

// SetDeviceLabel sets the label of the attached device in memory. Call
// WriteFlash (or WriteDeviceLabel) to persist it.
func (p *phidget) SetDeviceLabel(label string) error {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return p.phidgetError(C.Phidget_setDeviceLabel(p.handle, clabel))
}

// Close closes the channel and deletes its handle
func (p *phidget) Close() error {
	if err := p.phidgetError(C.Phidget_close(p.handle)); err != nil {
		return err
	}
	// Phidget_close waits for in-flight handlers, so the contexts are safe to release
	for _, h := range p.ctxs {
		h.Delete()
	}
	p.ctxs = nil
	return p.phidgetError(C.Phidget_delete(&p.handle))
}

// GetChannelClassName gets the name of the channel class the channel belongs to.
func (p *phidget) GetChannelClassName() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getChannelClassName(p.handle, r) })
}

// SetIsHubPortDevice sets a phidget sensor as a remote device
func (p *phidget) SetIsHubPortDevice(b bool) error {
	return p.phidgetError(C.Phidget_setIsHubPortDevice(p.handle, boolToCInt(b)))
}

// GetDeviceLabel retrieves the label for the device of this handle
func (p *phidget) GetDeviceLabel() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getDeviceLabel(p.handle, r) })
}

// GetChannelName retrieves the channel name of this handle
func (p *phidget) GetChannelName() (string, error) {
	return getString(p, func(r **C.char) C.PhidgetReturnCode { return C.Phidget_getChannelName(p.handle, r) })
}

// GetAttached returns whether the Phidget device is attached
func (p *phidget) GetAttached() (bool, error) {
	return getBool(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getAttached(p.handle, r) })
}

// GetChannel retrives which channel this handle is attached to
func (p *phidget) GetChannel() (int, error) {
	return getInt(p, func(r *C.int) C.PhidgetReturnCode { return C.Phidget_getChannel(p.handle, r) })
}

// String returns a string description of the handle
// implements Gos stringer interface to ease printing
func (p *phidget) String() string {
	name, _ := p.GetChannelName()
	class, _ := p.GetChannelClassName()
	channel, _ := p.GetChannel()
	serial, _ := p.GetDeviceSerialNumber()
	label, _ := p.GetDeviceLabel()
	port, _ := p.GetHubPort()

	return fmt.Sprintf("%s: %s channel %d [ser=%x] [label=%s] [port=%d]",
		name, class, channel, serial, label, port)
}

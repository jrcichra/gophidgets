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

// PhidgetDictionary wraps a Phidget Dictionary
type PhidgetDictionary struct {
	phidget
	handle C.PhidgetDictionaryHandle
}

// Create creates a PhidgetDictionary handle
func (p *PhidgetDictionary) Create() {
	C.PhidgetDictionary_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// withKV runs f with key and value as C strings
func (p *PhidgetDictionary) withKV(key, value string, f func(k, v *C.char) C.PhidgetReturnCode) error {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	return p.phidgetError(f(ckey, cvalue))
}

// read runs f with arg as a C string and a maxLen-byte result buffer
func (p *PhidgetDictionary) read(arg string, maxLen int, f func(arg, buf *C.char, n C.size_t) C.PhidgetReturnCode) (string, error) {
	if maxLen <= 0 {
		return "", errors.New("gophidgets: buffer length must be positive")
	}
	carg := C.CString(arg)
	defer C.free(unsafe.Pointer(carg))
	buf := C.malloc(C.size_t(maxLen))
	defer C.free(buf)
	if cerr := f(carg, (*C.char)(buf), C.size_t(maxLen)); cerr != C.EPHIDGET_OK {
		return "", p.phidgetError(cerr)
	}
	return string(C.GoBytes(buf, C.int(maxLen))), nil
}

// Add adds a key/value pair to the dictionary
func (p *PhidgetDictionary) Add(key, value string) error {
	return p.withKV(key, value, func(k, v *C.char) C.PhidgetReturnCode { return C.PhidgetDictionary_add(p.handle, k, v) })
}

// Set sets the value for a key in the dictionary
func (p *PhidgetDictionary) Set(key, value string) error {
	return p.withKV(key, value, func(k, v *C.char) C.PhidgetReturnCode { return C.PhidgetDictionary_set(p.handle, k, v) })
}

// Update updates the value for a key in the dictionary
func (p *PhidgetDictionary) Update(key, value string) error {
	return p.withKV(key, value, func(k, v *C.char) C.PhidgetReturnCode { return C.PhidgetDictionary_update(p.handle, k, v) })
}

// Get returns the value associated with key. maxLen is the size in bytes of
// the buffer the value is read into; longer values are truncated.
func (p *PhidgetDictionary) Get(key string, maxLen int) (string, error) {
	return p.read(key, maxLen, func(k, buf *C.char, n C.size_t) C.PhidgetReturnCode {
		return C.PhidgetDictionary_get(p.handle, k, buf, n)
	})
}

// Remove removes the key from the dictionary
func (p *PhidgetDictionary) Remove(key string) error {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	return p.phidgetError(C.PhidgetDictionary_remove(p.handle, ckey))
}

// RemoveAll removes every key from the dictionary
func (p *PhidgetDictionary) RemoveAll() error {
	return p.phidgetError(C.PhidgetDictionary_removeAll(p.handle))
}

// Scan returns keys in the dictionary that start with the given prefix.
// maxLen is the size in bytes of the buffer the key list is read into.
func (p *PhidgetDictionary) Scan(prefix string, maxLen int) (string, error) {
	return p.read(prefix, maxLen, func(k, buf *C.char, n C.size_t) C.PhidgetReturnCode {
		return C.PhidgetDictionary_scan(p.handle, k, buf, n)
	})
}

// SetOnAddHandler sets a callback that fires when a key is added
func (p *PhidgetDictionary) SetOnAddHandler(f func(key, value string)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDictionary_setOnAddHandler(
		p.handle, (C.phidget_kv_fcn)(unsafe.Pointer(C.dictkvcallback)), ctx))
}

// SetOnUpdateHandler sets a callback that fires when a key's value is updated
func (p *PhidgetDictionary) SetOnUpdateHandler(f func(key, value string)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDictionary_setOnUpdateHandler(
		p.handle, (C.phidget_kv_fcn)(unsafe.Pointer(C.dictkvcallback)), ctx))
}

// SetOnRemoveHandler sets a callback that fires when a key is removed
func (p *PhidgetDictionary) SetOnRemoveHandler(f func(key string)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDictionary_setOnRemoveHandler(
		p.handle, (C.phidget_key_fcn)(unsafe.Pointer(C.dictkeycallback)), ctx))
}

// AddDictionary associates a label with a device serial number in the
// system dictionary.
func AddDictionary(serial int, label string) error {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return newPhidgetError(C.PhidgetDictionary_addDictionary(C.int(serial), clabel))
}

// RemoveDictionary removes the dictionary entry for the device serial number
func RemoveDictionary(serial int) error {
	return newPhidgetError(C.PhidgetDictionary_removeDictionary(C.int(serial)))
}

// LoadDictionary loads the dictionary from a JSON file
func LoadDictionary(serial int, file string) error {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	return newPhidgetError(C.PhidgetDictionary_loadDictionary(C.int(serial), cfile))
}

// EnableControlDictionary enables writing to the control dictionary
func EnableControlDictionary() error {
	return newPhidgetError(C.PhidgetDictionary_enableControlDictionary())
}

// EnableStatsDictionary enables updating the statistics dictionary
func EnableStatsDictionary() error {
	return newPhidgetError(C.PhidgetDictionary_enableStatsDictionary())
}

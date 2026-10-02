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

// Add adds a key/value pair to the dictionary
func (p *PhidgetDictionary) Add(key, value string) error {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	return p.phidgetError(C.PhidgetDictionary_add(p.handle, ckey, cvalue))
}

// Set sets the value for a key in the dictionary
func (p *PhidgetDictionary) Set(key, value string) error {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	return p.phidgetError(C.PhidgetDictionary_set(p.handle, ckey, cvalue))
}

// Update updates the value for a key in the dictionary
func (p *PhidgetDictionary) Update(key, value string) error {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	return p.phidgetError(C.PhidgetDictionary_update(p.handle, ckey, cvalue))
}

// Get returns the value associated with key. maxLen is the size in bytes of
// the buffer the value is read into; longer values are truncated.
func (p *PhidgetDictionary) Get(key string, maxLen int) (string, error) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	if maxLen <= 0 {
		return "", errors.New("gophidgets: value length must be positive")
	}
	buf := C.malloc(C.size_t(maxLen))
	defer C.free(buf)
	if cerr := C.PhidgetDictionary_get(p.handle, ckey, (*C.char)(buf), C.size_t(maxLen)); cerr != C.EPHIDGET_OK {
		return "", p.phidgetError(cerr)
	}
	return string(C.GoBytes(buf, C.int(maxLen))), nil
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
	cprefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cprefix))
	if maxLen <= 0 {
		return "", errors.New("gophidgets: key list length must be positive")
	}
	buf := C.malloc(C.size_t(maxLen))
	defer C.free(buf)
	if cerr := C.PhidgetDictionary_scan(p.handle, cprefix, (*C.char)(buf), C.size_t(maxLen)); cerr != C.EPHIDGET_OK {
		return "", p.phidgetError(cerr)
	}
	return string(C.GoBytes(buf, C.int(maxLen))), nil
}

// SetOnAddHandler sets a callback that fires when a key is added
func (p *PhidgetDictionary) SetOnAddHandler(f func(key, value string)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDictionary_setOnAddHandler(
		p.handle, (C.phidget_kv_fcn)(unsafe.Pointer(C.cdictkvcallback)), ctx))
}

// SetOnUpdateHandler sets a callback that fires when a key's value is updated
func (p *PhidgetDictionary) SetOnUpdateHandler(f func(key, value string)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDictionary_setOnUpdateHandler(
		p.handle, (C.phidget_kv_fcn)(unsafe.Pointer(C.cdictkvcallback)), ctx))
}

// SetOnRemoveHandler sets a callback that fires when a key is removed
func (p *PhidgetDictionary) SetOnRemoveHandler(f func(key string)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDictionary_setOnRemoveHandler(
		p.handle, (C.phidget_key_fcn)(unsafe.Pointer(C.cdictkeycallback)), ctx))
}

// Close closes the handle and deletes it
func (p *PhidgetDictionary) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetDictionary_delete(&p.handle))
}

// AddDictionary associates a label with a device serial number in the
// system dictionary.
func AddDictionary(serial int, label string) error {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return managerError(C.PhidgetDictionary_addDictionary(C.int(serial), clabel))
}

// RemoveDictionary removes the dictionary entry for the device serial number
func RemoveDictionary(serial int) error {
	return managerError(C.PhidgetDictionary_removeDictionary(C.int(serial)))
}

// LoadDictionary loads the dictionary from a JSON file
func LoadDictionary(serial int, file string) error {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	return managerError(C.PhidgetDictionary_loadDictionary(C.int(serial), cfile))
}

// EnableControlDictionary enables writing to the control dictionary
func EnableControlDictionary() error {
	return managerError(C.PhidgetDictionary_enableControlDictionary())
}

// EnableStatsDictionary enables updating the statistics dictionary
func EnableStatsDictionary() error {
	return managerError(C.PhidgetDictionary_enableStatsDictionary())
}

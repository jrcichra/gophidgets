package phidgets

/*
#include <phidget22.h>
#include "phidgets.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// PhidgetDistanceSensor is the struct that is a phidget distance sensor
type PhidgetDistanceSensor struct {
	phidget
	handle C.PhidgetDistanceSensorHandle
}

// Create creates a phidget distance sensor
func (p *PhidgetDistanceSensor) Create() {
	C.PhidgetDistanceSensor_create(&p.handle)
	p.rawHandle(unsafe.Pointer(p.handle))
}

// GetDistance - The most recent distance value that the channel has reported.
func (p *PhidgetDistanceSensor) GetDistance() (int, error) {
	var r C.uint
	cerr := C.PhidgetDistanceSensor_getDistance(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, p.phidgetError(cerr)
	}
	return int(r), nil
}

// SetDistanceChangeTrigger sets the interrupt trigger point
func (p *PhidgetDistanceSensor) SetDistanceChangeTrigger(distance uint32) error {
	return p.phidgetError(C.PhidgetDistanceSensor_setDistanceChangeTrigger(p.handle, C.uint(distance)))
}

// SetOnDistanceChangeHandler sets a callback that is called when the
// reported distance changes. The callback receives the distance (mm).
func (p *PhidgetDistanceSensor) SetOnDistanceChangeHandler(f func(uint32)) error {
	ctx := p.save(f)
	return p.phidgetError(C.PhidgetDistanceSensor_setOnDistanceChangeHandler(
		p.handle, (C.phidget_uint32_fcn)(unsafe.Pointer(C.uint32callback)), ctx))
}

// GetSonarReflections - The most recent reflection values that the channel has reported.
func (p *PhidgetDistanceSensor) GetSonarReflections() ([]uint32, []uint32, error) {
	var cDistances [8]C.uint
	var cAmplitudes [8]C.uint
	var count C.uint
	cerr := C.PhidgetDistanceSensor_getSonarReflections(p.handle, &cDistances, &cAmplitudes, &count)
	if cerr != C.EPHIDGET_OK {
		return nil, nil, p.phidgetError(cerr)
	}

	iCount := int(count)
	// trim arrays to count size (removes 2^32 − 1 values)
	distances := make([]uint32, 0, iCount)
	for i := 0; i < iCount; i++ {
		distances = append(distances, uint32(cDistances[i]))
	}
	amplitudes := make([]uint32, 0, iCount)
	for i := 0; i < iCount; i++ {
		amplitudes = append(amplitudes, uint32(cAmplitudes[i]))
	}

	if len(distances) != len(amplitudes) {
		return distances, amplitudes, fmt.Errorf("length of distances: %d does not match the length of amplitudes: %d", len(distances), len(amplitudes))
	}

	return distances, amplitudes, nil
}

func (p *PhidgetDistanceSensor) SetSonarQuietMode(val bool) error {
	return p.phidgetError(C.PhidgetDistanceSensor_setSonarQuietMode(p.handle, boolToCInt(val)))
}

func (p *PhidgetDistanceSensor) GetSonarQuietMode() (bool, error) {
	var r C.int
	err := p.phidgetError(C.PhidgetDistanceSensor_getSonarQuietMode(p.handle, &r))
	return r > 0, err
}

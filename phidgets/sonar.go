package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*distance_callback_fcn)(void* handle, void* ctx, uint32_t distance);
void cdistancecallback(void* handle, void* ctx, uint32_t distance);  // Forward declaration.
typedef void (*reflection_callback_fcn)(void* handle, void* ctx, const uint32_t distances[8], const uint32_t amplitudes[8], uint32_t count);
void creflectioncallback(void* handle, void* ctx, const uint32_t distances[8], const uint32_t amplitudes[8], uint32_t count);  // Forward declaration.
*/
import "C"
import (
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

// SetOnDistanceChangeHandler - interrupt for distance changes calls a function
// func (p *PhidgetDistanceSensor) SetOnDistanceChangeHandler(f func(uint32)) error {
// 	//make a c function pointer to a go function pointer and pass it through the phidget context
// 	var passthrough DistancePassthrough
// 	passthrough.f = f
// 	pt := gopointer.Save(passthrough)
// 	cerr := C.PhidgetDistanceSensor_setOnDistanceChangeHandler(p.handle, (C.distance_callback_fcn)(unsafe.Pointer(C.cdistancecallback)), pt)
// 	if cerr != C.EPHIDGET_OK {
// 		return p.phidgetError(cerr)
// 	}
// 	return nil
// }

// // setOnSonarReflectionsUpdateHandler - interrupt for sonar reflections
// func (p *PhidgetDistanceSensor) setOnSonarReflectionsUpdateHandler(f func([8]uint32, [8]uint32, uint32)) error {
// 	//make a c function pointer to a go function pointer and pass it through the phidget context
// 	var passthrough ReflectionPassthrough
// 	passthrough.f = f
// 	pt := gopointer.Save(passthrough)
// 	cerr := C.PhidgetDistanceSensor_setOnDistanceChangeHandler(p.handle, (C.reflection_callback_fcn)(unsafe.Pointer(C.creflectioncallback)), pt)
// 	if cerr != C.EPHIDGET_OK {
// 		return p.phidgetError(cerr)
// 	}
// 	return nil
// }

// GetSonarReflections - The most recent reflection values that the channel has reported.
func (p *PhidgetDistanceSensor) GetSonarReflections() ([]uint32, []uint32, uint32, error) {
	var cDistances [8]C.uint
	var cAmplitudes [8]C.uint
	var count C.uint
	cerr := C.PhidgetDistanceSensor_getSonarReflections(p.handle, &cDistances, &cAmplitudes, &count)
	if cerr != C.EPHIDGET_OK {
		return nil, nil, 0, p.phidgetError(cerr)
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

	return distances, amplitudes, uint32(count), nil
}

func (p *PhidgetDistanceSensor) SetSonarQuietMode(val bool) error {
	return p.phidgetError(C.PhidgetDistanceSensor_setSonarQuietMode(p.handle, boolToCInt(val)))
}

func (p *PhidgetDistanceSensor) GetSonarQuietMode() (bool, error) {
	var r C.int
	err := p.phidgetError(C.PhidgetDistanceSensor_getSonarQuietMode(p.handle, &r))
	return r > 0, err
}

// Close - close the handle and delete it
func (p *PhidgetDistanceSensor) Close() error {
	if err := p.phidget.Close(); err != nil {
		return err
	}
	return p.phidgetError(C.PhidgetDistanceSensor_delete(&p.handle))
}

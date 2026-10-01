package phidgets

/*
#include <phidget22.h>
*/
import "C"

// GetEulerAngles returns the current orientation as Euler angles (roll, pitch, heading) in degrees.
// Hand-written because the C API returns a struct, which the generator doesn't handle.
func (p *PhidgetSpatial) GetEulerAngles() (roll, pitch, heading float64, err error) {
	var angles C.PhidgetSpatial_SpatialEulerAngles
	if cerr := C.PhidgetSpatial_getEulerAngles(p.handle, &angles); cerr != C.EPHIDGET_OK {
		return 0, 0, 0, p.phidgetError(cerr)
	}
	return float64(angles.roll), float64(angles.pitch), float64(angles.heading), nil
}

// GetQuaternion returns the current orientation as a quaternion [w, x, y, z].
func (p *PhidgetSpatial) GetQuaternion() ([]float64, error) {
	var q C.PhidgetSpatial_SpatialQuaternion
	if cerr := C.PhidgetSpatial_getQuaternion(p.handle, &q); cerr != C.EPHIDGET_OK {
		return nil, p.phidgetError(cerr)
	}
	return []float64{float64(q.w), float64(q.x), float64(q.y), float64(q.z)}, nil
}

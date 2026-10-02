#ifndef GOPHIDGETS_H
#define GOPHIDGETS_H

#include <stdint.h>

/*
 * Callback function pointer types used by all device files.
 * Defined once here so each file doesn't need its own typedef.
 */
typedef void (*phidget_double_fcn)   (void*, void*, double);
typedef void (*phidget_motion_fcn)   (void*, void*, const double*, double);
typedef void (*phidget_quaternion_fcn)(void*, void*, const double*, double);
typedef void (*phidget_void_fcn)     (void*, void*);
typedef void (*phidget_encoder_fcn)  (void*, void*, int, double, int);
typedef void (*phidget_count_fcn)    (void*, void*, uint64_t, double);
typedef void (*phidget_twofloat_fcn) (void*, void*, double, double);
typedef void (*phidget_threefloat_fcn)(void*, void*, double, double, double);
typedef void (*phidget_spatial_fcn)  (void*, void*, const double*, const double*, const double*, double);
typedef void (*phidget_sound_fcn)    (void*, void*, double, double, double, const double*);

typedef void (*phidget_state_fcn)  (void*, void*, int);
typedef void (*phidget_int_fcn)    (void*, void*, int);
typedef void (*phidget_uint32_fcn) (void*, void*, uint32_t);
typedef void (*phidget_error_fcn)  (void*, void*, int, const char*);
typedef void (*phidget_rfid_fcn)   (void*, void*, const char*, int);

/*
 * Forward declarations of C shim functions defined in cdefs.go.
 * Device files reference these as function pointers for callbacks.
 */
void ccallback(void*, void*, double);
void cmotioncallback(void*, void*, const double*, double);
void cquaternioncallback(void*, void*, const double*, double);
void cvoidcallback(void*, void*);
void cencodercallback(void*, void*, int, double, int);
void ccountcallback(void*, void*, uint64_t, double);
void ctwofloatcallback(void*, void*, double, double);
void cthreefloatcallback(void*, void*, double, double, double);
void cspatialcallback(void*, void*, const double*, const double*, const double*, double);
void csoundcallback(void*, void*, double, double, double, const double*);
void cstatecallback(void*, void*, int);
void cintcallback(void*, void*, int);
void cuint32callback(void*, void*, uint32_t);
void cerrorcallback(void*, void*, int, const char*);
void crfidtagcallback(void*, void*, const char*, int);
void cmanager_detach_callback(PhidgetManagerHandle, void*, PhidgetHandle);

#endif /* GOPHIDGETS_H */

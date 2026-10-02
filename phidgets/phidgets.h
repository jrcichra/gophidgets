#ifndef GOPHIDGETS_H
#define GOPHIDGETS_H

#include <stdint.h>
#include <phidget22.h>

/*
 * Callback function pointer types used by all device files.
 * Defined once here so each file doesn't need its own typedef.
 */
typedef void (*phidget_double_fcn)   (void*, void*, double);
typedef void (*phidget_motion_fcn)   (void*, void*, const double*, double);
typedef void (*phidget_void_fcn)     (void*, void*);
typedef void (*phidget_encoder_fcn)  (void*, void*, int, double, int);
typedef void (*phidget_count_fcn)    (void*, void*, uint64_t, double);
typedef void (*phidget_twofloat_fcn) (void*, void*, double, double);
typedef void (*phidget_threefloat_fcn)(void*, void*, double, double, double);
typedef void (*phidget_spatial_fcn)  (void*, void*, const double*, const double*, const double*, double);
typedef void (*phidget_sound_fcn)    (void*, void*, double, double, double, const double*);

typedef void (*phidget_state_fcn)  (void*, void*, int);
typedef void (*phidget_uint32_fcn) (void*, void*, uint32_t);
typedef void (*phidget_error_fcn)  (void*, void*, int, const char*);
typedef void (*phidget_rfid_fcn)   (void*, void*, const char*, int);
typedef void (*phidget_kv_fcn)     (void*, void*, const char*, const char*);
typedef void (*phidget_key_fcn)    (void*, void*, const char*);
typedef void (*phidget_ircode_fcn) (void*, void*, const char*, uint32_t, int);
typedef void (*phidget_irraw_fcn)  (void*, void*, const uint32_t*, size_t);
typedef void (*phidget_irlearn_fcn)(void*, void*, const char*, void*);
typedef void (*phidget_prop_fcn)   (void*, void*, const char*);
typedef void (*phidget_manager_fcn)(PhidgetManagerHandle, void*, PhidgetHandle);

/*
 * Go functions exported via //export in common.go and manager.go, passed
 * directly to libphidget22 as handlers (cast to the typedefs above).
 * Signatures match the cgo-generated _cgo_export.h.
 */
void callback(void*, void*, double);
void motioncallback(void*, void*, double*, double);
void quaternioncallback(void*, void*, double*, double);
void voidcallback(void*, void*);
void encodercallback(void*, void*, int, double, int);
void countcallback(void*, void*, uint64_t, double);
void twofloatcallback(void*, void*, double, double);
void threefloatcallback(void*, void*, double, double, double);
void spatialcallback(void*, void*, double*, double*, double*, double);
void soundcallback(void*, void*, double, double, double, double*);
void statecallback(void*, void*, int);
void uint32callback(void*, void*, uint32_t);
void errorcallback(void*, void*, int, char*);
void rfidtagcallback(void*, void*, char*, int);
void dictkvcallback(void*, void*, char*, char*);
void dictkeycallback(void*, void*, char*);
void ircodecallback(void*, void*, char*, uint32_t, int);
void irrawcallback(void*, void*, uint32_t*, size_t);
void irlearncallback(void*, void*, char*, void*);
void propcallback(void*, void*, char*);
void attach_handler(PhidgetManagerHandle, void*, PhidgetHandle);
void manager_detach_handler(PhidgetManagerHandle, void*, PhidgetHandle);

/* handlectx turns a runtime/cgo.Handle into the void* ctx libphidget22 expects. */
static inline void* handlectx(uintptr_t h) { return (void*)h; }

#endif /* GOPHIDGETS_H */

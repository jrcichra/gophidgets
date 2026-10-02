package phidgets

/*
#include <stdlib.h>
#include <phidget22.h>
#include "phidgets.h"

// Single float64 callback
void callback(void*, void*, double);
void ccallback(void* handle, void* ctx, double b) {
  callback(handle, ctx, b);
}

// Sound sensor callback
void soundcallback(void*, void*, double, double, double, const double[10]);
void csoundcallback(void* handle, void* ctx, double dB, double dBA, double dBC, const double octaves[]) {
  soundcallback(handle, ctx, dB, dBA, dBC, octaves);
}

// Motion callback: 3-element double array + timestamp (accelerometer, gyroscope, magnetometer)
void motioncallback(void*, void*, double*, double);
void cmotioncallback(void* handle, void* ctx, const double* arr, double timestamp) {
  motioncallback(handle, ctx, (double*)arr, timestamp);
}

// Quaternion callback: 4-element double array + timestamp (spatial algorithm)
void quaternioncallback(void*, void*, double*, double);
void cquaternioncallback(void* handle, void* ctx, const double* arr, double timestamp) {
  quaternioncallback(handle, ctx, (double*)arr, timestamp);
}

// Void callback: no data arguments (stepper stopped, touch end)
void voidcallback(void*, void*);
void cvoidcallback(void* handle, void* ctx) {
  voidcallback(handle, ctx);
}

// Encoder callback: positionChange, timeChange, indexTriggered
void encodercallback(void*, void*, int, double, int);
void cencodercallback(void* handle, void* ctx, int positionChange, double timeChange, int indexTriggered) {
  encodercallback(handle, ctx, positionChange, timeChange, indexTriggered);
}

// Count callback: counts (uint64), timeChange
void countcallback(void*, void*, uint64_t, double);
void ccountcallback(void* handle, void* ctx, uint64_t counts, double timeChange) {
  countcallback(handle, ctx, counts, timeChange);
}

// Two-float callback: two double values (GPS heading: heading + velocity)
void twofloatcallback(void*, void*, double, double);
void ctwofloatcallback(void* handle, void* ctx, double a, double b) {
  twofloatcallback(handle, ctx, a, b);
}

// Three-float callback: three double values (GPS position: lat, lon, alt)
void threefloatcallback(void*, void*, double, double, double);
void cthreefloatcallback(void* handle, void* ctx, double a, double b, double c) {
  threefloatcallback(handle, ctx, a, b, c);
}

// Spatial data callback: three 3-element arrays + timestamp
void spatialcallback(void*, void*, double*, double*, double*, double);
void cspatialcallback(void* handle, void* ctx, const double* accel, const double* angularRate, const double* magneticField, double timestamp) {
  spatialcallback(handle, ctx, (double*)accel, (double*)angularRate, (double*)magneticField, timestamp);
}

// State callback: boolean state as int (digital input, stepper engaged, ...)
void statecallback(void*, void*, int);
void cstatecallback(void* handle, void* ctx, int state) {
  statecallback(handle, ctx, state);
}

// Uint32 callback: raw uint32 value (distance sensor distance, ...)
void uint32callback(void*, void*, uint32_t);
void cuint32callback(void* handle, void* ctx, uint32_t value) {
  uint32callback(handle, ctx, value);
}

// Channel error callback: EEPHIDGET_* event code + error message
void errorcallback(void*, void*, int, const char*);
void cerrorcallback(void* handle, void* ctx, int code, const char* message) {
  errorcallback(handle, ctx, code, message);
}

// RFID tag / tag-lost callback: tag data string + protocol
void rfidtagcallback(void*, void*, const char*, int);
void crfidtagcallback(void* handle, void* ctx, const char* tag, int protocol) {
  rfidtagcallback(handle, ctx, tag, protocol);
}

// Manager detach callback
void manager_detach_handler(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel);
void cmanager_detach_callback(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel) {
  manager_detach_handler(man, ctx, channel);
}

// Manager attach callback
void attach_handler(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel);
void cattach_callback(PhidgetManagerHandle man, void *ctx, PhidgetHandle channel) {
  attach_handler(man, ctx, channel);
}
*/
import "C"

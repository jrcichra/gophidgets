package phidgets

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L .
#include <stdlib.h>
void callback(void*, void*, double);
void ccallback(void* handle, void* ctx, double b) {
  callback(handle,ctx,b);
}
void soundcallback(void*, void*, double,double,double, const double[10]);
void csoundcallback(void* handle, void* ctx, double dB, double dBA, double dBC, const double octaves[]) {
  soundcallback(handle,ctx,dB,dBA,dBC,octaves);
}
*/
import "C"

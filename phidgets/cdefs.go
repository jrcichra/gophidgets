package phidgets

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L .
#include <stdlib.h>
void callback(void*, void*, double);
void ccallback(void* handle, void* ctx, double b) {
  callback(handle,ctx,b);
}
*/
import "C"

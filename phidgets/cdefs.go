package phidgets

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L .
#include <stdlib.h>
typedef void (*callback_fcn)(double i);
void ccallback(void* handle, void* ctx, double b) {
  callback(b);
}
*/
import "C"

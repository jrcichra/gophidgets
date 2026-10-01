package phidgets

import (
	"testing"

	gopointer "github.com/mattn/go-pointer"
)

// Handlers pass the Go func itself as the C context; voidcallback must restore and call it.
func TestVoidCallbackCallsFunc(t *testing.T) {
	called := false
	ctx := gopointer.Save(func() { called = true })
	defer gopointer.Unref(ctx)
	voidcallback(nil, ctx)
	if !called {
		t.Fatal("callback not invoked")
	}
}

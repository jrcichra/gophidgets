package phidgets

import "testing"

// Handlers pass the Go func itself as the C context; voidcallback must restore and call it.
func TestVoidCallbackCallsFunc(t *testing.T) {
	called := false
	var p phidget
	ctx := p.save(func() { called = true })
	defer p.ctxs[0].Delete()
	voidcallback(nil, ctx)
	if !called {
		t.Fatal("callback not invoked")
	}
}

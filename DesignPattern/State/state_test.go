package State

import "testing"

func TestContext_Handle(t *testing.T) {
	context := NewContext()
	context.Handle()
	context.Handle()
	context.Handle()
	context.Handle()
	context.Handle()
}

package Abstract_Factory

import "testing"

func TestNewFactory(t *testing.T) {
	factory := NewFactory("mouse")
	AsusMouse := factory.produceMouse("Asus")
	AsusMouse.click()
}

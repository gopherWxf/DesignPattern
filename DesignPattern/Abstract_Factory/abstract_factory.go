package Abstract_Factory

import "fmt"

type mouse interface {
	click()
}
type AsusMouse struct {
}

func (a *AsusMouse) click() {
	fmt.Println("Asus Mouse Check screen")
}

type LogitechMouse struct {
}

func (l *LogitechMouse) click() {
	fmt.Println("Logitech Mouse Check screen")
}

type keyboard interface {
	press()
}
type AsusKeyboard struct {
}

func (a *AsusKeyboard) press() {
	fmt.Println("Asus keyboard press")
}

type LogitechKeyboard struct {
}

func (l *LogitechKeyboard) press() {
	fmt.Println("Logitech Keyboard press")
}

type AbstractFactory interface {
	produceMouse(brand string) mouse
	produceKeyboard(brand string) keyboard
}
type MouseFactory struct {
}

func (m *MouseFactory) produceMouse(brand string) mouse {
	switch brand {
	case "Asus":
		return &AsusMouse{}
	case "Logitech":
		return &LogitechMouse{}
	}
	return nil
}

func (m *MouseFactory) produceKeyboard(brand string) keyboard {
	return nil
}

type KeyboardFactory struct {
}

func (k KeyboardFactory) produceMouse(brand string) mouse {
	return nil
}

func (k KeyboardFactory) produceKeyboard(brand string) keyboard {
	switch brand {
	case "Asus":
		return &AsusKeyboard{}
	case "Logitech":
		return &LogitechKeyboard{}
	}
	return nil
}
func NewFactory(factory string) AbstractFactory {
	switch factory {
	case "mouse":
		return &MouseFactory{}
	case "keyboard":
		return &KeyboardFactory{}
	}
	return nil
}

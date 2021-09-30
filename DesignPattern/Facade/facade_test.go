package Facade

import (
	"testing"
)

func TestCarFacade_CreateCompleteCar(t *testing.T) {
	car:=NewCarFacade()
	car.CreateCompleteCar()
	car.MethodA()
	car.MethodB()
}

package Decorator

import "testing"

func TestDecorator_Operation(t *testing.T) {
	objectA := &ConcreteComponent{}
	cdA := &DecoratorA{}
	cdB := &DecoratorB{}
	cdA.SetComponent(objectA)
	cdB.SetComponent(cdA)
	cdB.Operation()
}

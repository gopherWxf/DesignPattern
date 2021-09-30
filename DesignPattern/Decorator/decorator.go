package Decorator

import "fmt"

type Component interface {
	Operation()
}
type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() {
	fmt.Println("具体的对象开始操作...")
}

type Decorator struct {
	cc Component
}

func (d *Decorator) SetComponent(c Component) {
	d.cc = c
}

func (d *Decorator) Operation() {
	if d.cc != nil {
		d.cc.Operation()
	}
}

type DecoratorA struct {
	Decorator
}

func (d *DecoratorA) Operation() {
	d.cc.Operation()
	d.IndependentMethod()
}
func (d *DecoratorA) IndependentMethod() {
	fmt.Println("装饰A扩展的方法")
}

type DecoratorB struct {
	Decorator
}

func (d *DecoratorB) Operation() {
	d.cc.Operation()
	fmt.Println(d.String())
}
func (d *DecoratorB) String() string {
	return "装饰B扩展的方法"
}

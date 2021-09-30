package State

import "fmt"

type Context struct {
	state State
}

func NewContext() *Context {
	return &Context{state: &ConcreteStateA{}}
}

func (c *Context) SetState(state State) {
	c.state = state
}
func (c *Context) Handle() {
	c.state.Handler(c)
}

type State interface {
	Handler(*Context)
}

type ConcreteStateA struct {
}

func (c *ConcreteStateA) Handler(context *Context) {
	fmt.Println("当前状态是A")
	context.SetState(&ConcreteStateB{})
}

type ConcreteStateB struct {
}

func (c *ConcreteStateB) Handler(context *Context) {
	fmt.Println("当前状态是B")
	context.SetState(&ConcreteStateC{})

}

type ConcreteStateC struct {
}

func (c *ConcreteStateC) Handler(context *Context) {
	fmt.Println("当前状态是C")
	context.SetState(&ConcreteStateA{})
}

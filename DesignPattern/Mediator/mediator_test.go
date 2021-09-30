package Mediator

import "testing"

func TestConcreteMediator_Send(t *testing.T) {
	concreteMediator := &ConcreteMediator{}
	c1 := &ConcreteColleague1{}
	c2 := &ConcreteColleague2{}
	c1.SetMediator(concreteMediator)
	c2.SetMediator(concreteMediator)
	concreteMediator.c1 = c1
	concreteMediator.c2 = c2
	c1.Send("吃饭了吗c2")
	c2.Send("我吃过了c1")
}

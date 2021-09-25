package Abstract_Factory

import "fmt"

type Lunch interface {
	Cook()
}
type Rise struct {
}

func (r *Rise) Cook() {
	fmt.Println("it is Rise.")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("it is Tomato.")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}
type SimpleLunchFactory struct {
}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}

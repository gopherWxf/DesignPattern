package Builder

import (
	"testing"
)

func TestDirector_Construct(t *testing.T) {
	b1 := &ConcreteBuilder1{product: Product{parts: make([]string, 0)}}
	b2 := &ConcreteBuilder2{product: Product{parts: make([]string, 0)}}
	director := &Director{}
	director.Construct(b1)
	director.Construct(b2)
	p1 := b1.GetResult()
	p2 := b2.GetResult()
	p1.ShowParts()
	p2.ShowParts()
}

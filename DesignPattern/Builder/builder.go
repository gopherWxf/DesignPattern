package Builder

import "fmt"

type Product struct {
	parts []string
}

func (p *Product) AddParts(s string) {
	p.parts = append(p.parts, s)
}
func (p *Product) ShowParts() {
	fmt.Println(p.parts)
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetResult() Product
}
type ConcreteBuilder1 struct {
	product Product
}

func (c *ConcreteBuilder1) BuildPartA() {
	c.product.AddParts("人的头")
}

func (c *ConcreteBuilder1) BuildPartB() {
	c.product.AddParts("两只手")
}

func (c *ConcreteBuilder1) BuildPartC() {
	c.product.AddParts("两条腿")
}

func (c *ConcreteBuilder1) GetResult() Product {
	return c.product
}

type ConcreteBuilder2 struct {
	product Product
}

func (c *ConcreteBuilder2) BuildPartA() {
	c.product.AddParts("狗的头")
}

func (c *ConcreteBuilder2) BuildPartB() {
	c.product.AddParts("狗的爪子")
}

func (c *ConcreteBuilder2) BuildPartC() {
	c.product.AddParts("狗的腿")

}

func (c *ConcreteBuilder2) GetResult() Product {
	return c.product
}

type Director struct {
}

func (d *Director) Construct(builder Builder) {
	builder.BuildPartA()
	builder.BuildPartB()
	builder.BuildPartC()
}

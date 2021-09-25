package Builder

type Builder interface {
	Build()
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) Director {
	return Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.Build()
}

type ConcreteBuilder struct {
	built bool
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{built: false}
}

func (b *ConcreteBuilder) Build() {
	b.built = true
}

type Product struct {
	built bool
}

func (b *ConcreteBuilder) GetResult() Product {
	return Product{built: b.built}
}

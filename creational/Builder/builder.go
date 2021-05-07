package Builder
// 建造者模式

type Builder interface {
	Build()
}

type Director struct {
	builder Builder
}

func newDirector(b Builder) Director {
	return Director{
		builder: b,
	}
}

func (e *Director) Construct ()  {
	e.builder.Build()
}

type ConcreteBuilder struct {
	built bool
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{false}
}

func (c *ConcreteBuilder) Build()  {
	c.built = true
}

type Product struct {
	Build bool
}

func (c *ConcreteBuilder) GetResult() Product {
	return Product{c.built}
}
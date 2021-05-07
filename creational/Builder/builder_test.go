package Builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	builder := NewConcreteBuilder()

	director := newDirector(&builder)

	director.Construct()

	product := builder.GetResult()

	fmt.Println(product.Build)
}

func TestDirector1_GetPerson(t *testing.T) {
	old := NewOlderBuilder()
	director1 := NewDirector1(&old)
	fmt.Println(director1.GetPerson())

	child := NewChildBuilder()

	director1.builder1 = &child

	fmt.Println(director1.GetPerson())

}

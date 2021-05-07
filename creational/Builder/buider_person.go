package Builder
// 建造者模式

type Builder1 interface {
	GetPerson()
}

type Person struct {
	age int
	spend string
	knowledge string
}

type OlderBuilder struct {
	Person
}
func NewOlderBuilder() OlderBuilder {
	old := OlderBuilder{}
	old.setAge()
	old.setSpend()
	old.setKnowledge()
	return old
}
func (o *OlderBuilder) setAge() {
	o.age = 70
}

func (o *OlderBuilder) setSpend() {
	o.spend = "low"
}
func (o *OlderBuilder) setKnowledge() {
	o.knowledge = "more"
}
func (o *OlderBuilder) GetPerson() {
}

type ChildBuilder struct {
	Person
}

func (c *ChildBuilder) setAge() {
	c.age = 10
}

func (c *ChildBuilder) setSpend() {
	c.spend = "fast"
}
func (c *ChildBuilder) setKnowledge() {
	c.knowledge = "poor"
}
func (o *ChildBuilder) GetPerson() {

}
func NewChildBuilder() ChildBuilder {
	child := ChildBuilder{}
	child.setAge()
	child.setSpend()
	child.setKnowledge()
	return child
}
type Director1 struct {
	builder1 Builder1
}
func NewDirector1(b Builder1) Director1 {
	return Director1{
		builder1: b,
	}
}
func (d *Director1) GetPerson() Builder1 {
	return d.builder1
}
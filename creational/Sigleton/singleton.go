package Sigleton
// 单例模式
import "sync"

var (
	p *Pet
	once sync.Once
)

func init()  {
	once.Do(func() {
		p =&Pet{}
	})
}

type Pet struct {
	name string
	age int
	mux sync.Mutex
}

func GetInstance() *Pet {
	return p
}

func (p *Pet) SetName(n string)  {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.name = n
}

func (p *Pet) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}
func (p *Pet) IncrementAge()  {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.age ++
	p.mux.Unlock()
}

func (p *Pet) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
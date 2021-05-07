package Factory
// 工厂方法模式
import "fmt"

type Factory interface {
	GetGameName()
}

type Warcraft struct {

}

func (w *Warcraft) GetGameName() {
	fmt.Println("World of Warcraft")
}

type Dota struct {

}

func (d *Dota) GetGameName()  {
	fmt.Println("Dota")
}

func NewFactory(s string) Factory {
	switch s {
	case "w":
		return &Warcraft{}
	case "d":
		return &Dota{}
	}
	return nil
}
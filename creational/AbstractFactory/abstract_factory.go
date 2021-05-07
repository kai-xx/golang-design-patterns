package AbstractFactory
// 抽象工厂模式

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

type GameFactory interface {
	GetOnlineGame() Factory
	GetStandAloneGame() Factory
}

type simpleGameFactory struct {

}

func (s *simpleGameFactory) GetOnlineGame() Factory {
	return &Warcraft{}
}
func (s *simpleGameFactory) GetStandAloneGame() Factory {
	return &Dota{}
}
func NewSimpleGameFactory() GameFactory {
	return &simpleGameFactory{}
}


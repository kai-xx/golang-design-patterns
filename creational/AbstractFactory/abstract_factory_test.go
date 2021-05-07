package AbstractFactory

import "testing"

func TestNewSimpleGameFactory(t *testing.T) {
	game := NewSimpleGameFactory()
	game.GetOnlineGame().GetGameName()
	game.GetStandAloneGame().GetGameName()
}
package Factory

import "testing"

func TestNewFactory(t *testing.T) {
	NewFactory("d").GetGameName()
	NewFactory("w").GetGameName()
}

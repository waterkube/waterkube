package game

import "testing"

func TestGenerate(t *testing.T) {
	game := New(nil)
	game.Generate()

	if len(game.Grids) != 100 {
		t.Error("missing grids")
	}
}

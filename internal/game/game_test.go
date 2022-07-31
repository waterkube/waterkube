package game

import "testing"

func TestGenerate(t *testing.T) {
	game := New(nil, nil, nil)
	game.Generate()

	if game.Player == nil {
		t.Error("missing player")
	}

	if len(game.Grids) != 100 {
		t.Error("missing grids")
	}
}

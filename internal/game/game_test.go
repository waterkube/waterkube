package game

import "testing"

func TestMapGenerate(t *testing.T) {
	game := New(nil, nil, nil)
	game.MapGenerate()

	if game.Player == nil {
		t.Error("missing player")
	}

	if len(game.Grids) != 100 {
		t.Error("missing grids")
	}
}

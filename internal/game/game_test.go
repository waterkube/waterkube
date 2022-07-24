package game

import "testing"

func TestCreate(t *testing.T) {
	game := New(nil)
	game.Create()

	if len(game.Grids) != 100 {
		t.Error("missing grids")
	}
}

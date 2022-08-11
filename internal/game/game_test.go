package game

import "testing"

func TestMapGenerate(t *testing.T) {
	game := New(nil, nil, nil)
	game.MapGenerate()

	if game.Player == nil {
		t.Error("missing player")
	}

	for _, grid := range game.Grids {
		if grid == nil {
			t.Errorf("missing grid: %s", grid.Name)
		}
	}
}

func TestIsValidGridName(t *testing.T) {
	game := New(nil, nil, nil)

	tests := []struct {
		name     string
		expected bool
		got      bool
	}{
		{"Short Name", false, game.isValidGridName("a")},
		{"Long Name", false, game.isValidGridName("abcd")},
		{"Only Chars", false, game.isValidGridName("ab")},
		{"Only Numbers", false, game.isValidGridName("11")},
		{"Lower", true, game.isValidGridName("a0")},
		{"Upper", true, game.isValidGridName("A0")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != tt.got {
				t.Errorf("expected: %v, got: %v", tt.expected, tt.got)
			}
		})
	}
}

func TestNewGrids(t *testing.T) {
	game := New(nil, nil, nil)
	grids := game.newGrids()

	if len(grids) != 100 {
		t.Error("missing grids")
	}
}

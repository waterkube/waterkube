package game

import (
	"strconv"

	"github.com/waterkube/waterkube/internal/models"
)

const (
	// ShallowCount number.
	ShallowCount = 60

	// DeepCount number.
	DeepCount = 40

	// LegendaryCount number.
	LegendaryCount = 1

	// CombinableCount number.
	CombinableCount = 2
)

var (
	// Cols slice.
	Cols = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
)

// Game type.
type Game struct {
	gridRepository models.GridRepository
	Grids          []*models.Grid
}

// New function.
func New(gridRepository models.GridRepository) *Game {
	return &Game{
		gridRepository: gridRepository,
	}
}

// Create function.
func (g *Game) Create() {
	g.Grids = make([]*models.Grid, 100)

	for row := 0; row < 10; row++ {
		for i, col := range Cols {
			grid := models.NewGrid()
			grid.Name = col + strconv.Itoa(row)

			g.Grids[row*10+i] = grid
		}
	}
}

// Load function.
func (g *Game) Load() {
	// TODO
}

// Delete function.
func (g *Game) Delete() {
	// TODO
}

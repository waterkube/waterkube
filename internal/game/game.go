package game

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/waterkube/waterkube/internal/artifact"
	"github.com/waterkube/waterkube/internal/models"
)

const (
	// ShallowCount number.
	ShallowCount = 60

	// ShallowCombinableCount number.
	ShallowCombinableCount = 1

	// DeepCombinableCount number.
	DeepCombinableCount = 1

	// LegendaryCount number.
	LegendaryCount = 1
)

var (
	// Cols slice.
	Cols = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	// Rows slice.
	Rows = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
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

// Generate function.
func (g *Game) Generate() {
	cellCount := len(Cols) * len(Rows)
	g.Grids = make([]*models.Grid, cellCount)

	var combinable *artifact.Combinable
	combinableCount := ShallowCombinableCount

	for i := 0; i < ShallowCount; i++ {
		g.Grids[i] = models.NewGrid()
		g.Grids[i].Type = models.Shallow

		if combinable != nil {
			g.Grids[i].ArtifactType = models.Combinable
			g.Grids[i].Artifact = combinable.Pair
			combinable = nil
		} else if combinableCount > 0 {
			g.Grids[i].ArtifactType = models.Combinable
			g.Grids[i].Artifact, combinable = randFromMap(artifact.ShallowCombinable)
			combinableCount--
		} else {
			g.Grids[i].ArtifactType = models.Unique
			g.Grids[i].Artifact, _ = randFromMap(artifact.ShallowUnique)
		}
	}

	combinable = nil
	combinableCount = DeepCombinableCount
	legendaryCount := LegendaryCount

	for i := ShallowCount; i < cellCount; i++ {
		g.Grids[i] = models.NewGrid()
		g.Grids[i].Type = models.Deep

		if combinable != nil {
			g.Grids[i].ArtifactType = models.Combinable
			g.Grids[i].Artifact = combinable.Pair
			combinable = nil
		} else if combinableCount > 0 {
			g.Grids[i].ArtifactType = models.Combinable
			g.Grids[i].Artifact, combinable = randFromMap(artifact.DeepCombinable)
			combinableCount--
		} else if legendaryCount > 0 {
			g.Grids[i].ArtifactType = models.Legendary
			g.Grids[i].Artifact, _ = randFromMap(artifact.LegendaryUnique)
			legendaryCount--
		} else {
			g.Grids[i].ArtifactType = models.Unique
			g.Grids[i].Artifact, _ = randFromMap(artifact.DeepUnique)
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.Grids), func(i, j int) {
		g.Grids[i], g.Grids[j] = g.Grids[j], g.Grids[i]
	})

	for _, row := range Rows {
		for i, col := range Cols {
			g.Grids[row*len(Cols)+i].Name = col + strconv.Itoa(row)
		}
	}
}

// Create function.
func (g *Game) Create() {
	// TODO
}

// Load function.
func (g *Game) Load() {
	// TODO
}

// Delete function.
func (g *Game) Delete() {
	// TODO
}

func randFromMap[K comparable, V any](m map[K]V) (K, V) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(m))
	i := 0

	var k K
	var v V

	for k, v = range m {
		if r == i {
			return k, v
		}

		i++
	}

	return k, v
}

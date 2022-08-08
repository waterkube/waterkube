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

	// DiverPrice value.
	DiverPrice = 5

	// SubmarinePrice value.
	SubmarinePrice = 10
)

var (
	// Cols slice.
	Cols = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	// Rows slice.
	Rows = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

// Game type.
type Game struct {
	explorationRepository models.ExplorationRepository
	gridRepository        models.GridRepository
	playerRepository      models.PlayerRepository
	Player                *models.Player
	Grids                 []*models.Grid
}

// New function.
func New(
	explorationRepository models.ExplorationRepository,
	gridRepository models.GridRepository,
	playerRepository models.PlayerRepository,
) *Game {
	return &Game{
		explorationRepository: explorationRepository,
		gridRepository:        gridRepository,
		playerRepository:      playerRepository,
	}
}

// MapGenerate function.
func (g *Game) MapGenerate() {
	g.Player = models.NewPlayer()
	g.Grids = g.newGrids()

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

	for i := ShallowCount; i < len(Cols)*len(Rows); i++ {
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

// MapCreate function.
func (g *Game) MapCreate() error {
	g.MapGenerate()

	err := g.playerRepository.CreateOrUpdate(g.Player)
	if err != nil {
		return err
	}

	for _, grid := range g.Grids {
		err := g.gridRepository.CreateOrUpdate(grid)
		if err != nil {
			return err
		}
	}

	return nil
}

// MapLoad function.
func (g *Game) MapLoad() error {
	player, err := g.playerRepository.Find()
	if err != nil {
		return err
	}

	grids := g.newGrids()

	for _, row := range Rows {
		for i, col := range Cols {
			grid, err := g.gridRepository.Find(col + strconv.Itoa(row))
			if err != nil {
				return err
			}

			grids[row*len(Cols)+i] = grid
		}
	}

	g.Player = player
	g.Grids = grids

	return nil
}

// MapDelete function.
func (g *Game) MapDelete() error {
	err := g.playerRepository.Delete()
	if err != nil {
		return err
	}

	for _, row := range Rows {
		for _, col := range Cols {
			err := g.gridRepository.Delete(col + strconv.Itoa(row))
			if err != nil {
				return err
			}
		}
	}

	g.Player = nil
	g.Grids = nil

	return nil
}

// ArtifactCombine function.
func (g *Game) ArtifactCombine(nameA, nameB string) {
	// TODO
}

// ArtifactDonate function.
func (g *Game) ArtifactDonate(name string) {
	// TODO
}

// ArtifactSell function.
func (g *Game) ArtifactSell(name string) {
	// TODO
}

// DiverExplore function.
func (g *Game) DiverExplore(name string) {
	// TODO
}

// DiverHire function.
func (g *Game) DiverHire() error {
	err := g.MapLoad()
	if err != nil {
		return err
	}

	if g.Player.Money < DiverPrice {
		return ErrNoMoney
	}

	g.Player.Money -= DiverPrice
	g.Player.DiverCount++

	err = g.playerRepository.CreateOrUpdate(g.Player)
	if err != nil {
		return err
	}

	return nil
}

// SubmarineExplore function.
func (g *Game) SubmarineExplore(name string) {
	// TODO
}

// SubmarineBuy function.
func (g *Game) SubmarineBuy() error {
	err := g.MapLoad()
	if err != nil {
		return err
	}

	if g.Player.Money < SubmarinePrice {
		return ErrNoMoney
	}

	g.Player.Money -= SubmarinePrice
	g.Player.SubmarineCount++

	err = g.playerRepository.CreateOrUpdate(g.Player)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) newGrids() []*models.Grid {
	return make([]*models.Grid, len(Cols)*len(Rows))
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

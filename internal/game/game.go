package game

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/waterkube/waterkube/internal/artifact"
	"github.com/waterkube/waterkube/internal/models"
	"golang.org/x/exp/slices"
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
	DiverPrice = 50

	// SubmarinePrice value.
	SubmarinePrice = 100
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
	Explorations          []string
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
	var explorations []string

	for _, row := range Rows {
		for i, col := range Cols {
			grid, err := g.gridRepository.Find(col + strconv.Itoa(row))
			if err != nil {
				return err
			}

			grids[row*len(Cols)+i] = grid

			exploration, err := g.explorationRepository.Find(grid)
			if err != nil {
				continue
			}

			if exploration != nil {
				explorations = append(explorations, grid.Name)
			}
		}
	}

	g.Player = player
	g.Grids = grids
	g.Explorations = explorations

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

// MapProgress function.
func (g *Game) MapProgress() int {
	completed := len(g.Grids)

	for _, grid := range g.Grids {
		if grid.Status == models.Undiscovered {
			completed--
		}
	}

	return int(float32(completed) / float32(len(g.Grids)) * 100)
}

// FreeUnits function.
func (g *Game) FreeUnits() (int, int, int) {
	busyBoat := 0
	busyDiver := 0
	busySubmarine := 0

	for _, grid := range g.Grids {
		if !slices.Contains(g.Explorations, grid.Name) {
			continue
		}

		busyBoat++

		if grid.Type == models.Shallow {
			busyDiver++
		} else {
			busySubmarine++
		}
	}

	return g.Player.BoatCount - busyBoat,
		g.Player.DiverCount - busyDiver,
		g.Player.SubmarineCount - busySubmarine
}

// DiscoveredArtifacts function.
func (g *Game) DiscoveredArtifacts() []*artifact.Artifact {
	var discoveredArtifacts []*artifact.Artifact
	var artifactNames []string

	for _, grid := range g.Grids {
		if grid.Status != models.Discovered {
			continue
		}

		if slices.Contains(g.Explorations, grid.Name) {
			continue
		}

		if slices.Contains(artifactNames, grid.Artifact) {
			discoveredArtifacts[slices.Index(artifactNames, grid.Artifact)].Quantity++
		} else {
			discoveredArtifacts = append(discoveredArtifacts, &artifact.Artifact{
				Name:     grid.Artifact,
				Type:     string(grid.ArtifactType),
				Price:    g.artifactPrice(grid),
				Quantity: 1,
			})

			artifactNames = append(artifactNames, grid.Artifact)
		}
	}

	sort.SliceStable(discoveredArtifacts, func(i, j int) bool {
		return discoveredArtifacts[i].Name < discoveredArtifacts[j].Name
	})

	return discoveredArtifacts
}

// DonatedArtifacts function.
func (g *Game) DonatedArtifacts() []*artifact.Artifact {
	var donatedArtifacts []*artifact.Artifact
	var artifactNames []string

	for _, grid := range g.Grids {
		if grid.Status != models.Donated {
			continue
		}

		if slices.Contains(artifactNames, grid.Artifact) {
			donatedArtifacts[slices.Index(artifactNames, grid.Artifact)].Quantity++
		} else {
			donatedArtifacts = append(donatedArtifacts, &artifact.Artifact{
				Name:     grid.Artifact,
				Type:     string(grid.ArtifactType),
				Quantity: 1,
			})

			artifactNames = append(artifactNames, grid.Artifact)
		}
	}

	sort.SliceStable(donatedArtifacts, func(i, j int) bool {
		return donatedArtifacts[i].Name < donatedArtifacts[j].Name
	})

	return donatedArtifacts
}

// ArtifactCombine function.
func (g *Game) ArtifactCombine(artifactName1, artifactName2 string) (*models.Grid, error) {
	grid1 := g.artifactGrid(artifactName1)
	if grid1 == nil {
		return nil, ErrNoArtifact
	}

	if grid1.ArtifactType != models.Combinable {
		return nil, ErrInvalidGridArtifactType
	}

	grid2 := g.artifactGrid(artifactName2)
	if grid2 == nil {
		return nil, ErrNoArtifact
	}

	if grid2.ArtifactType != models.Combinable {
		return nil, ErrInvalidGridArtifactType
	}

	if grid1.Type != grid2.Type {
		return nil, ErrInvalidGridType
	}

	var combinable *artifact.Combinable

	if grid1.Type == models.Shallow {
		combinable = artifact.ShallowCombinable[grid1.Artifact]
	} else {
		combinable = artifact.DeepCombinable[grid1.Artifact]
	}

	if grid2.Artifact != combinable.Pair {
		return nil, ErrInvalidArtifactPair
	}

	grid1.Artifact = combinable.Result
	grid1.ArtifactType = models.Result

	err := g.gridRepository.CreateOrUpdate(grid1)
	if err != nil {
		return nil, err
	}

	grid2.Status = models.Combined

	err = g.gridRepository.CreateOrUpdate(grid2)
	if err != nil {
		return nil, err
	}

	return grid1, nil
}

// ArtifactDonate function.
func (g *Game) ArtifactDonate(artifactName string) (*models.Grid, error) {
	grid := g.artifactGrid(artifactName)
	if grid == nil {
		return nil, ErrNoArtifact
	}

	grid.Status = models.Donated

	err := g.gridRepository.CreateOrUpdate(grid)
	if err != nil {
		return nil, err
	}

	g.Player.BoatCount++

	err = g.playerRepository.CreateOrUpdate(g.Player)
	if err != nil {
		return nil, err
	}

	return grid, nil
}

// ArtifactSell function.
func (g *Game) ArtifactSell(artifactName string) (*models.Grid, error) {
	grid := g.artifactGrid(artifactName)
	if grid == nil {
		return nil, ErrNoArtifact
	}

	grid.Status = models.Sold

	err := g.gridRepository.CreateOrUpdate(grid)
	if err != nil {
		return nil, err
	}

	g.Player.Money += g.artifactPrice(grid)

	err = g.playerRepository.CreateOrUpdate(g.Player)
	if err != nil {
		return nil, err
	}

	return grid, nil
}

// DiverExplore function.
func (g *Game) DiverExplore(gridName string) (*models.Grid, error) {
	if !g.isValidGridName(gridName) {
		return nil, ErrInvalidGridName
	}

	freeBoat, freeDiver, _ := g.FreeUnits()

	if freeBoat == 0 {
		return nil, ErrNoBoat
	}

	if freeDiver == 0 {
		return nil, ErrNoDiver
	}

	grid := g.undiscoveredGrid(gridName)
	if grid == nil {
		return nil, ErrInvalidGridName
	}

	if grid.Type != models.Shallow {
		return nil, ErrInvalidGridType
	}

	if grid.Status != models.Undiscovered {
		return nil, ErrInvalidGridStatus
	}

	err := g.explorationRepository.Create(models.NewExploration(grid))
	if err != nil {
		return nil, err
	}

	grid.Status = models.Discovered

	err = g.gridRepository.CreateOrUpdate(grid)
	if err != nil {
		return nil, err
	}

	return grid, nil
}

// DiverHire function.
func (g *Game) DiverHire() error {
	if g.Player.Money < DiverPrice {
		return ErrNoMoney
	}

	g.Player.Money -= DiverPrice
	g.Player.DiverCount++

	err := g.playerRepository.CreateOrUpdate(g.Player)
	if err != nil {
		return err
	}

	return nil
}

// SubmarineExplore function.
func (g *Game) SubmarineExplore(gridName string) (*models.Grid, error) {
	if !g.isValidGridName(gridName) {
		return nil, ErrInvalidGridName
	}

	freeBoat, _, freeSubmarine := g.FreeUnits()

	if freeBoat == 0 {
		return nil, ErrNoBoat
	}

	if freeSubmarine == 0 {
		return nil, ErrNoSubmarine
	}

	grid := g.undiscoveredGrid(gridName)
	if grid == nil {
		return nil, ErrInvalidGridName
	}

	if grid.Type != models.Deep {
		return nil, ErrInvalidGridType
	}

	if grid.Status != models.Undiscovered {
		return nil, ErrInvalidGridStatus
	}

	err := g.explorationRepository.Create(models.NewExploration(grid))
	if err != nil {
		return nil, err
	}

	grid.Status = models.Discovered

	err = g.gridRepository.CreateOrUpdate(grid)
	if err != nil {
		return nil, err
	}

	return grid, nil
}

// SubmarineBuy function.
func (g *Game) SubmarineBuy() error {
	if g.Player.Money < SubmarinePrice {
		return ErrNoMoney
	}

	g.Player.Money -= SubmarinePrice
	g.Player.SubmarineCount++

	err := g.playerRepository.CreateOrUpdate(g.Player)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) newGrids() []*models.Grid {
	return make([]*models.Grid, len(Cols)*len(Rows))
}

func (g *Game) isValidGridName(gridName string) bool {
	if len(gridName) != 2 {
		return false
	}

	col := strings.ToUpper(gridName)[0:1]

	if !slices.Contains(Cols, col) {
		return false
	}

	row, err := strconv.Atoi(gridName[1:2])
	if err != nil {
		return false
	}

	if !slices.Contains(Rows, row) {
		return false
	}

	return true
}

func (g *Game) undiscoveredGrid(gridName string) *models.Grid {
	for _, grid := range g.Grids {
		if grid.Status != models.Undiscovered {
			continue
		}

		if !strings.EqualFold(grid.Name, gridName) {
			continue
		}

		return grid
	}

	return nil
}

func (g *Game) artifactGrid(artifactName string) *models.Grid {
	for _, grid := range g.Grids {
		if grid.Status != models.Discovered {
			continue
		}

		if slices.Contains(g.Explorations, grid.Name) {
			continue
		}

		if !strings.EqualFold(grid.Artifact, artifactName) {
			continue
		}

		return grid
	}

	return nil
}

func (g *Game) artifactPrice(grid *models.Grid) int {
	if grid.ArtifactType == models.Unique {
		if grid.Type == models.Shallow {
			return artifact.ShallowUnique[grid.Artifact].Price
		}

		return artifact.DeepUnique[grid.Artifact].Price
	}

	if grid.ArtifactType == models.Legendary {
		return artifact.LegendaryUnique[grid.Artifact].Price
	}

	if grid.ArtifactType == models.Combinable {
		if grid.Type == models.Shallow {
			return artifact.ShallowCombinable[grid.Artifact].Price
		}

		return artifact.DeepCombinable[grid.Artifact].Price
	}

	if grid.Type == models.Shallow {
		return artifact.ShallowResult[grid.Artifact].Price
	}

	return artifact.DeepResult[grid.Artifact].Price
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

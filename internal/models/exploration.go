package models

import "time"

const (
	// ExplorationTime duration.
	ExplorationTime = 15 * time.Second
)

// Exploration type.
type Exploration struct {
	Grid string `json:"grid" redis:"grid"`
}

// ExplorationRepository type.
type ExplorationRepository interface {
	Create(*Exploration) error
	Find(*Grid) (*Exploration, error)
}

// NewExploration function.
func NewExploration(grid *Grid) *Exploration {
	return &Exploration{
		Grid: grid.Name,
	}
}

package models

import "time"

const (
	// ExplorationTime duration.
	ExplorationTime = 10 * time.Second
)

// Exploration type.
type Exploration struct {
	Grid string `json:"grid" redis:"grid"`
}

// ExplorationRepository type.
type ExplorationRepository interface {
}

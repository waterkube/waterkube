package models

import "time"

const (
	// ExplorationTime duration.
	ExplorationTime = 10 * time.Second
)

type Exploration struct {
	Grid     string   `json:"grid" redis:"grid"`
	GridType GridType `json:"grid_type" redis:"grid_type"`
}

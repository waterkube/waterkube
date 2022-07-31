package models

// Player type.
type Player struct {
	Money          int `json:"money" redis:"money"`
	BoatCount      int `json:"boat_count" redis:"boat_count"`
	DiverCount     int `json:"diver_count" redis:"diver_count"`
	SubmarineCount int `json:"submarine_count" redis:"submarine_count"`
}

// PlayerRepository type.
type PlayerRepository interface {
	Create(*Player) error
	Find() (*Player, error)
	Update(*Player) error
	Delete() error
}

// NewPlayer function.
func NewPlayer() *Player {
	return &Player{
		Money:          0,
		BoatCount:      1,
		DiverCount:     1,
		SubmarineCount: 0,
	}
}

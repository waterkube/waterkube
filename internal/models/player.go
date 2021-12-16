package models

// Player type.
type Player struct {
	BoatCount      int `json:"boat_count" redis:"boat_count"`
	DiverCount     int `json:"diver_count" redis:"diver_count"`
	SubmarineCount int `json:"submarine_count" redis:"submarine_count"`
}

// PlayerRepository type.
type PlayerRepository interface {
}

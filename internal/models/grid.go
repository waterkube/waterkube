package models

// GridType type.
type GridType string

// GridStatus type.
type GridStatus string

const (
	// Shallow type.
	Shallow GridType = "shallow"

	// ShallowCount number.
	ShallowCount = 60

	// Deep type.
	Deep GridType = "deep"

	// DeepCount number.
	DeepCount = 40

	// Undiscovered status.
	Undiscovered GridStatus = "undiscovered"

	// Discovered status.
	Discovered GridStatus = "discovered"
)

// Grid type.
type Grid struct {
	Name     string     `json:"name" redis:"name"`
	Type     GridType   `json:"type" redis:"type"`
	Status   GridStatus `json:"status" redis:"status"`
	Artifact string     `json:"artifact" redis:"artifact"`
}

// GridRepository type.
type GridRepository interface {
}

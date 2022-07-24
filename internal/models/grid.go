package models

// GridType type.
type GridType string

// GridArtifactType type.
type GridArtifactType string

// GridStatus type.
type GridStatus string

const (
	// Shallow type.
	Shallow GridType = "shallow"

	// Deep type.
	Deep GridType = "deep"

	// Unique artifact type.
	Unique GridArtifactType = "unique"

	// Legendary artifact type.
	Legendary GridArtifactType = "legendary"

	// Combinable artifact type.
	Combinable GridArtifactType = "combinable"

	// Undiscovered status.
	Undiscovered GridStatus = "undiscovered"

	// Discovered status.
	Discovered GridStatus = "discovered"

	// Sold status.
	Sold GridStatus = "sold"
)

// Grid type.
type Grid struct {
	Name         string           `json:"name" redis:"name"`
	Type         GridType         `json:"type" redis:"type"`
	Artifact     string           `json:"artifact" redis:"artifact"`
	ArtifactType GridArtifactType `json:"artifact_type" redis:"artifact_type"`
	Status       GridStatus       `json:"status" redis:"status"`
}

// GridRepository type.
type GridRepository interface {
}

// NewGrid function.
func NewGrid() *Grid {
	return &Grid{
		Status: Undiscovered,
	}
}

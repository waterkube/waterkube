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

	// Result artifact type.
	Result GridArtifactType = "result"

	// Undiscovered status.
	Undiscovered GridStatus = "undiscovered"

	// Discovered status.
	Discovered GridStatus = "discovered"

	// Combined status.
	Combined GridStatus = "combined"

	// Donated status.
	Donated GridStatus = "donated"

	// Sold status.
	Sold GridStatus = "sold"
)

// Grid type.
type Grid struct {
	Name         string           `json:"name" redis:"name"`
	Type         GridType         `json:"type" redis:"type"`
	Artifact     string           `json:"-" redis:"artifact"`
	ArtifactType GridArtifactType `json:"-" redis:"artifact_type"`
	Status       GridStatus       `json:"status" redis:"status"`
}

// GridRepository type.
type GridRepository interface {
	CreateOrUpdate(*Grid) error
	Find(string) (*Grid, error)
	Delete(string) error
}

// NewGrid function.
func NewGrid() *Grid {
	return &Grid{
		Status: Undiscovered,
	}
}

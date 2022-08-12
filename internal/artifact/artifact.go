package artifact

// Artifact type.
type Artifact struct {
	Grid  string `json:"grid"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Price int    `json:"price"`
}

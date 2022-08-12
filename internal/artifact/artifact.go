package artifact

// Artifact type.
type Artifact struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

package models

// Command type.
type Command struct {
	ID        int64     `json:"id" redis:"id"`
	Value     string    `json:"value" redis:"value"`
	CreatedAt Timestamp `json:"created_at" redis:"created_at"`
}

// CommandRepository type.
type CommandRepository interface {
}

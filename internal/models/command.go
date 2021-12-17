package models

// Command type.
type Command struct {
	ID        int64     `json:"id" redis:"id"`
	GroupName string    `json:"group_name" redis:"group_name"`
	Name      string    `json:"name" redis:"name"`
	Arguments []string  `json:"arguments" redis:"arguments"`
	CreatedAt Timestamp `json:"created_at" redis:"created_at"`
}

// CommandRepository type.
type CommandRepository interface {
}

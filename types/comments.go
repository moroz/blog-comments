package types

import (
	"uuid"
)

type Comment struct {
	ID         uuid.UUID     `json:"id"`
	PostPath   string        `json:"post_path"`
	Author     string        `json:"author"`
	Body       string        `json:"body"`
	Website    *string       `json:"website,omitempty"`
	InsertedAt UnixTimestamp `json:"inserted_at"`
}

type CommentParams struct {
	PostPath string
	Author   string
	Body     string
	Website  string
}

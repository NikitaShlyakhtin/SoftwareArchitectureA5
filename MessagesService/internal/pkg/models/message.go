package models

import "github.com/google/uuid"

// Message represents a feed message
type Message struct {
	ID       uuid.UUID `json:"id"` // UUIDv7
	Username string    `json:"username"`
	Content  string    `json:"content"`
	IsLiked  bool      `json:"is_liked"`
}

// NewMessage creates a new Message instance
func NewMessage(
	id uuid.UUID,
	username string,
	content string,
	isLiked bool,
) *Message {
	return &Message{
		ID:       id,
		Username: username,
		Content:  content,
		IsLiked:  isLiked,
	}
}

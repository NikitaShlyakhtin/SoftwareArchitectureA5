package types

import "github.com/google/uuid"

// Message represents a feed message
type Message struct {
	ID       uuid.UUID `db:"id" json:"id"` // UUIDv7
	Username string    `db:"username" json:"username"`
	Content  string    `db:"content" json:"content"`
	IsLiked  bool      `db:"is_liked" json:"is_liked"`
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

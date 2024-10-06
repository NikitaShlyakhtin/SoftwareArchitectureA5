package dependencies

import (
	"MessagesService/internal/pkg/types"
	"github.com/google/uuid"
)

// IStore defines the methods for working with storage
type IStore interface {
	Start() error
	Stop() error

	InsertMessage(msg *types.Message) (*types.Message, error)
	LikeMessage(id uuid.UUID) (*types.Message, error)
}

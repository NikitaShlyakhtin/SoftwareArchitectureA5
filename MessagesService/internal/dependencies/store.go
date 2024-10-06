package dependencies

import (
	"MessagesService/internal/pkg/types"
	"context"
	"github.com/google/uuid"
)

// IStore defines the methods for working with storage
type IStore interface {
	InsertMessage(context.Context, *types.Message) (*types.Message, error)
	LikeMessage(context.Context, string, uuid.UUID) (*types.Message, error)
}

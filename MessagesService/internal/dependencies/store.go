package dependencies

import (
	"MessagesService/internal/pkg/models"
	"context"
	"github.com/google/uuid"
)

// IStore defines the methods for working with storage
type IStore interface {
	InsertMessage(context.Context, *models.Message) (*models.Message, error)
	LikeMessage(context.Context, string, uuid.UUID) (*models.Message, error)
}

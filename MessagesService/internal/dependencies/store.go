package dependencies

import (
	"MessagesService/internal/pkg/models"
	"context"
)

// IStore defines the methods for working with storage
type IStore interface {
	InsertMessage(context.Context, *models.Message) (*models.Message, error)
}

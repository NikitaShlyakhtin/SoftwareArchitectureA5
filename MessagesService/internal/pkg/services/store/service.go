package store

import (
	"MessagesService/internal/dependencies"
	"MessagesService/internal/pkg/models"
	"context"
	"errors"
	"go.uber.org/zap"
)

// Store implementation of IStore interface
type Store struct {
	Logger *zap.SugaredLogger
}

var _ dependencies.IStore = (*Store)(nil)

// NewStore creates a new instance of Store
func NewStore(l *zap.Logger) (dependencies.IStore, error) {
	if l == nil {
		return nil, errors.New("logger must be provided")
	}

	return &Store{
		Logger: l.Sugar(),
	}, nil
}

func (s *Store) InsertMessage(ctx context.Context, msg *models.Message) (*models.Message, error) {
	return nil, errors.New("not implemented")
}

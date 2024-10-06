package messages_service

import (
	"MessagesService/internal/dependencies"
	"errors"
	"go.uber.org/zap"
)

// MessageService implementation of IMessageService interface
type MessageService struct {
	Logger *zap.SugaredLogger
}

var _ dependencies.IMessageService = (*MessageService)(nil)

// NewMessageService creates a new instance of MessageService
func NewMessageService(l *zap.Logger) (dependencies.IMessageService, error) {
	if l == nil {
		return nil, errors.New("logger must be provided")
	}

	return &MessageService{
		Logger: l.Sugar(),
	}, nil
}

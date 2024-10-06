package app

import (
	"MessagesService/internal/dependencies"
	"go.uber.org/zap"
)

// Application holds the application state and dependencies
type Application struct {
	Logger         *zap.SugaredLogger
	MessageService dependencies.IMessageService
}

// NewApplication initializes a new Application instance
func NewApplication(
	logger *zap.Logger,
	messageService dependencies.IMessageService,
) *Application {
	return &Application{
		Logger:         logger.Sugar(),
		MessageService: messageService,
	}
}

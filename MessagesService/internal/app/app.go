package app

import (
	"MessagesService/internal/dependencies"
	"go.uber.org/zap"
)

// Application holds the application state and dependencies
type Application struct {
	Logger *zap.SugaredLogger
	Store  dependencies.IStore
}

// NewApplication initializes a new Application instance
func NewApplication(
	logger *zap.Logger,
	store dependencies.IStore,
) *Application {
	return &Application{
		Logger: logger.Sugar(),
		Store:  store,
	}
}

package server

import (
	"MessagesService/internal/app"
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.SugaredLogger
	echo   *echo.Echo
	app    *app.Application
}

func NewServer(app *app.Application) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	s := &Server{
		logger: app.Logger,
		app:    app,
		echo:   e,
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	s.echo.POST("/messages/create", s.app.CreateMessage())
	s.echo.PUT("/messages/like", s.app.LikeMessage())
}

func (s *Server) Start(address string) error {
	s.logger.Infof("Starting store on address: %v", address)

	return s.echo.Start(address)
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down store")

	return s.echo.Shutdown(ctx)
}

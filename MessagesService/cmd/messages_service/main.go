package main

import (
	"MessagesService/internal/app"
	"MessagesService/internal/pkg/services/messages_service"
	"MessagesService/internal/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(getFxOptions()).Run()
}

func getFxOptions() fx.Option {
	return fx.Options(
		fx.WithLogger(getEventLogger),
		fx.Provide(
			zap.NewDevelopment,
			messages_service.NewMessageService,
			app.NewApplication,
			server.NewServer,
		),
		fx.Invoke(startServer),
	)
}

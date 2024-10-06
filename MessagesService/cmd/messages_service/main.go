package main

import (
	"MessagesService/internal/app"
	"MessagesService/internal/dependencies"
	"MessagesService/internal/pkg/services/store"
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
			fx.Annotate(store.NewStore, fx.As(new(dependencies.IStore))),
			app.NewApplication,
			server.NewServer,
		),
		fx.Invoke(startServer, startStore),
	)
}

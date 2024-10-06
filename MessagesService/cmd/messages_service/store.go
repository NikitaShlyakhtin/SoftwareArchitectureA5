package main

import (
	"MessagesService/internal/dependencies"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func startStore(lc fx.Lifecycle, s dependencies.IStore, l *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return s.Start()
		},
		OnStop: func(ctx context.Context) error {
			return s.Stop()
		},
	})
}

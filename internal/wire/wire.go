//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/devendrapratap307/go-crypto-service/internal/config"
	"github.com/devendrapratap307/go-crypto-service/internal/handler"
	"github.com/devendrapratap307/go-crypto-service/internal/keys"
	"github.com/devendrapratap307/go-crypto-service/internal/service"
	"github.com/google/wire"
)

// func InitializeApp() (*handler.RestHandler, *handler.GRPCServer, error) {
// 	wire.Build(
// 		config.Load,
// 		keys.NewKeyManager,
// 		service.NewCryptoService,
// 		handler.NewRestHandler,
// 		handler.NewGRPCServer,
// 	)
// 	// return &handler.RestHandler{}, &handler.GRPCServer{}, nil
// 	return nil, nil, nil
// }

type App struct {
	RestHandler *handler.RestHandler
	GRPCServer  *handler.GRPCServer
}

func NewApp(rest *handler.RestHandler, grpc *handler.GRPCServer) *App {
	return &App{RestHandler: rest, GRPCServer: grpc}
}

func InitializeApp() (*App, error) {
	wire.Build(
		config.Load,
		keys.NewKeyManager,
		service.NewCryptoService,
		handler.NewRestHandler,
		handler.NewGRPCServer,
		NewApp,
	)
	return nil, nil
}

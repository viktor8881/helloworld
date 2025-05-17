package main

import (
	"context"
	genServerEndps "github.com/viktor8881/helloworld/generated/http/server"
	"github.com/viktor8881/helloworld/inner/alive"
	"github.com/viktor8881/helloworld/inner/mainpage"
	"github.com/viktor8881/service-utilities/http/server"
	"go.uber.org/zap"
	"net/http"
)

func RegisterRoutes(ctx context.Context, mux *http.ServeMux, logger *zap.Logger) []func() {
	// http endpoints
	aliveService := alive.NewService()
	mainService := mainpage.NewService()

	tr := server.NewTransport(mux)

	genServerEndps.MainPage(
		tr,
		server.DecodeRequest,
		mainService.MainPage,
		server.EncodeResponse,
		server.ErrorHandler,
		logger,
	)

	genServerEndps.AliveEndpoint(
		tr,
		server.DecodeRequest,
		aliveService.Alive,
		server.EncodeResponse,
		server.ErrorHandler,
		logger,
	)

	return []func(){}
}

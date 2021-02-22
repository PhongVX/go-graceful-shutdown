package app

import (
	"context"
	"graceful-shutdown/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	app struct {
		Server *http.Server
	}

	App interface {
		Start() error
		Stop(ctx context.Context) error
	}
)

const (
	ADDR = ":8081"
)

func New() App {
	router := mux.NewRouter()
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()
	handler.New(apiV1Router)

	httpServer := &http.Server{
		Addr:    ADDR,
		Handler: router,
	}
	return app{
		Server: httpServer,
	}
}

func (a app) Start() error {
	log.Printf("Server is listening at %s", ADDR)
	return a.Server.ListenAndServe()
}

func (a app) Stop(ctx context.Context) error {
	return a.Server.Shutdown(ctx)
}

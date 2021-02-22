package router

import (
	"graceful-shutdown/internal/handler"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()
	handler.New(apiV1Router)
	return router
}
package api

import (
	"github.com/gorilla/mux"
	"github.com/jolienai/unicorn-factory/middleware"
)

func NewRouter(handler *UnicornHandler) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthenticationMiddleware)

	var api = r.PathPrefix("/v1").Subrouter()

	api.HandleFunc("/unicorn/{id:[0-9]+}", handler.Get).Methods("GET")
	api.HandleFunc("/unicorn", handler.Create).Methods("POST")

	return r
}

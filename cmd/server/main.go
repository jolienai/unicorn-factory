package main

import (
	"log"
	"net/http"

	"github.com/jolienai/unicorn-factory/api"
	"github.com/jolienai/unicorn-factory/internal/unicorn"
)

func main() {
	serverAddr := ":8081"

	repository := unicorn.NewUnicornRepository()
	factory := unicorn.NewFactory()

	handler := api.NewUnicornHandler(repository, factory)

	router := api.NewRouter(handler)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

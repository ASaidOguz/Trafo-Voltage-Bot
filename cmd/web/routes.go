package main

import (
	"TrafoApp/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

// routes handles application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/send", http.HandlerFunc(handlers.SendData))

	return mux
}

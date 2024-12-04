package main

import (
	"fmt"
	"go-api/internal/middlewares"
	"go-api/internal/handlers/item"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middlewares.Authorization)

	// Routes
	r.Route("/item", func(r chi.Router) {
		r.Post("/list", item.List)
	})

	// Start the server
	fmt.Println("Starting Item server on :8001")
	http.ListenAndServe(":8001", r)
}
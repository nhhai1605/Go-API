package main

import (
	"fmt"
	"go-api/internal/handlers/item"
	"go-api/internal/middlewares"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        panic(err)
    }
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middlewares.Authorization)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Routes
	r.Route("/item", func(r chi.Router) {
		r.Post("/list", item.List)
	})

	// Start the server
	fmt.Println("Starting Item server on :8001")
	http.ListenAndServe(":8001", r)
}
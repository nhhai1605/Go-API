package main

import (
	"fmt"
	"go-api/internal/handlers/auth"
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
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	// Start the server
	fmt.Println("Starting Auth server on :8000")
	http.ListenAndServe(":8000", r)
}
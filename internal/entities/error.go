package entities

import (
	"net/http"
	"github.com/go-chi/render"
)

type Error struct {
	Code int
	Message string
}

func (e *Error) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)
	return nil
}
package core

import (
	"go-api/internal/entities"
	"github.com/go-chi/render"
)

func ErrRender(err entities.Error) render.Renderer {
	return &err;
}

func CreateError(code int, message string) *entities.Error {
	return &entities.Error{
		Code: code,
		Message: message,
	}
}
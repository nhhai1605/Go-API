package core

import (
	"go-api/internal/entities"
	"net/http"
	"github.com/go-chi/render"
)

func ErrInvalidRequest(err error) render.Renderer {
	return &entities.Error{
		Code: http.StatusBadRequest,
		Message: err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &entities.Error{
		Code: http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func ErrUnauthorized(err error) render.Renderer {
	errorMessage := "Unauthorized"
	if(err != nil) {
		errorMessage = err.Error()
	}
	return &entities.Error{
		Code: http.StatusUnauthorized,
		Message: errorMessage,
	}
}
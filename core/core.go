package core

import (
	"net/http"
	"github.com/go-chi/render"
)

type Error struct
{
	Code int
	Message string
}
func (e *Error) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &Error{
		Code: http.StatusBadRequest,
		Message: err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &Error{
		Code: http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func ErrUnauthorized(err error) render.Renderer {
	return &Error{
		Code: http.StatusUnauthorized,
		Message: err.Error(),
	}
}
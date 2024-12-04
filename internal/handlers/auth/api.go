package auth

import (
	"go-api/core"
	"go-api/internal/databases/auth"
	"go-api/internal/entities"
	"net/http"
	"github.com/go-chi/render"
)

func Login(w http.ResponseWriter, r *http.Request) {
	params := &entities.LoginRequest{}
	
	if  err := render.Bind(r, params); err != nil {
		render.Render(w, r, core.ErrRender(*core.CreateError(http.StatusBadRequest, "Invalid request payload")))
		return
	}
	token, err := auth.GetToken(*params)
	if err != nil {
		render.Render(w, r, core.ErrRender(*err))
		return
	}
	render.JSON(w, r, map[string]string{"token": token})
}
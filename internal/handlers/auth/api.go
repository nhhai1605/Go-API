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
	err := render.Bind(r, params);
	if  err != nil {
		render.Render(w, r, core.ErrInvalidRequest(err))
		return
	}
	db, err := auth.NewDatabase()
	if err != nil {
		render.Render(w, r, core.ErrRender(err))
		return
	}
	token, err := (*db).GetToken(*params)
	if err != nil {
		render.Render(w, r, core.ErrUnauthorized(err))
		return
	}
	render.JSON(w, r, map[string]string{"token": token})
}
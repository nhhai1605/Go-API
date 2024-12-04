package entities

import "net/http"

type LoginRequest struct {
	Username string
	Password string
}

func (a *LoginRequest) Bind(r *http.Request) error {
	return nil
}
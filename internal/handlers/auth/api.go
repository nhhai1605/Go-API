package auth

import (
	"go-api/internal/databases/auth"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	params := &auth.LoginRequest{}
	
	//get the request body from POST
	//parse the request body into the params struct

}
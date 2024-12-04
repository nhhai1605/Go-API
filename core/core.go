package core

import (
	"encoding/json"
	"net/http"
)

type Error struct
{
	Code int
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	err := Error{Code: code, Message: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}

var (
	NotFoundHandle = func(w http.ResponseWriter) {
		writeError(w, http.StatusNotFound, "Not Found")
	}
	InternalErrorHandle = func(w http.ResponseWriter, message string) {
		writeError(w, http.StatusInternalServerError, message)
	}
	UnauthorizedHandle = func(w http.ResponseWriter) {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
	}
)
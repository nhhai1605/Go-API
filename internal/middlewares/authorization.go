package middlewares

import (
	"fmt"
	"go-api/core"
	_ "go-api/internal/databases/auth"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/render"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get "token" from the request header
		token := r.Header.Get("token")
		if token == "" {
			render.Render(w, r, core.ErrRender(*core.CreateError(http.StatusBadRequest, "Token is required")))
			return
		}
		plainText, err := core.Decrypt(token)
		if err != nil {
			render.Render(w, r, core.ErrRender(*core.CreateError(http.StatusBadRequest, "Invalid token")))
			return
		}
		splitString := strings.Split(plainText, "~")
		if len(splitString) != 4 {
			render.Render(w, r, core.ErrRender(*core.CreateError(http.StatusBadRequest, "Invalid token")))
			return
		}
		expireTimeUnix, err := strconv.ParseInt(splitString[3], 10, 64)
		if err != nil {
			render.Render(w, r, core.ErrRender(*core.CreateError(http.StatusBadRequest, "Invalid token")))
			return
		}
		fmt.Println(expireTimeUnix)
		curTimeUnix := time.Now().UTC().Unix()
		deltaSeconds := curTimeUnix - expireTimeUnix
		expireSeconds, err := strconv.ParseInt(os.Getenv("TOKEN_EXPIRE_SECONDS"), 10, 64)
		if err != nil {
			render.Render(w, r, core.ErrRender(*core.CreateError(http.StatusInternalServerError, "Internal server error")))
			return
		}
		if deltaSeconds > expireSeconds {
			render.Render(w, r, core.ErrRender(*core.CreateError(http.StatusBadRequest, "Token expired")))
			return
		}
		next.ServeHTTP(w, r)
	})
}
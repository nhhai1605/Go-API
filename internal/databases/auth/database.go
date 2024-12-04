package auth

import (
	"database/sql"
	"fmt"
	"go-api/core"
	"go-api/internal/entities"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "postgres"
)


func GetToken(loginRequest entities.LoginRequest) (string, *entities.Error) {
	var token string = ""
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return token, core.CreateError(http.StatusInternalServerError, err.Error())
	}
	queryString := `SELECT * FROM auth.login($1, $2)`
    rows, err := db.Query(queryString, loginRequest.Username, loginRequest.Password)
	if err != nil {
		return token, core.CreateError(http.StatusBadRequest, err.Error())
	}
	cols, _ := rows.Columns()
	maps, err := core.DatabaseMapping(rows, cols)
	if err != nil {
		return token, core.CreateError(http.StatusInternalServerError, err.Error())
	}
	if err := rows.Err(); err != nil {
		return token, core.CreateError(http.StatusInternalServerError, err.Error())
	}
	type User struct {
		ID int 
		Username string 
		Password string 
	}
	var users []User
	err = core.StructMapping(maps, &users)
	fmt.Println(users)
	if err != nil {
		return token, core.CreateError(http.StatusInternalServerError, err.Error())
	}
    token, err = core.Encrypt(fmt.Sprintf("%d~%s~%s~%d", users[0].ID, users[0].Username, users[0].Password, time.Now().UTC().Unix()))
	if err != nil {
		return token, core.CreateError(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	defer db.Close()
	return token, nil
}


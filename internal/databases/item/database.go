package item

import (
	"database/sql"
	"fmt"
	"go-api/core"
	"go-api/internal/entities"
	"net/http"
	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "postgres"
)

func GetItemList() ([]entities.Item, *entities.Error) {
	var items []entities.Item
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return items, core.CreateError(http.StatusInternalServerError, err.Error())
	}
	queryString := `SELECT * FROM item.item`
    rows, err := db.Query(queryString)
	if err != nil {
		return items, core.CreateError(http.StatusBadRequest, err.Error())
	}
	cols, _ := rows.Columns()
	maps, err := core.DatabaseMapping(rows, cols)
	if err != nil {
		return items, core.CreateError(http.StatusInternalServerError, err.Error())
	}
	if err := rows.Err(); err != nil {
		return items, core.CreateError(http.StatusInternalServerError, err.Error())
	}
	err = core.StructMapping(maps, &items)
	if err != nil {
		return items, core.CreateError(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	defer db.Close()
	return items, nil
}

package item

import (
	"go-api/core"
	"go-api/internal/databases/item"
	"net/http"
	"github.com/go-chi/render"
)

func List(w http.ResponseWriter, r *http.Request) {
	itemList, err := item.GetItemList()
	if err != nil {
		render.Render(w, r, core.ErrRender(*err))
		return
	}
	render.JSON(w, r, itemList)
}
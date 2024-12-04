package entities

import "net/http"

type Item struct {
	ID int
	Name string
	Price float64
}

func (a *Item) Bind(r *http.Request) error {
	return nil
}
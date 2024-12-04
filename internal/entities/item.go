package entities

import "net/http"

type Item struct {
	ID int 
	Name string 
	Description string 
}

func (a *Item) Bind(r *http.Request) error {
	return nil
}
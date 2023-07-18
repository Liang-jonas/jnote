package DTO

import "github.com/Liang-jonas/golib/object"

type API struct {
	ID          int64   `json:"id"`
	Name        *string `json:"name"`
	Path        *string `json:"path"`
	Description *string `json:"description"`
}

func (a *API) String() string {
	return object.ToString(a)
}

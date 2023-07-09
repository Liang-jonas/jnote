package Entity

import "github.com/Liang-jonas/golib/db"

type Image struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Owner       int64  `json:"owner"`
	db.Model
}

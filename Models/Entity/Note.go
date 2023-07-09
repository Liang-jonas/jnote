package Entity

import "github.com/Liang-jonas/golib/db"

type Note struct {
	Title string `json:"title"`
	Owner int64  `json:"owner"`
	db.Model
}

package Entity

import "github.com/Liang-jonas/golib/db"

type WorkSpace struct {
	Name  string `json:"name"`
	Owner int64  `json:"owner"`
	db.Model
}

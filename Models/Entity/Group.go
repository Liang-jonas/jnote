package Entity

import "github.com/Liang-jonas/golib/db"

type Group struct {
	Name        string `json:"name`
	Description string `json:"description"`
	IsDel       bool   `json:"is-del"`
	db.Model
}

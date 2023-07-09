package Entity

import "github.com/Liang-jonas/golib/db"

type Theme struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Context     string `json:"context"`
	db.Model
}

package Entity

import "github.com/Liang-jonas/golib/db"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ThemeID  int64  `json:"theme-id"`
	GroupID  int64  `json:"group-id"`
	Disable  bool   `json:"disable"`
	Avatar   string `json:"avatar"`
	db.Model
}

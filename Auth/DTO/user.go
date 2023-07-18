package DTO

import "github.com/Liang-jonas/golib/object"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Group    string `json:"group"`
	GroupID  int64  `json:"group-id"`
	Theme    string `json:"theme"`
	Avatar   string `json:"avatar"`
	Disable  bool   `json:"disable"`
	Empty    bool   `json:"empty"`
}

func (u *User) String() string {
	return object.ToString(u)
}

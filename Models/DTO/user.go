package DTO

import "github.com/Liang-jonas/golib/object"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Group    string `json:"group"`
	GroupID  int    `json:"group-id"`
	Theme    string `json:"theme"`
	Avatar   string `json:"avatar"`
	Disable  bool   `json:"disable"`
	Empty    bool   `json:"empty"`

	ExpiresAt int64
}

func (u *User) String() string {
	return object.ToString(u)
}

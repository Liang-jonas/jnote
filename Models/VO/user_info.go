package VO

import "github.com/Liang-jonas/golib/object"

type UserInfo struct {
	Username      string   `json:"username"`
	Avatar        string   `json:"avatar"`
	Theme         string   `json:"theme"`
	BusinessPages []string `json:"businessPages"`
	UtilsPages    []string `json:"utilsPages"`
}

func (u *UserInfo) String() string {
	return object.ToString(u)
}

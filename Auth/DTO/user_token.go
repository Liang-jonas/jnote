package DTO

import "github.com/Liang-jonas/golib/object"

type UserToken struct {
	UID          int64  `json:"user-id"`
	IP           string `json:"ip"`
	UA           string `json:"user-agent"`
	AccessToken  string `json:"access-token"`
	AccessTTL    int    `json:"access-ttl"`
	RefreshToken string `json:"refresh-token"`
	RefreshTTL   int    `json:"refresh-ttl"`
}

func (u *UserToken) String() string {
	return object.ToString(u)
}

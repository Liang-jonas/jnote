package VO

import "github.com/Liang-jonas/golib/object"

type AuthToken struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	AccessTimeout  int    `json:"accessTimeout"`
	RefreshTimeout int    `json:"refreshTimeout"`
}

func (l *AuthToken) String() string {
	return object.ToString(l)
}

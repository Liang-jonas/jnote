package JWT

import (
	"errors"
	"github.com/Liang-jonas/jnote/Models/DTO"
	"strings"

	"github.com/golang-jwt/jwt"
)

func (e *engine) Parse(text string) (*DTO.User, error) {
	token, err := new(jwt.Parser).ParseWithClaims(
		text,
		&userClaims{},
		func(_ *jwt.Token) (interface{}, error) {
			return e.Signature, nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*userClaims)
	if ok &&
		token.Valid &&
		strings.Compare(claims.Subject, e.Subject) == 0 &&
		strings.Compare(claims.Id, e.JwtId) == 0 &&
		strings.Compare(claims.Issuer, e.Issuer) == 0 {
		u := new(DTO.User)
		u.Username = claims.Audience
		u.Group = claims.Group
		u.ID = claims.UID
		u.ExpiresAt = claims.ExpiresAt
		return u, err
	}
	return nil, errors.New("parse Jwt error")
}

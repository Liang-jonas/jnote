package JWT

import "github.com/golang-jwt/jwt"

type userClaims struct {
	jwt.StandardClaims
	Group string `json:"group"`
	UID   int64  `json:"user-id"`
	Magic int64  `json:"magic"`
}

package JWT

import (
	"github.com/Liang-jonas/jnote/Models/DTO"
	"time"

	"github.com/golang-jwt/jwt"
)

func (e *engine) genToken(user *DTO.User, exp int64) (string, error) {
	nowTime := time.Now().Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  user.Username,
			ExpiresAt: nowTime + exp,
			Id:        e.JwtId,
			IssuedAt:  nowTime,
			Issuer:    e.Issuer,
			NotBefore: nowTime,
			Subject:   e.Subject,
		},
		Group: user.Group,
		UID:   user.ID,
		Magic: nowTime,
	})
	return claims.SignedString(e.Signature)
}

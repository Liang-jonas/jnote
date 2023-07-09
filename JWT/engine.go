package JWT

import (
	"github.com/Liang-jonas/jnote/Conf"
	"github.com/Liang-jonas/jnote/Models/DTO"
)

var _ Engine = (*engine)(nil)

type Engine interface {
	i()
	genToken(user *DTO.User, exp int64) (string, error)
	GenAccess(user *DTO.User) (string, error)
	GenRefresh(user *DTO.User) (string, error)
	GetAccessTTL() int
	GetQ1() int64
	GetRefreshTTL() int
	Parse(text string) (*DTO.User, error)
}

type engine struct {
	Subject         string
	Issuer          string
	Signature       []byte
	JwtId           string
	LoginAccessTTL  int64
	LoginRefreshTTL int64
	Q1              int64 // 第一四分位数
}

func (*engine) i() {}

func NewJwtEngine(cfg Conf.JWT) Engine {
	jwtEngine := new(engine)
	jwtEngine.Issuer = cfg.Issuer
	jwtEngine.Signature = []byte(cfg.Signature)
	jwtEngine.JwtId = cfg.JwtId
	jwtEngine.Subject = cfg.Subject
	jwtEngine.LoginAccessTTL = cfg.LoginAccessTTL
	jwtEngine.LoginRefreshTTL = cfg.LoginRefreshTTL
	jwtEngine.Q1 = cfg.LoginRefreshTTL / 100 * 25
	return jwtEngine
}

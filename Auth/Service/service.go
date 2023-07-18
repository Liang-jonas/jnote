package Service

import (
	DTO2 "github.com/Liang-jonas/jnote/Auth/DTO"
	"github.com/Liang-jonas/jnote/DB/Mysql"
	"github.com/Liang-jonas/jnote/DB/Redis"
	"github.com/Liang-jonas/jnote/Models/Params"
)

type Service interface {
	i()
	GetUser(payload *Params.Login) (*DTO2.User, error)
	CachePolicy(user *DTO2.User) error
	CacheUserToken(token *DTO2.UserToken) error
}

type service struct {
	mDB Mysql.Repo
	rDB Redis.Repo
}

func New(mDB Mysql.Repo, rDB Redis.Repo) Service {
	return &service{mDB: mDB, rDB: rDB}
}

func (*service) i() {}

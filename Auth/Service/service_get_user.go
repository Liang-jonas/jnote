package Service

import (
	"github.com/Liang-jonas/jnote/Auth/DAO"
	"github.com/Liang-jonas/jnote/Auth/DTO"
	"github.com/Liang-jonas/jnote/Models/Params"
)

func (s *service) GetUser(payload *Params.Login) (*DTO.User, error) {
	userByMySQLBinder := DAO.NewUserByMySQLBinder()
	userByMySQLBinder.WherePassword(payload.Password)
	userByMySQLBinder.WhereUsername(payload.Username)
	userByMySQLBinder.WhereDeleteAt(nil)
	userByMySQLBinder.WhereDisable(false)
	return userByMySQLBinder.GetUser(s.mDB)
}

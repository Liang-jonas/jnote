package JWT

import (
	"github.com/Liang-jonas/jnote/Auth/DTO"
)

func (e *engine) GenRefresh(user *DTO.User) (string, error) {
	return e.genToken(user, e.LoginRefreshTTL)
}

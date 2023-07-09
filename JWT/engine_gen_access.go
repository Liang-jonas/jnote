package JWT

import "github.com/Liang-jonas/jnote/Models/DTO"

func (e *engine) GenAccess(user *DTO.User) (string, error) {
	return e.genToken(user, e.LoginAccessTTL)
}

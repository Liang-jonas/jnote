package Service

import (
	"github.com/Liang-jonas/jnote/Auth/DAO"
	"github.com/Liang-jonas/jnote/Auth/DTO"
)

func (s *service) CacheUserToken(token *DTO.UserToken) error {
	cacheUserTokenBinder := DAO.NewCacheUserTokenByRedisBinder()
	cacheUserTokenBinder.SetUserToken(token)
	if err := cacheUserTokenBinder.CacheAccess(s.rDB); err != nil {
		return err
	}

	if err := cacheUserTokenBinder.CacheRefresh(s.rDB); err != nil {
		cacheUserTokenBinder.RemoveAccess(s.rDB)
		return err
	}
	return nil
}

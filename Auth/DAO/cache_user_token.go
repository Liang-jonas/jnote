package DAO

import (
	"fmt"
	"github.com/Liang-jonas/golib/cipher"
	"github.com/Liang-jonas/jnote/Auth/DTO"
	"github.com/Liang-jonas/jnote/Constants"
	"github.com/Liang-jonas/jnote/DB/Redis"
)

type cacheUserTokenByRedisBinder struct {
	token *DTO.UserToken
	key   string
}

func NewCacheUserTokenByRedisBinder() *cacheUserTokenByRedisBinder {
	return new(cacheUserTokenByRedisBinder)
}

func (c *cacheUserTokenByRedisBinder) SetUserToken(token *DTO.UserToken) {
	c.token = token
	c.key = cipher.Md5WithoutSalt(token.IP + token.UA + fmt.Sprint(token.UID))
}

func (c *cacheUserTokenByRedisBinder) CacheAccess(db Redis.Repo) error {
	return db.SSet(Constants.UserAccessTokenFlag+c.key, c.token.AccessToken, c.token.AccessTTL)
}

func (c *cacheUserTokenByRedisBinder) CacheRefresh(db Redis.Repo) error {
	return db.SSet(Constants.UserRefreshTokenFlag+c.key, c.token.RefreshToken, c.token.RefreshTTL)
}

func (c *cacheUserTokenByRedisBinder) GetAccess(db Redis.Repo) (string, error) {
	return db.SGet(Constants.UserAccessTokenFlag + c.key)
}

func (c *cacheUserTokenByRedisBinder) GetAccessLen(db Redis.Repo) (int, error) {
	return db.SLen(Constants.UserAccessTokenFlag + c.key)
}

func (c *cacheUserTokenByRedisBinder) GetRefresh(db Redis.Repo) (string, error) {
	return db.SGet(Constants.UserRefreshTokenFlag + c.key)
}

func (c *cacheUserTokenByRedisBinder) RemoveAccess(db Redis.Repo) error {
	return db.SDel(Constants.UserAccessTokenFlag + c.key)
}

func (c *cacheUserTokenByRedisBinder) RemoveRefresh(db Redis.Repo) error {
	return db.SDel(Constants.UserRefreshTokenFlag + c.key)
}

package DAO

import (
	"fmt"
	"github.com/Liang-jonas/jnote/Auth/DTO"
	"github.com/Liang-jonas/jnote/Constants"
	"github.com/Liang-jonas/jnote/DB/Redis"
)

type cacheUserPolicesByRedisBinder struct {
	uid       string
	apis      []string
	resources []int64
}

func NewCacheUserPolicesByRedisBinder() *cacheUserPolicesByRedisBinder {
	return new(cacheUserPolicesByRedisBinder)
}

func (c *cacheUserPolicesByRedisBinder) SetUserID(user *DTO.User) {
	c.uid = fmt.Sprint(user.ID)
}

func (c *cacheUserPolicesByRedisBinder) SetApiPolices(apis []DTO.API) {
	for _, api := range apis {
		c.apis = append(c.apis, *api.Path)
	}
}

func (c *cacheUserPolicesByRedisBinder) SetNotePolices(notes []DTO.MapPolicyNote) {
	for _, note := range notes {
		c.resources = append(c.resources, note.NoteID)
	}
}

func (c *cacheUserPolicesByRedisBinder) SetImagePolices(images []DTO.MapPolicyImage) {
	for _, image := range images {
		c.resources = append(c.resources, image.ImageID)
	}
}

func (c *cacheUserPolicesByRedisBinder) CacheUserPolices(db Redis.Repo) error {
	if len(c.apis) != 0 {
		if err := db.SetSet(Constants.UserApiFlag+c.uid, c.apis); err != nil {
			return err
		}
	} else {
		if err := db.SetSet(Constants.UserApiFlag+c.uid, nil); err != nil {
			return err
		}
	}

	if len(c.resources) != 0 {
		if err := db.SetSet(Constants.UserResourceFlag+c.uid, c.resources); err != nil {
			db.Del(Constants.UserApiFlag + c.uid)
			return err
		}
	} else {
		if err := db.SetSet(Constants.UserResourceFlag+c.uid, nil); err != nil {
			db.Del(Constants.UserApiFlag + c.uid)
			return err
		}
	}

	return nil
}

func (c *cacheUserPolicesByRedisBinder) DelUserPolices(db Redis.Repo) error {
	if err := db.Del(Constants.UserApiFlag + c.uid); err != nil {
		return err
	}

	if err := db.Del(Constants.UserResourceFlag + c.uid); err != nil {
		return err
	}
	return nil
}

func (c *cacheUserPolicesByRedisBinder) VerifyApi(api string, db Redis.Repo) (bool, error) {
	return db.SetExist(Constants.UserApiFlag+c.uid, api)
}

package Service

import (
	"github.com/Liang-jonas/jnote/Auth/DAO"
	"github.com/Liang-jonas/jnote/Auth/DTO"
)

func (s *service) CachePolicy(user *DTO.User) error {
	getPolicyBinder := DAO.NewPolicyByMysqlBinder()
	getPolicyBinder.WhereUserId(user.ID)
	getPolicyBinder.WhereUserGroupId(user.GroupID)
	policy, err := getPolicyBinder.GetPolicy(s.mDB)
	if err != nil {
		return err
	}
	getImagePolicyBinder := DAO.NewPolicyByMysqlBinder()
	getImagePolicyBinder.WhereMapPolicyImageWithPid(policy.PID)
	imagePolices, err := getImagePolicyBinder.GetImagePolices(s.mDB)
	if err != nil {
		return err
	}

	// 读取笔记策略
	getNotePolicyBinder := DAO.NewPolicyByMysqlBinder()
	getNotePolicyBinder.WhereMapPolicyNoteWithPid(policy.PID)
	notePolices, err := getNotePolicyBinder.GetNotePolices(s.mDB)
	if err != nil {
		return err
	}

	// 读取api策略
	getApiPolicyBinder := DAO.NewPolicyByMysqlBinder()
	getApiPolicyBinder.WhereMapPolicyAndApi()
	getApiPolicyBinder.WhereMapPolicyApiWithPid(policy.PID)
	apiPolices, err := getApiPolicyBinder.GetApiPolices(s.mDB)
	if err != nil {
		return err
	}

	cacheUserPolicesBinder := DAO.NewCacheUserPolicesByRedisBinder()
	cacheUserPolicesBinder.SetUserID(user)
	cacheUserPolicesBinder.SetApiPolices(apiPolices)
	cacheUserPolicesBinder.SetImagePolices(imagePolices)
	cacheUserPolicesBinder.SetNotePolices(notePolices)
	return cacheUserPolicesBinder.CacheUserPolices(s.rDB)
}

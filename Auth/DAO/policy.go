package DAO

import (
	"errors"
	"fmt"
	"github.com/Liang-jonas/golib/db"
	"github.com/Liang-jonas/jnote/Auth/DTO"
	"github.com/Liang-jonas/jnote/DB/Mysql"
	"strings"
)

const (
	getPolicySql      string = "SELECT `policy`.`id`, `policy`.`user-id`, `policy`.`group-id` FROM `policy`"
	getApiPolicySql   string = "SELECT `api`.`id`, `api`.`name`, `api`.`path`, `api`.`description` FROM `api`, `map-policy-api`"
	getImagePolicySql string = "SELECT `map-policy-image`.`id`, `map-policy-image`.`image-id`, `map-policy-image`.`policy-id` FROM `map-policy-image`"
	getNotePolicySql  string = "SELECT `map-policy-note`.`id`, `map-policy-note`.`note-id`, `map-policy-note`.`policy-id` FROM `map-policy-note`"
)

type policyByMysqlBinder struct {
	db.Object
}

func NewPolicyByMysqlBinder() *policyByMysqlBinder {
	p := new(policyByMysqlBinder)
	p.Object = db.NewDbObject()
	return p
}

func (p *policyByMysqlBinder) WhereUserId(uid int64) {
	p.SetAndParams("`policy`.`user-id` = ?")
	p.SetAndValues(uid)
}

func (p *policyByMysqlBinder) WhereUserGroupId(uid int64) {
	p.SetAndParams("`policy`.`group-id` = ?")
	p.SetAndValues(uid)
}

func (p *policyByMysqlBinder) WhereMapPolicyApiWithPid(pid int64) {
	p.whereDbNameAndPolicyId("map-policy-api", pid)
}

func (p *policyByMysqlBinder) WhereMapPolicyImageWithPid(pid int64) {
	p.whereDbNameAndPolicyId("map-policy-image", pid)
}

func (p *policyByMysqlBinder) WhereMapPolicyNoteWithPid(pid int64) {
	p.whereDbNameAndPolicyId("map-policy-note", pid)
}

func (p *policyByMysqlBinder) WhereMapPolicyAndApi() {
	p.SetAndParams("`map-policy-api`.`api-id` = `api`.`id`")
}

func (p *policyByMysqlBinder) whereDbNameAndPolicyId(dbName string, pid int64) {
	p.SetAndParams(fmt.Sprintf("`%s`.`group-id` = ?", dbName))
	p.SetAndValues(pid)
}

func (p *policyByMysqlBinder) genSqlBuf(baseSql string) *strings.Builder {
	buf := new(strings.Builder)
	buf.WriteString(baseSql)
	buf.WriteString(" WHERE ")
	p.GetAndParamsToBuf(buf, " AND ")
	return buf
}

func (p *policyByMysqlBinder) GetPolicy(db Mysql.Repo) (*DTO.Policy, error) {
	buf := p.genSqlBuf(getPolicySql)

	rows, err := db.GetDB().QueryContext(db.GetCtx(), buf.String(), p.GetAndValues()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("the policy is empty")
	}

	policy := new(DTO.Policy)

	if err := rows.Scan(
		&policy.PID,
		&policy.UID,
		&policy.GID,
	); err != nil {
		return nil, err
	}

	return policy, nil
}

func (p *policyByMysqlBinder) GetApiPolices(db Mysql.Repo) ([]DTO.API, error) {
	buf := p.genSqlBuf(getApiPolicySql)

	rows, err := db.GetDB().QueryContext(db.GetCtx(), buf.String(), p.GetAndValues()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apis []DTO.API

	for rows.Next() {
		api := DTO.API{}
		if err := rows.Scan(
			&api.ID,
			&api.Name,
			&api.Path,
			&api.Description,
		); err != nil {
			return nil, err
		}
		apis = append(apis, api)
	}

	return apis, nil
}

func (p *policyByMysqlBinder) GetImagePolices(db Mysql.Repo) ([]DTO.MapPolicyImage, error) {
	buf := p.genSqlBuf(getImagePolicySql)

	rows, err := db.GetDB().QueryContext(db.GetCtx(), buf.String(), p.GetAndValues()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []DTO.MapPolicyImage

	for rows.Next() {
		image := DTO.MapPolicyImage{}
		if err := rows.Scan(
			&image.ID,
			&image.ImageID,
			&image.PolicyID,
		); err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	return images, nil
}

func (p *policyByMysqlBinder) GetNotePolices(db Mysql.Repo) ([]DTO.MapPolicyNote, error) {
	buf := p.genSqlBuf(getNotePolicySql)
	rows, err := db.GetDB().QueryContext(db.GetCtx(), buf.String(), p.GetAndValues()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []DTO.MapPolicyNote

	for rows.Next() {
		note := DTO.MapPolicyNote{}
		if err := rows.Scan(
			&note.ID,
			&note.NoteID,
			&note.PolicyID,
		); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

package DAO

import (
	"github.com/Liang-jonas/golib/db"
	"github.com/Liang-jonas/jnote/Auth/DTO"
	"github.com/Liang-jonas/jnote/DB/Mysql"
	"strings"
	"time"
)

type userByMySQLBinder struct {
	db.Object
}

const (
	queryUserSql string = "SELECT `user`.`id`,`user`.`username`,`group`.`name`, `user`.`group-id`, `theme`.`name`,`user`.`avatar`,`user`.`disable` FROM `user`, `theme`, `group` where `group`.`id` = `user`.`group-id` AND `theme`.`id` = `user`.`theme-id`"
)

func NewUserByMySQLBinder() *userByMySQLBinder {
	u := new(userByMySQLBinder)
	u.Object = db.NewDbObject()
	return u
}

func (u *userByMySQLBinder) GetUser(db Mysql.Repo) (*DTO.User, error) {
	var sqlBuf strings.Builder
	sqlBuf.WriteString(queryUserSql)
	u.GetAndParamsToBuf(&sqlBuf, " and ")

	rows, err := db.GetDB().QueryContext(db.GetCtx(), sqlBuf.String(), u.GetAndValues()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := new(DTO.User)

	if !rows.Next() {
		user.Empty = true
		return user, nil
	}

	if err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Group,
		&user.GroupID,
		&user.Theme,
		&user.Avatar,
		&user.Disable,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userByMySQLBinder) WhereUsername(username string) {
	if username == "" {
		return
	}
	u.SetAndParams("`user`.`username` = ?")
	u.SetAndValues(username)
}

func (u *userByMySQLBinder) WherePassword(password string) {
	if password == "" {
		return
	}
	u.SetAndParams("`user``password` = ?")
	u.SetAndValues(password)
}

func (u *userByMySQLBinder) WhereDeleteAt(t *time.Time) {
	u.SetAndParams("`user`.`delete-at` <=> ?")
	u.SetAndValues(t)
}

func (u *userByMySQLBinder) WhereDisable(disable bool) {
	u.SetAndParams("`user`.`disable` = ?")
	u.SetAndValues(disable)
}

func (u *userByMySQLBinder) WhereUserId(uid int64) {
	if uid == 0 {

		return
	}
	u.SetAndValues(uid)
	u.SetAndParams(" `id` = ? ")
}

func (u *userByMySQLBinder) Where() {}

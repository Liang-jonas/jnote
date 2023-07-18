package Mysql

import (
	"fmt"
	"github.com/Liang-jonas/jnote/Conf"
	"github.com/Liang-jonas/jnote/Models/Entity"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	cfg := Conf.Mysql{
		Ip:             "192.168.78.128",
		Port:           "3306",
		Username:       "root",
		Password:       "abcd1234",
		Dbname:         "JNA",
		Options:        "charset=utf8mb4&parseTime=True&loc=Local",
		PoolSize:       20,
		ConnectTimeout: 60,
	}
	db, err := NewDB(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	sql := "UPDATE JNA.`user` SET `update-at`=? WHERE id=123;"
	_, err = db.GetDB().ExecContext(db.GetCtx(), sql, time.Now())
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMysql(t *testing.T) {
	cfg := Conf.Mysql{
		Ip:             "192.168.78.128",
		Port:           "3306",
		Username:       "root",
		Password:       "abcd1234",
		Dbname:         "JNA",
		Options:        "charset=utf8mb4&parseTime=True&loc=Local",
		PoolSize:       20,
		ConnectTimeout: 60,
	}
	db, err := NewDB(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	sql := "SELECT id, username, password, `theme-id`, `group-id`, disable, avatar, `create-at`, `update-at`, `delete-at` FROM JNA.`user` WHERE id=123;"
	rows, err := db.GetDB().QueryContext(db.GetCtx(), sql)
	if err != nil {
		t.Error(err)
		return
	}
	if !rows.Next() {
		return
	}
	u := new(Entity.User)
	if err := rows.Scan(
		&u.ID,
		&u.Username,
		&u.Password,
		&u.ThemeID,
		&u.GroupID,
		&u.Disable,
		&u.Avatar,
		&u.CreateAt,
		&u.UpdateAt,
		&u.DeleteAt,
	); err != nil {
		t.Error(err)
		return
	}
	fmt.Println(u)
}

func TestI(t *testing.T) {
	var x int64 = 92883
	fmt.Println(fmt.Sprint(x))
}

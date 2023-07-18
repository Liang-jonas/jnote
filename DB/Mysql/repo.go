package Mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Liang-jonas/jnote/Conf"
	_ "github.com/go-sql-driver/mysql"
)

const dsnFormat = "%s:%s@tcp(%s:%s)/%s?%s"

var _ Repo = (*dbRepo)(nil)

// Repo Mysql数据库存储对象
type Repo interface {
	i()

	// Close
	// @Title	Close
	// @Description 关闭数据库连接
	// @Tags DB.Mysql
	// @Return error "关闭时可能出现的错误"
	Close() error

	// GetDB
	// @Title	GetDB
	// @Description 获取数据库对象
	// @Tags DB.Mysql
	// @Return sql.DB "数据库对象实例"
	GetDB() *sql.DB

	// GetCtx
	// @Title	GetCtx
	// @Description 获取上下文对象
	// @Tags DB.Mysql
	// @Return context.Context
	GetCtx() context.Context
}

type dbRepo struct {
	db     *sql.DB
	ctx    context.Context
	cancel context.CancelFunc
}

func (*dbRepo) i() {}

func NewDB(cfg Conf.Mysql) (Repo, error) {
	repo := new(dbRepo)
	repo.ctx, repo.cancel = context.WithCancel(context.Background())
	db, err := sql.Open("mysql", fmt.Sprintf(dsnFormat, cfg.Username, cfg.Password, cfg.Ip, cfg.Port, cfg.Dbname, cfg.Options))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.PoolSize)
	db.SetMaxOpenConns(cfg.PoolSize)
	repo.db = db
	return repo, nil
}

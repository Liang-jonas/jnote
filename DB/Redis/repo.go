package Redis

import (
	"context"
	"fmt"
	"github.com/Liang-jonas/jnote/Conf"
	"time"

	"github.com/go-redis/redis"
)

const addrFormat = "%s:%s"

var _ Repo = (*dbRepo)(nil)

// Repo Redis数据库存储对象
type Repo interface {
	i()

	// Close
	// @Title Close
	// @Description 关闭数据库连接
	// @Tags DB.Redis
	// @Return error "关闭时可能出现的错误"
	Close() error

	// Del
	// @Title Del
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	Del(key string) error

	// HDel
	// @Title HDel
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	HDel(key, field string) error

	// HExist
	// @Title HExist
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	HExist(key string, field string) (bool, error)

	// HGet
	// @Title HGet
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	HGet(key, field string) (string, error)

	// HSet
	// @Title HSet
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	HSet(key, field string, value interface{}) error

	// SetDel
	// @Title SetDel
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	SetDel(key, field string) error

	// SetExist
	// @Title SetExist
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	SetExist(key, value string) (bool, error)

	// SetGetAll
	// @Title SetGetAll
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	SetGetAll(key string) ([]string, error)

	// SetSet
	// @Title SetSet
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	SetSet(key string, value ...interface{}) error

	// SDel
	// @Title SDel
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	SDel(key string) error

	// SGet
	// @Title SGet
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	SGet(key string) (string, error)

	// SSet
	// @Title SSet
	// @Description
	// @Tags DB.Redis
	// @Return error "时可能出现的错误"
	SSet(key, value string, ttl int) error

	SLen(key string) (int, error)
}

type dbRepo struct {
	dbKeyPrefix string
	db          *redis.Client
	ctx         context.Context
	cancel      context.CancelFunc
}

func (*dbRepo) i() {}

func NewDB(cfg Conf.Redis) (Repo, error) {
	repo := new(dbRepo)
	repo.ctx, repo.cancel = context.WithCancel(context.Background())
	repo.db = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf(addrFormat, cfg.Ip, cfg.Port),
		Password:    cfg.Password,
		DB:          cfg.Dbname,
		PoolSize:    cfg.PoolSize,
		DialTimeout: time.Duration(cfg.ConnectTimeout) * time.Second,
	})
	repo.dbKeyPrefix = cfg.KeyPrefix
	if err := repo.db.Ping().Err(); err != nil {
		return nil, err
	}
	return repo, nil
}

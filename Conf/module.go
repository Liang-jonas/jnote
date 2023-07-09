package Conf

import "github.com/Liang-jonas/golib/object"

type Base struct {
	Ip        string
	LogPath   string
	Mode      string
	Port      string
	UriPrefix string
	LogColor  bool
}

type Mysql struct {
	Ip             string
	Port           string
	Password       string
	Username       string
	Dbname         string
	Options        string
	PoolSize       int
	ConnectTimeout int
}

type Redis struct {
	Ip             string
	Port           string
	Password       string
	KeyPrefix      string
	Dbname         int
	PoolSize       int
	ConnectTimeout int
}

type JWT struct {
	Signature       string
	Issuer          string
	JwtId           string
	Subject         string
	LoginAccessTTL  int64
	LoginRefreshTTL int64
}

type Config struct {
	BaseCfg  Base
	MysqlCfg Mysql
	RedisCfg Redis
	JwtCfg   JWT
}

func (b *Base) String() string {
	return object.ToString(b)
}

func (m *Mysql) String() string {
	return object.ToString(m)
}

func (r *Redis) String() string {
	return object.ToString(r)
}

func (j *JWT) String() string {
	return object.ToString(j)
}

func (c *Config) String() string {
	return object.ToString(c)
}

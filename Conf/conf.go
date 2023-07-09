package Conf

import "github.com/go-ini/ini"

var _ Conf = (*conf)(nil)

// Conf 配置文件读取对象
type Conf interface {
	i()

	//GetCfg
	// @Title GetCfg
	// @Description 获取config对象
	// @Tags Conf
	// @Return *Config "config对象"
	GetCfg() *Config

	// Read
	// @Title Read
	// @Description 读取指定路径下的配置文件
	// @Tags Conf
	// @Param string "配置文件路径"
	// @Return error "读取时可能会出现的错误"
	Read(cfgPath string) error
}

type conf struct {
	ini *ini.File
	cfg *Config
}

func (*conf) i() {}

func NewConf() Conf {
	c := new(conf)
	c.cfg = new(Config)
	return c
}

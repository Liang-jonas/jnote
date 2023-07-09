package Conf

import (
	"errors"
	"github.com/Liang-jonas/golib/file"
	"github.com/go-ini/ini"
)

func (c *conf) Read(cfgPath string) error {
	if !file.FileExist(cfgPath) {
		return errors.New("The file does not exist: " + cfgPath)
	}
	var err error = nil
	c.ini, err = ini.Load(cfgPath)
	if err != nil {
		return err
	}
	c.cfg.BaseCfg.Ip = c.ini.Section("base").Key("ip").String()
	c.cfg.BaseCfg.LogPath = c.ini.Section("base").Key("log-path").String()
	c.cfg.BaseCfg.Mode = c.ini.Section("base").Key("mode").String()
	c.cfg.BaseCfg.Port = c.ini.Section("base").Key("port").String()
	c.cfg.BaseCfg.UriPrefix = c.ini.Section("base").Key("uri-prefix").String()
	if c.cfg.BaseCfg.LogColor, err = c.ini.Section("base").Key("log-color").Bool(); err != nil {
		return err
	}

	c.cfg.MysqlCfg.Ip = c.ini.Section("mysql").Key("ip").String()
	c.cfg.MysqlCfg.Port = c.ini.Section("mysql").Key("port").String()
	c.cfg.MysqlCfg.Username = c.ini.Section("mysql").Key("username").String()
	c.cfg.MysqlCfg.Password = c.ini.Section("mysql").Key("password").String()
	c.cfg.MysqlCfg.Dbname = c.ini.Section("mysql").Key("dbname").String()
	c.cfg.MysqlCfg.Options = c.ini.Section("mysql").Key("options").String()
	if c.cfg.MysqlCfg.ConnectTimeout, err = c.ini.Section("mysql").Key("connect-timeout").Int(); err != nil {
		return err
	}
	if c.cfg.MysqlCfg.PoolSize, err = c.ini.Section("mysql").Key("pool-size").Int(); err != nil {
		return err
	}

	c.cfg.RedisCfg.Ip = c.ini.Section("redis").Key("ip").String()
	c.cfg.RedisCfg.Port = c.ini.Section("redis").Key("port").String()
	c.cfg.RedisCfg.Password = c.ini.Section("redis").Key("password").String()
	c.cfg.RedisCfg.KeyPrefix = c.ini.Section("redis").Key("key-prefix").String()
	if c.cfg.RedisCfg.Dbname, err = c.ini.Section("redis").Key("dbname").Int(); err != nil {
		return err
	}
	if c.cfg.RedisCfg.ConnectTimeout, err = c.ini.Section("redis").Key("connect-timeout").Int(); err != nil {
		return err
	}
	if c.cfg.RedisCfg.PoolSize, err = c.ini.Section("redis").Key("pool-size").Int(); err != nil {
		return err
	}

	c.cfg.JwtCfg.Signature = c.ini.Section("jwt").Key("signature").String()
	c.cfg.JwtCfg.Issuer = c.ini.Section("jwt").Key("issuer").String()
	c.cfg.JwtCfg.JwtId = c.ini.Section("jwt").Key("jwt-id").String()
	c.cfg.JwtCfg.Subject = c.ini.Section("jwt").Key("subject").String()
	if c.cfg.JwtCfg.LoginAccessTTL, err = c.ini.Section("jwt").Key("login-access-ttl").Int64(); err != nil {
		return err
	}
	if c.cfg.JwtCfg.LoginRefreshTTL, err = c.ini.Section("jwt").Key("login-refresh-ttl").Int64(); err != nil {
		return err
	}

	return nil
}

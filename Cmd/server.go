package Cmd

import (
	"fmt"
	"github.com/Liang-jonas/golib"
	"github.com/Liang-jonas/golib/logger"
	"github.com/Liang-jonas/jnote/Conf"
	"github.com/Liang-jonas/jnote/DB/Mysql"
	"github.com/Liang-jonas/jnote/DB/Redis"
	"github.com/Liang-jonas/jnote/JWT"
	"github.com/Liang-jonas/jnote/Middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var defaultCfgPath string

func RunServer(cfgPath string) {
	if cfgPath == "" {
		cfgPath = defaultCfgPath
	}

	cfgEngine := Conf.NewConf()
	if err := cfgEngine.Read(cfgPath); err != nil {
		fmt.Println("cfg error: ", err)
		return
	}

	cfg := cfgEngine.GetCfg()
	logger, err := logger.NewLogger(cfg.BaseCfg.LogColor, cfg.BaseCfg.LogPath)
	if err != nil {
		fmt.Println("logger error: ", err)
		return
	}

	mDB, err := Mysql.NewDB(cfg.MysqlCfg)
	if err != nil {
		logger.Errorf("mysqlDB error: %s", err)
		return
	}
	defer mDB.Close()
	logger.Infof("mysql database connection to %s:%s, dbName: %s, options: %s", cfg.MysqlCfg.Ip,
		cfg.MysqlCfg.Port, cfg.MysqlCfg.Dbname, cfg.MysqlCfg.Options)

	rDB, err := Redis.NewDB(cfg.RedisCfg)
	if err != nil {
		logger.Errorf("redisDB error: %s", err)
		return
	}
	defer rDB.Close()
	logger.Infof("Redis database connection to %s:%s", cfg.RedisCfg.Ip, cfg.RedisCfg.Port)

	jwtEngine := JWT.NewJwtEngine(cfg.JwtCfg)
	middlewares := Middleware.NewMiddleware(logger, jwtEngine, mDB, rDB)

	ginEngine := gin.New()

	ginEngine.NoRoute(func(ctx *gin.Context) {
		golib.FailWithMsg(golib.NoUriCode, golib.NoUriMsg, ctx)
	})

	ginEngine.MaxMultipartMemory = 8 << 20

	ginEngine.Use(
		//middlewares.HandleError,
		middlewares.Logger,
		middlewares.Cors,
		gzip.Gzip(gzip.DefaultCompression),
	)

	if cfg.BaseCfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := ginEngine.Run(fmt.Sprintf("%s:%s", cfg.BaseCfg.Ip, cfg.BaseCfg.Port)); err != nil {
		logger.Error(err)
	}
}

func ServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Short:   "Run the jnote server.",
		Long:    "Run the video-download service and start it by default or by specifying a configuration file.",
		GroupID: rootGroupID,
		Run: func(cmd *cobra.Command, args []string) {
			var cfgPath string
			cmd.Flags().StringVarP(&cfgPath, "file", "f", "", "set config path")
			cmd.Flags().Parse(args)
			//RunServer(cfgPath)
		},
		DisableSuggestions: true,
		DisableFlagParsing: true,
	}

	cmd.CompletionOptions.HiddenDefaultCmd = true

	return cmd
}

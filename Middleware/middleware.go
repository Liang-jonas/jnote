package Middleware

import (
	"github.com/Liang-jonas/jnote/DB/Mysql"
	"github.com/Liang-jonas/jnote/DB/Redis"
	"github.com/Liang-jonas/jnote/JWT"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var _ Middleware = (*middleware)(nil)

type Middleware interface {
	i()
	AuthImage(ctx *gin.Context)
	AuthUser(ctx *gin.Context)
	HandleError(ctx *gin.Context)
	Cors(c *gin.Context)
	Logger(ctx *gin.Context)
}

type middleware struct {
	logger    *logrus.Logger
	jwtEngine JWT.Engine
	mDB       Mysql.Repo
	rDB       Redis.Repo
	//authorizedService Authorized.Service
}

func (*middleware) i() {}

func NewMiddleware(logger *logrus.Logger, jwtEngine JWT.Engine, mDB Mysql.Repo, rDB Redis.Repo) Middleware {
	m := new(middleware)
	m.logger = logger
	m.mDB = mDB
	m.rDB = rDB
	m.jwtEngine = jwtEngine
	//m.authorizedService = Authorized.NewService(mDB, rDB)
	return m
}

package Auth

import (
	"github.com/Liang-jonas/jnote/Auth/Service"
	"github.com/Liang-jonas/jnote/Constants"
	"github.com/Liang-jonas/jnote/DB/Mysql"
	"github.com/Liang-jonas/jnote/DB/Redis"
	"github.com/Liang-jonas/jnote/JWT"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Controller interface {
	i()
	loginApi(ctx *gin.Context)
	logoutApi(ctx *gin.Context)
	refreshApi(ctx *gin.Context)
	registerApi(ctx *gin.Context)
}

type controller struct {
	logger    *logrus.Logger
	svc       Service.Service
	jwtEngine JWT.Engine
}

func New(
	uriPrefix string,
	ginEngine *gin.Engine,
	logger *logrus.Logger,
	mDB Mysql.Repo,
	rDB Redis.Repo,
	jwtEngine JWT.Engine) Controller {
	c := &controller{
		logger:    logger,
		jwtEngine: jwtEngine,
		svc:       Service.New(mDB, rDB),
	}
	ginEngine.POST(uriPrefix+Constants.LoginApi, c.loginApi)
	ginEngine.POST(uriPrefix+Constants.LogoutApi, c.logoutApi)
	ginEngine.POST(uriPrefix+Constants.RefreshApi, c.refreshApi)
	ginEngine.POST(uriPrefix+Constants.RegisterApi, c.registerApi)
	return c
}

func (*controller) i() {}

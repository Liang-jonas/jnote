package Auth

import (
	"github.com/Liang-jonas/golib"
	"github.com/Liang-jonas/jnote/Auth/DTO"
	"github.com/Liang-jonas/jnote/Constants"
	"github.com/Liang-jonas/jnote/Models/Params"
	"github.com/Liang-jonas/jnote/Models/VO"
	"github.com/gin-gonic/gin"
)

func (c *controller) loginApi(ctx *gin.Context) {
	payload := new(Params.Login)
	if err := ctx.ShouldBindJSON(payload); err != nil {
		golib.FailWithMsg(Constants.ParamParseErrorCode, Constants.ParamParseErrorMsg, ctx)
		return
	}

	if payload.Username == "" {
		golib.FailWithMsg(Constants.ParamElementEmptyCode, Constants.ParamUsernameEmptyMsg, ctx)
		return
	}
	if payload.Password == "" {
		golib.FailWithMsg(Constants.ParamElementEmptyCode, Constants.ParamPasswordEmptyMsg, ctx)
		return
	}

	user, err := c.svc.GetUser(payload)
	if err != nil {
		c.logger.Error(err)
		golib.FailWithInternal(ctx)
		return
	}

	if user.Empty {
		golib.FailWithMsg(Constants.UserNotFoundCode, Constants.UserNotFoundMsg, ctx)
		return
	}

	if user.Disable {
		golib.FailWithMsg(Constants.UserDisableCode, Constants.UserDisableMsg, ctx)
		return
	}

	if err := c.svc.CachePolicy(user); err != nil {
		c.logger.Error(err)
		golib.FailWithInternal(ctx)
		return
	}

	accessToken, err := c.jwtEngine.GenAccess(user)
	if err != nil {
		c.logger.Error(err)
		golib.FailWithInternal(ctx)
		return
	}

	refreshToken, err := c.jwtEngine.GenRefresh(user)
	if err != nil {
		c.logger.Error(err)
		golib.FailWithInternal(ctx)
		return
	}

	userTokenDTO := new(DTO.UserToken)
	userTokenDTO.AccessToken = accessToken
	userTokenDTO.AccessTTL = c.jwtEngine.GetAccessTTL()
	userTokenDTO.RefreshToken = refreshToken
	userTokenDTO.RefreshTTL = c.jwtEngine.GetRefreshTTL()
	userTokenDTO.UID = user.ID
	userTokenDTO.IP = ctx.ClientIP()
	userTokenDTO.UA = ctx.GetHeader("User-Agent")

	// TODO: 考虑是否做一个多端登录的认证的过程

	if err := c.svc.CacheUserToken(userTokenDTO); err != nil {
		c.logger.Error(err)
		golib.FailWithInternal(ctx)
		return
	}

	authToken := new(VO.AuthToken)
	authToken.AccessToken = userTokenDTO.AccessToken
	authToken.RefreshToken = userTokenDTO.RefreshToken
	authToken.AccessTimeout = userTokenDTO.AccessTTL
	authToken.RefreshTimeout = userTokenDTO.RefreshTTL

	golib.OkWithData(authToken, ctx)
}

package Middleware

import (
	"fmt"
	"github.com/Liang-jonas/golib"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func (m *middleware) HandleError(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			m.logger.Error(fmt.Sprintf("%+v", err), "\nStack: ", string(debug.Stack()))
			golib.FailWithMsg(golib.InternalErrCode, golib.InternalErrMsg, ctx)
		}
	}()
	ctx.Next()
}

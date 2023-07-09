package Middleware

import (
	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthUser(ctx *gin.Context) {

	ctx.Next()
}

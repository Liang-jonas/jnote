package Auth

import (
	"github.com/Liang-jonas/golib"
	"github.com/gin-gonic/gin"
)

func (c *controller) logoutApi(ctx *gin.Context) {

	golib.FailWithNotOpenApi(ctx)
}

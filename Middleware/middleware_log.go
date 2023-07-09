package Middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var timeMap map[int]string = map[int]string{
	0: "ns",
	1: "μs",
	2: "ms",
	3: "s",
}

func timeMapToString(latencyTime int64, i int) string {
	t := latencyTime / 1e3
	if t == 0 {
		return fmt.Sprintf("%d %s", latencyTime, timeMap[i])
	} else {
		return timeMapToString(t, i+1)
	}
}

func (m *middleware) Logger(ctx *gin.Context) {
	startTime := time.Now().UnixNano()
	ctx.Next()
	endTime := time.Now().UnixNano()
	latencyTime := endTime - startTime
	reqMethod := ctx.Request.Method
	reqUri := ctx.Request.RequestURI
	statusCode := ctx.Writer.Status()
	clientIP := ctx.ClientIP()
	m.logger.Infof("| %3d | %15s |  %13v  | %s | %s ",
		statusCode,
		timeMapToString(latencyTime, 0),
		clientIP,
		reqMethod,
		reqUri,
	)
}

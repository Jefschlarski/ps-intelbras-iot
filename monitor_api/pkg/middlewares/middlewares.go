package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware
func Logger(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		next(ctx)
		fmt.Printf("%s %s %s %s %s %d %s\n", start.Format("02/Jan/2006:15:04:05 -0700"), ctx.Request.Method, ctx.Request.URL.Path, ctx.Request.Proto, ctx.ClientIP(), ctx.Writer.Status(), time.Since(start))
	}
}

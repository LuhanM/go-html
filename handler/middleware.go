package handler

import (
	"fmt"

	"github.com/LuhanM/go-htmx/schemas"
	"github.com/gin-gonic/gin"
)

func StatsMiddleware(v *schemas.Viewers) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		fmt.Println(ctx.Request.RemoteAddr)
		v.Increase(ctx.Request.RemoteAddr)
	}
	return gin.HandlerFunc(fn)
}

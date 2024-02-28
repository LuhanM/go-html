package router

import (
	"github.com/LuhanM/go-htmx/handler"
	"github.com/LuhanM/go-htmx/schemas"
	"github.com/gin-gonic/gin"
)

func Initialize(visitors *schemas.Viewers) {
	r := gin.Default()

	r.Use(handler.StatsMiddleware(visitors))

	r.GET("/", handler.HomePageHandler)
	r.GET("/stats", handler.VisitorsCounterHandler(visitors))

	r.Run(":8080")
}

package router

import (
	"github.com/LuhanM/go-htmx/handler"
	"github.com/LuhanM/go-htmx/schemas"
	"github.com/gin-gonic/gin"
)

func Initialize(visitors *schemas.Viewers) {
	r := gin.Default()

	r.Use(handler.StatsMiddleware(visitors))
	//r.Use(gin.Logger()) it's not necessary because we using gin.Default()
	r.Static("/static", "./static")

	r.GET("/", handler.HomePageHandler)
	r.GET("/stats", handler.VisitorsCounterHandler(visitors))

	r.Run(":8080")
}

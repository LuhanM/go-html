package handler

import (
	"html/template"
	"io"

	"github.com/LuhanM/go-htmx/schemas"
	"github.com/gin-gonic/gin"
)

var pageAddress = "stats.html"

func VisitorsCounterHandler(v *schemas.Viewers) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		parseTemplatePage(v, ctx.Writer)
	}
	return gin.HandlerFunc(fn)
}

func parseTemplatePage(v *schemas.Viewers, wr io.Writer) {
	tmpl, _ := template.ParseFiles(pageAddress)
	data := map[string]int{
		"ViewerCounterValue": v.GetValue(),
	}
	tmpl.ExecuteTemplate(wr, pageAddress, data)
}

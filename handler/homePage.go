package handler

import (
	"html/template"
	"io"

	"github.com/gin-gonic/gin"
)

func HomePageHandler(ctx *gin.Context) {
	parseHomePage(ctx.Writer)
}

func parseHomePage(wr io.Writer) {
	tmpl, _ := template.ParseFiles("index.html")
	data := map[string]int{}
	tmpl.ExecuteTemplate(wr, "index.html", data)
}

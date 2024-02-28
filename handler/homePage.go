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
	tmpl, _ := template.New("").ParseFiles("templates/home.html")
	tmpl.ExecuteTemplate(wr, "Home", nil)
}

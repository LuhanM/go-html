package main

import (
	"time"

	"github.com/LuhanM/go-htmx/router"
	"github.com/LuhanM/go-htmx/schemas"
)

// var visitors = &schemas.Viewers{
// 	List: make(map[string]time.Time),
// }

func main() {
	visitors := &schemas.Viewers{
		List: make(map[string]time.Time),
	}

	router.Initialize(visitors)
}

// func homeHandler(ctx *gin.Context) {
// 	tmpl, _ := template.ParseFiles("index.html")
// 	data := map[string]int{
// 		"ViewerCounterValue": viewers.GetValue(),
// 	}
// 	tmpl.ExecuteTemplate(ctx.Writer, "index.html", data)
// }

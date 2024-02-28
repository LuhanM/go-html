package main

import (
	"time"

	"github.com/LuhanM/go-htmx/router"
	"github.com/LuhanM/go-htmx/schemas"
)

func main() {
	visitors := &schemas.Viewers{
		List: make(map[string]time.Time),
	}

	router.Initialize(visitors)
}

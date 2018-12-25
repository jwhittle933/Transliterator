package main

import (
	"transliterator/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	//e.File registers new route with static file to serve
	e.File("/", "views/index.html")
	e.File("/about", "views/about.html")
	e.POST("/", handler.HomeHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

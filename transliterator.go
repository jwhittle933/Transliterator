package main

import (
	"io"
	"net/http"
	"text/template"
	"transliterator/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Template for view render
type Template struct {
	templates *template.Template
}

//Render method
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Static("/assets"))

	//Static route serves from assets directory
	e.Static("/", "assets/transliterator/dist")

	t := &Template{
		//uncomment below to move template rendering to /views
		// templates: template.Must(template.ParseGlob("views/*.html")),

		//uncoment below to use Vue in Echo Template, add {{ define "index"}} && {{ end }} to index.html in /dist every time the Vue project is built
		templates: template.Must(template.ParseGlob("assets/transliterator/dist/*.html")),
	}
	e.Renderer = t

	e.GET("/", Index)
	e.GET("/about", About)
	e.POST("/", handler.HomeHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

//Index hander function
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "Transliterator")
}

//About handler function
func About(c echo.Context) error {
	return c.Render(http.StatusOK, "about", "About Page")
}

package main

import (
	"io"
	"net/http"
	"text/template"
	"transliterator/handler"

	"github.com/labstack/echo"
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

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
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

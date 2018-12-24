package main

import (
	"io"
	"text/template"
	"transliterator/handler"

	"github.com/labstack/echo"
)

// Template struct
type Template struct {
	templates *template.Template
}

// Render interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	//e.File registers new route with static file to serve
	e.File("/", "views/index.html")
	e.File("/about", "views/about.html")
	e.POST("/", handler.HomeHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

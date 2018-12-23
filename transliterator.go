package main

import (
	"io"
	"net/http"
	"strings"
	"text/template"

	handler "transliterator/handler"

	"github.com/labstack/echo"
)

var mGreek = map[string]string{"α": "a", "β": "b", "γ": "g", "δ": "d", "ε": "e", "ζ": "z", "η": "e", "θ": "th", "ι": "i", "κ": "k", "λ": "l", "μ": "m", "ν": "n", "ξ": "ks", "ο": "o", "π": "p", "ρ": "r", "σ": "s", "ς": "s", "τ": "t", "υ": "y", "φ": "ph", "χ": "ch", "ψ": "ps", "ω": "o"}
var mHebrew = map[string]string{"א": "'", "ב": "b", "ג": "g", "ד": "d", "ה": "h", "ו": "w", "ז": "z", "ח": "ch", "ט": "t", "י": "y", "כ": "k", "ך": "k", "ל": "l", "מ": "m", "ם": "m", "נ": "n", "ן": "n", "ס": "s", "ע": "'", "פ": "p", "ף": "ph", "צ": "ts", "ץ": "ts", "ק": "q", "ר": "r", "שׁ": "s", "שׂ": "sh", "ת": "th"}

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
	e.GET("/", handler.HomeHandler)
	e.GET("/words/:word", transliterate)
	e.Logger.Fatal(e.Start(":1323"))
}

//Function to handle language input and transliteration
func transliterate(c echo.Context) error {
	// lang := c.FormValue("language")
	lang := "GREEK"
	lang = strings.ToUpper(lang)
	text := c.Param("word")
	var str string
	if lang == "GREEK" {
		for _, value := range text {
			letter := string(value)
			str += mGreek[letter]
		}
	}
	return c.String(http.StatusOK, str)
}

package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

var mGreek = map[string]string{"α": "a", "β": "b", "γ": "g", "δ": "d", "ε": "e", "ζ": "z", "η": "e", "θ": "th", "ι": "i", "κ": "k", "λ": "l", "μ": "m", "ν": "n", "ξ": "ks", "ο": "o", "π": "p", "ρ": "r", "σ": "s", "ς": "s", "τ": "t", "υ": "y", "φ": "ph", "χ": "ch", "ψ": "ps", "ω": "o"}
var mHebrew = map[string]string{"א": "'", "ב": "b", "ג": "g", "ד": "d", "ה": "h", "ו": "w", "ז": "z", "ח": "ch", "ט": "t", "י": "y", "כ": "k", "ך": "k", "ל": "l", "מ": "m", "ם": "m", "נ": "n", "ן": "n", "ס": "s", "ע": "'", "פ": "p", "ף": "ph", "צ": "ts", "ץ": "ts", "ק": "q", "ר": "r", "שׁ": "s", "שׂ": "sh", "ת": "th"}

//HomeHandler path controller
func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "HOME",
	})
}

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

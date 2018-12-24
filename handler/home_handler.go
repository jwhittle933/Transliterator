package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

//HomeHandler path controller
func HomeHandler(c echo.Context) error {
	// Please note the the second parameter "index.html" is the template name and should
	// be equal to the value stated in the {{ define }} statement in "view/home.html"
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "HOME",
	})
}

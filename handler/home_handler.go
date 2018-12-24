package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

//HomeHandler path controller
func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "HOME",
	})
}

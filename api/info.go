package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//Hello
func GetAPIInfo(c echo.Context) error {
	return c.String(http.StatusOK, "Service run normally")
}

//Hello
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Bunjiu")
}

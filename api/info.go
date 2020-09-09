package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//Hello
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

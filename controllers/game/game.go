package game

import (
	"net/http"

	"github.com/labstack/echo"
)

func AddGame(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, "Not Implemented")
}

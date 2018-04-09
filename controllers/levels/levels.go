package levels

import (
	"database/sql"
	"net/http"

	"api.meet.the/components/database"
	"api.meet.the/components/response"
	model "api.meet.the/models"
	"github.com/labstack/echo"
)

func GetLevels(c echo.Context) error {

	l := model.Level{}

	levels, err := l.GetLevels(database.DB)

	if err != nil {

		switch err {
		case sql.ErrNoRows:
			return response.ErrorBadRequestWithKey(c, "User.Error.NotFound")
		}

		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return response.Success(c, levels)

}

package levels

import (
	"database/sql"

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
			return response.ErrorBadRequestWithKey(c, "SQL.Error.NoRows")
		}

		return response.ErrorBadRequestWithKey(c, "Application.Error.Unknown")

	}

	return response.Success(c, levels)

}

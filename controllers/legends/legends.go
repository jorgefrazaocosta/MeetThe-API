package legends

import (
	"database/sql"

	database "api.meet.the/components/database"
	"api.meet.the/components/response"
	model "api.meet.the/models"

	"github.com/labstack/echo"
)

func GetLegends(c echo.Context) error {

	l := model.Legend{}

	legends, err := l.GetLegends(database.DB)

	if err != nil {

		switch err {
		case sql.ErrNoRows:
			return response.ErrorBadRequestWithKey(c, "SQL.Error.NoRows")
		}

		return response.ErrorBadRequestWithKey(c, "Application.Error.Unknown")

	}

	return response.Success(c, legends)

}

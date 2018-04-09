package legends

import (
	"database/sql"
	"net/http"

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
			return response.ErrorBadRequestWithKey(c, "User.Error.NotFound")
		}

		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return response.Success(c, legends)

}

package question

import (
	"database/sql"
	"net/http"

	"api.meet.the/components/database"
	"api.meet.the/components/response"
	"api.meet.the/components/validator"
	model "api.meet.the/models"
	"github.com/labstack/echo"
)

func GetQuestion(c echo.Context) error {

	g := new(model.Game)
	q := new(model.Question)

	if err := c.Bind(g); err != nil {
		return response.ErrorBadRequestWithKey(c, "Application.Error.Unknown")
	}

	if err := validator.ValidateStruct(c, g); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := q.GetRandomQuestion(g.PersonID, g.Level, database.DB)

	if err != nil {

		switch err {
		case sql.ErrNoRows:
			return response.ErrorBadRequestWithKey(c, "User.Error.NotFound")
		}

		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return response.Success(c, q)

}

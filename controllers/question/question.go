package question

import (
	"database/sql"

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
		return response.ErrorBadRequestWithKey(c, "Application.Validation.Error")
	}

	err := q.GetRandomQuestion(g.PersonID, g.Level)

	if err != nil {

		switch err {
		case sql.ErrNoRows:
			return response.ErrorBadRequestWithKey(c, "SQL.Error.NoRows")
		}

		return response.ErrorBadRequestWithKey(c, "Application.Error.Unknown")

	}

	return response.Success(c, q)

}

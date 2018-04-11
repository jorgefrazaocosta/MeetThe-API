package game

import (
	"database/sql"
	"net/http"

	"api.meet.the/components/database"
	"api.meet.the/components/response"
	"api.meet.the/components/validator"
	model "api.meet.the/models"
	"github.com/labstack/echo"
)

func AddGame(c echo.Context) error {

	g := new(model.Game)

	if err := c.Bind(g); err != nil {
		return response.ErrorBadRequestWithKey(c, "Application.Error.Unknown")
	}

	if err := validator.ValidateStruct(c, g); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := g.AddNewGame(database.DB)

	if err != nil {

		switch err {
		case sql.ErrNoRows:
			return response.ErrorBadRequestWithKey(c, "User.Error.NotFound")
		}

		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return response.Success(c, true)

}

func RegisterQuestionResult(c echo.Context) error {

	gt := new(model.GameTrack)

	if err := c.Bind(gt); err != nil {
		return response.ErrorBadRequestWithKey(c, "Application.Error.Unknown")
	}

	if err := validator.ValidateStruct(c, gt); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := gt.RegisterGameAnswer(database.DB)

	if err != nil {

		switch err {
		case sql.ErrNoRows:
			return response.ErrorBadRequestWithKey(c, "User.Error.NotFound")
		}

		return c.JSON(http.StatusBadRequest, err.Error())

	}

	if result {
		return response.Success(c, true)
	}

	q := new(model.Question)
	q.GetUnansweredQuestion(gt.PeopleQuestionID, database.DB)

	// Also send email to user and to admin

	return response.Success(c, q)

}

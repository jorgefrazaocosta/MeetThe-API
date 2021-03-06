package game

import (
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
		return response.ErrorBadRequestWithKey(c, "Application.Validation.Error")
	}

	err := g.AddNewGame()

	if err != nil {
		return response.ErrorBadRequestWithKey(c, "Add.New.Game.Error")
	}

	return response.Success(c, true)

}

func RegisterQuestionResult(c echo.Context) error {

	gt := new(model.GameTrack)

	if err := c.Bind(gt); err != nil {
		return response.ErrorBadRequestWithKey(c, "Application.Error.Unknown")
	}

	if err := validator.ValidateStruct(c, gt); err != nil {
		return response.ErrorBadRequestWithKey(c, "Application.Validation.Error")
	}

	result, err := gt.RegisterGameAnswer()

	if err != nil {
		return response.ErrorBadRequestWithKey(c, "Register.User.Response.Error")
	}

	if result {
		return response.Success(c, true)
	}

	q := new(model.Question)
	err = q.GetUnansweredQuestion(gt.PeopleQuestionID)

	if err != nil {
		return response.ErrorBadRequestWithKey(c, "Get.Unanswered.Question.Error")
	}

	// Also send email to user and to admin

	return response.Success(c, q)

}

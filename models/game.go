package models

import (
	"api.meet.the/components/database"
)

type Game struct {
	PersonID int `form:"person" query:"person" validate:"required"`
	Level    int `form:"level" query:"level" validate:"required"`
}

type GameTrack struct {
	GameID           int `form:"gameId" validate:"required"`
	PeopleID         int `form:"peopleId" validate:"required"`
	PeopleQuestionID int `form:"peopleQuestionId" validate:"required"`
	AnswerID         int `form:"answerId" validate:"required"`
}

func (g *Game) AddNewGame() error {

	stmt, err := database.DB.Prepare("INSERT INTO Games (personId, levelId) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(g.PersonID, g.Level)
	if err != nil {
		return err
	}

	return nil

}

func (gt *GameTrack) RegisterGameAnswer() (bool, error) {

	result, err := gt.getAnswerConfirmation()

	if err != nil {
		return false, err
	}

	stmt, err := database.DB.Prepare("INSERT INTO GameTrack (gameId, peopleId, peopleQuestionId, result) VALUES (?,?,?,?)")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(gt.GameID, gt.PeopleID, gt.PeopleQuestionID, result)
	if err != nil {
		return false, err
	}

	return result, nil

}

func (gt *GameTrack) getAnswerConfirmation() (bool, error) {

	var result = false

	err := database.DB.QueryRow("SELECT isCorrect FROM Answers WHERE id = ?", gt.AnswerID).Scan(&result)

	if err != nil {
		return result, err
	}

	return result, nil

}

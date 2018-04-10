package models

import "database/sql"

type Game struct {
	PersonID int `form:"person" query:"person" validate:"required"`
	Level    int `form:"level" query:"level" validate:"required"`
}

type GameTrack struct {
	GameID           int  `form:"gameId" validate:"required"`
	PeopleID         int  `form:"peopleId" validate:"required"`
	PeopleQuestionID int  `form:"peopleQuestionId" validate:"required"`
	Result           bool `form:"result"`
}

func (g *Game) AddNewGame(db *sql.DB) error {

	stmt, err := db.Prepare("INSERT INTO Games (personId, levelId) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(g.PersonID, g.Level)
	if err != nil {
		panic(err.Error())
	}

	return nil

}

func (gt *GameTrack) RegisterGameAnswer(db *sql.DB) error {

	stmt, err := db.Prepare("INSERT INTO GameTrack (gameId, peopleId, peopleQuestionId, result) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(gt.GameID, gt.PeopleID, gt.PeopleQuestionID, gt.Result)
	if err != nil {
		panic(err.Error())
	}

	return nil

}

package models

import "api.meet.the/components/database"

type Level struct {
	LevelID     int    `json:"levelId" form:"levelId" binding:"required" validate:"required"`
	Description string `json:"description"`
}

func (l *Level) GetLevels() ([]Level, error) {

	var levels []Level

	results, err := database.DB.Query("SELECT id, description FROM Levels ORDER BY sort")
	if err != nil {
		return nil, err
	}

	for results.Next() {

		var level Level

		err = results.Scan(&level.LevelID, &level.Description)

		if err != nil {
			return nil, err
		}

		levels = append(levels, level)

	}

	return levels, nil

}

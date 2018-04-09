package models

type Game struct {
	PersonID int `query:"person" validate:"required"`
	Level    int `query:"level" validate:"required"`
}

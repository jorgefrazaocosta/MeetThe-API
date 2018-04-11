package models

type Answer struct {
	ID     int    `json:"answerId,omitempty" form:"answerdId"`
	Answer string `json:"answer" form:"answer"`
}

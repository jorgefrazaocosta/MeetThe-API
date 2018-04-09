package models

type ServerError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

package response

import (
	"net/http"

	"github.com/labstack/echo"

	localization "api.meet.the/components/localization"
)

func Success(c echo.Context, payload interface{}) error {
	return SuccessWithCode(c, http.StatusOK, payload)
}

func SuccessWithCode(c echo.Context, statusCode int, payload interface{}) error {

	finalPayload := map[string]interface{}{"data": payload}

	return c.JSON(statusCode, finalPayload)

}

func Error(c echo.Context, statusCode int, message string) error {

	errorPlayload := map[string]interface{}{"code": statusCode, "message": message}
	finalPayload := map[string]interface{}{"error": errorPlayload}

	return c.JSON(statusCode, finalPayload)

}

func ErrorKey(c echo.Context, statusCode int, key string) error {

	stringLocalized := localization.T(c.Request().Header.Get("Accept-Language"), key)
	return Error(c, statusCode, stringLocalized)

}

func ErrorBadRequestWithKey(c echo.Context, key string) error {
	return Error(c, http.StatusBadRequest, key)
}

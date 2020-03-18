package drivers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"nikan.dev/pronto/internals/exception"
)

func EchoExceptionDriver(err error, c echo.Context) {
	c.Logger().Error(err)
	if he, ok := err.(*echo.HTTPError); ok {
		str, ok := he.Message.(string)
		if ok != true {
			str = "something went wrong"
		}
		c.JSON(he.Code, struct {
			Message string `json:"message"`
		}{str})
		return
	}
	code := http.StatusInternalServerError
	exc, ok := err.(*exception.Exception)
	if ok != true {
		c.JSON(code, struct {
			Message string `json:"message"`
		}{err.Error()})
		return
	}
	c.JSON(getHttpStatusCode(exc.Status), exc)
}

func getHttpStatusCode(status exception.Status) int {
	switch status {
	case exception.NotFound:
		return 404
	case exception.AccessDenied:
		return 403
	case exception.AlreadyExists:
		return 409
	case exception.InvalidInput:
		return 422
	case exception.ServerError:
		return 500
	default:
		return 500
	}
}

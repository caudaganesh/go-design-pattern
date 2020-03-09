package echoserver

import (
	"github.com/labstack/echo"
)

// Helpers holds all the helpers
type Helpers struct {
	Response ResponseHelper
}

// ResponseHelper holds response helpers methods
type ResponseHelper interface {
	ResponseData(c echo.Context, args interface{}) error
	ErrorResponse(c echo.Context, httpStatus int, err error) error
}

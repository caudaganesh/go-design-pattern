package echoserver

import (
	"github.com/labstack/echo"
)

// UserHandler user handler
type UserHandler interface {
	GetUser(ctx echo.Context) error
}

package echorest

import (
	"github.com/labstack/echo"
)

// UserController user controller
type UserController interface {
	GetUser(ctx echo.Context) error
}

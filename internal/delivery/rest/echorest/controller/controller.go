package controller

import (
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest"
)

// Controller holds all the controllers
type Controller struct {
	UserController echorest.UserController
}

// NewController init new controller
func NewController(userController echorest.UserController) *Controller {
	return &Controller{
		userController,
	}
}

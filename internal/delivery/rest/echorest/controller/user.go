package controller

import (
	"errors"
	"strconv"

	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/constants"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest/helper"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/restentity"
	"github.com/labstack/echo"
)

type userController struct {
	handler echorest.UserHandler
	helper  echorest.Helpers
}

// NewUserController init user controller for echo
func NewUserController(handler echorest.UserHandler) echorest.UserController {
	helper := echorest.Helpers{
		Response: helper.NewResponseHelper(),
	}
	userCtrl := &userController{
		handler: handler,
		helper:  helper,
	}

	return userCtrl
}

func (u *userController) GetUser(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param(constants.UserID))
	if err != nil {
		return errors.New("invalid user id")
	}

	req := restentity.GetUserRequest{ID: userID}
	res, err := u.handler.GetUser(req)
	if err != nil {
		return err
	}

	return u.helper.Response.ResponseData(ctx, res)
}

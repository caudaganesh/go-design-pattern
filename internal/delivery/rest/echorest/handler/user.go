package handler

import (
	"errors"
	"strconv"

	"github.com/caudaganesh/go-design-pattern/internal/delivery"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/deliveryentity"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/constants"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest/helper"
	"github.com/labstack/echo"
)

type userHandler struct {
	adapter delivery.UserAdapter
	helper  echorest.Helpers
}

// NewUserHandler init user handler for echo
func NewUserHandler(adapter delivery.UserAdapter) echorest.UserHandler {
	helper := echorest.Helpers{
		Response: helper.NewResponseHelper(),
	}
	userCtrl := &userHandler{
		adapter: adapter,
		helper:  helper,
	}

	return userCtrl
}

func (u *userHandler) GetUser(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param(constants.UserID))
	if err != nil {
		return errors.New("invalid user id")
	}

	req := deliveryentity.GetUserRequest{ID: userID}
	res, err := u.adapter.GetUser(req)
	if err != nil {
		return err
	}

	return u.helper.Response.ResponseData(ctx, res)
}

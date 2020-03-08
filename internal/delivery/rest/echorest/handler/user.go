package handler

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/restentity"
)

type userHandler struct {
	userUC internal.UserUC
}

// NewUserHandler init new user handler
func NewUserHandler(userUC internal.UserUC) echorest.UserHandler {
	return &userHandler{
		userUC,
	}
}

func (u *userHandler) GetUser(request restentity.GetUserRequest) (restentity.GetUserResponse, error) {
	res := restentity.GetUserResponse{}
	user, err := u.userUC.GetUser(request.ID)
	if err != nil {
		return res, err
	}

	res.ID = user.ID
	res.UserName = user.UserName
	res.SetFullName(user.FirstName, user.LastName)

	return res, nil
}

package adapter

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/delivery"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/deliveryentity"
)

type userHandler struct {
	userUC internal.UserUC
}

// NewUserHandler init new user adapter
func NewUserHandler(userUC internal.UserUC) delivery.UserAdapter {
	return &userHandler{
		userUC,
	}
}

func (u *userHandler) GetUser(request deliveryentity.GetUserRequest) (deliveryentity.GetUserResponse, error) {
	res := deliveryentity.GetUserResponse{}
	user, err := u.userUC.GetUser(request.ID)
	if err != nil {
		return res, err
	}

	res.ID = user.ID
	res.UserName = user.UserName
	res.SetFullName(user.FirstName, user.LastName)

	return res, nil
}

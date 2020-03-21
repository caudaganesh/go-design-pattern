package adapter

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/deliveryentity"
)

type userHandlerAdapter struct {
	userUC internal.UserUC
}

// NewUserHandlerAdapter init new user adapter
func NewUserHandlerAdapter(userUC internal.UserUC) UserAdapter {
	return &userHandlerAdapter{
		userUC,
	}
}

func (u *userHandlerAdapter) GetUser(request deliveryentity.GetUserRequest) (deliveryentity.GetUserResponse, error) {
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

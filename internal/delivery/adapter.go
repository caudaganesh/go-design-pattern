package delivery

import (
	"github.com/caudaganesh/go-design-pattern/internal/delivery/deliveryentity"
)

// UserAdapter user handler
type UserAdapter interface {
	GetUser(request deliveryentity.GetUserRequest) (deliveryentity.GetUserResponse, error)
}

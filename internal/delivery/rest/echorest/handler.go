package echorest

import (
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/restentity"
)

// UserHandler user handler
type UserHandler interface {
	GetUser(request restentity.GetUserRequest) (restentity.GetUserResponse, error)
}

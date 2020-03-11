package internal

import (
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

// UserUC holds user use case methods
type UserUC interface {
	GetUser(ID int) (entity.User, error)
}

// AuthUC contains all the auth usecase methods
type AuthUC interface {
	Authorize(token string) (entity.Auth, error)
}

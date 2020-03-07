package internal

import (
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

// UserUC holds user use case methods
type UserUC interface {
	GetUser(ID int) (entity.User, error)
}

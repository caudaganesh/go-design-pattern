package internal

import (
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

// UserDBRepo holds mysql repo methods
type UserDBRepo interface {
	Find(ID int) (entity.User, error)
}

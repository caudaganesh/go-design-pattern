package mysql

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

type userDBRepo struct{}

// NewUserRepo init new user repo
func NewUserRepo() internal.UserDBRepo {
	return &userDBRepo{}
}

func (u *userDBRepo) Find(ID int) (entity.User, error) {
	return entity.User{}, nil
}

package usecase

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

type userUC struct {
	db internal.UserDBRepo
}

// NewUserUC init new user uc
func NewUserUC(db internal.UserDBRepo) internal.UserUC {
	return &userUC{db}
}

func (u *userUC) GetUser(ID int) (entity.User, error) {
	user, err := u.db.Find(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

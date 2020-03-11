package usecase

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

type authGoogle struct{}

// NewAuthGoogle init new auth google
func NewAuthGoogle() internal.AuthUC {
	return &authGoogle{}
}

func (a *authGoogle) Authorize(token string) (entity.Auth, error) {
	return entity.Auth{}, nil
}

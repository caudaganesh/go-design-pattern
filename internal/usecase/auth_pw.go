package usecase

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

type authPw struct{}

// NewAuthPW init new auth pw
func NewAuthPW() internal.AuthUC {
	return &authPw{}
}

func (a *authPw) Authorize(token string) (entity.Auth, error) {
	return entity.Auth{}, nil
}

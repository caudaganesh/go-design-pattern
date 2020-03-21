package composite

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

// AuthUserCompose composition of auth and user uc
type AuthUserCompose struct {
	internal.AuthUC
	internal.UserUC
}

// UserAuthComposeEntity composition of auth and user entity
type UserAuthComposeEntity struct {
	entity.User
	entity.Auth
}

// UserCategory user category
type UserCategory struct {
	Category       string
	ParentCategory *UserCategory
	ChildCategory  *UserCategory
}

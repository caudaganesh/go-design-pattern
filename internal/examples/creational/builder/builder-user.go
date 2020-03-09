package builder

import (
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

// UserBuilder user builder implementation
type UserBuilder struct {
	userType entity.UserType
}

// SetLevel user builder level setter
func (u *UserBuilder) SetLevel() UserTypeBuilder {
	u.userType.AccessLevel = entity.AccessLevel1
	return u
}

// SetType user builder type setter
func (u *UserBuilder) SetType() UserTypeBuilder {
	u.userType.Type = entity.UserTypeUser
	return u
}

// GetUserType return the user user type
func (u *UserBuilder) GetUserType() entity.UserType {
	return u.userType
}

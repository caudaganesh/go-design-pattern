package builder

import (
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

// UserTypeBuilder contains user type build methods
type UserTypeBuilder interface {
	SetLevel() UserTypeBuilder
	SetType() UserTypeBuilder
	GetUserType() entity.UserType
}

// UserTypeDirector director for user type
type UserTypeDirector struct {
	builder UserTypeBuilder
}

// SetBuilder set user type director builder
func (u *UserTypeDirector) SetBuilder(b UserTypeBuilder) {
	u.builder = b
}

// Construct constructor for user type director
func (u *UserTypeDirector) Construct() {
	u.builder.
		SetLevel().
		SetType()
}

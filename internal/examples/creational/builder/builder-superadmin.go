package builder

import (
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

// SuperAdminBuilder super admin builder implementation
type SuperAdminBuilder struct {
	userType entity.UserType
}

// SetLevel super admin builder level setter
func (s *SuperAdminBuilder) SetLevel() UserTypeBuilder {
	s.userType.AccessLevel = entity.AccessLevel3
	return s
}

// SetType super admin builder type setter
func (s *SuperAdminBuilder) SetType() UserTypeBuilder {
	s.userType.Type = entity.UserTypeAdmin
	return s
}

// GetUserType return the super admin user type
func (s *SuperAdminBuilder) GetUserType() entity.UserType {
	return s.userType
}

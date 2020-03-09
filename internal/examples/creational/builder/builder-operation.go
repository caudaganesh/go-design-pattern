package builder

import (
	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

// OperationAdminBuilder operat operation admin builder implementation
type OperationAdminBuilder struct {
	userType entity.UserType
}

// SetLevel operation admin builder level setter
func (o *OperationAdminBuilder) SetLevel() UserTypeBuilder {
	o.userType.AccessLevel = entity.AccessLevel1
	return o
}

// SetType operation admin builder type setter
func (o *OperationAdminBuilder) SetType() UserTypeBuilder {
	o.userType.Type = entity.UserTypeAdmin
	return o
}

// GetUserType return the operation admin user type
func (o *OperationAdminBuilder) GetUserType() entity.UserType {
	return o.userType
}

package builder

import (
	"reflect"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

func TestOperationAdminBuilder_SetLevel(t *testing.T) {
	o := &OperationAdminBuilder{}
	res := o.SetLevel()

	if !reflect.DeepEqual(res, o) {
		t.Error("returns invalid builder", res)
	}

	if o.userType.AccessLevel != 1 {
		t.Error("invalid user level, should be 1", o.userType.AccessLevel)
	}
}

func TestOperationAdminBuilder_SetType(t *testing.T) {
	o := &OperationAdminBuilder{}
	res := o.SetType()

	if !reflect.DeepEqual(res, o) {
		t.Error("returns invalid builder", res)
	}

	if o.userType.Type != entity.UserTypeAdmin {
		t.Error("invalid user type", o.userType.Type)
	}
}

func TestOperationAdminBuilder_GetUserType(t *testing.T) {
	o := &OperationAdminBuilder{}
	o.SetLevel()
	o.SetType()

	res := o.GetUserType()

	if !reflect.DeepEqual(res, o.userType) {
		t.Error("invalid user getter", res, o.userType)
	}
}

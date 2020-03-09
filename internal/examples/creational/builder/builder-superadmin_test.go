package builder

import (
	"reflect"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

func TestSuperAdminBuilder_SetLevel(t *testing.T) {
	s := &SuperAdminBuilder{}
	res := s.SetLevel()

	if !reflect.DeepEqual(res, s) {
		t.Error("returns invalid builder", res, s)
	}

	if s.userType.AccessLevel != entity.AccessLevel3 {
		t.Error("invalid user level, should be 3", s.userType.AccessLevel)
	}
}

func TestSuperAdminBuilder_SetType(t *testing.T) {
	s := &SuperAdminBuilder{}
	res := s.SetType()

	if !reflect.DeepEqual(s, res) {
		t.Error("returns invalid builder", s, res)
	}

	if s.userType.Type != entity.UserTypeAdmin {
		t.Error("invalid user type, should be ADMIN", s.userType.Type)
	}
}

func TestSuperAdminBuilder_GetUserType(t *testing.T) {
	s := &SuperAdminBuilder{}
	s.SetLevel()
	s.SetType()

	res := s.GetUserType()

	if !reflect.DeepEqual(res, s.userType) {
		t.Error("invalid user getter", res, s.userType)
	}
}

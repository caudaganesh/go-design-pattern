package builder

import (
	"reflect"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

func TestUserBuilder_SetLevel(t *testing.T) {
	u := &UserBuilder{}
	res := u.SetLevel()

	if !reflect.DeepEqual(res, u) {
		t.Error("returns invalid builder")
	}

	if u.userType.AccessLevel != 1 {
		t.Error("invalid user level, should be 1", u.userType.AccessLevel)
	}
}

func TestUserBuilder_SetType(t *testing.T) {
	u := &UserBuilder{}
	res := u.SetType()

	if !reflect.DeepEqual(u, res) {
		t.Error("returns invalid builder")
	}

	if u.userType.Type != entity.UserTypeUser {
		t.Error("invalid user type, should be USER", u.userType.Type)
	}

}

func TestUserBuilder_GetUserType(t *testing.T) {
	u := &UserBuilder{}
	u.SetLevel()
	u.SetType()

	res := u.GetUserType()

	if !reflect.DeepEqual(res, u.userType) {
		t.Error("invalid user getter")
	}
}

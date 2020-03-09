package builder

import (
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

func TestUserTypeDirector_SetBuilder(t *testing.T) {
	u := &UserTypeDirector{}
	ub := &UserBuilder{}
	u.SetBuilder(ub)
	u.Construct()

	if u.builder.GetUserType().Type != entity.UserTypeUser {
		t.Error("building wrong user type", u)
	}

	if u.builder.GetUserType().AccessLevel != entity.AccessLevel1 {
		t.Error("building wrong user access level", u)
	}
}

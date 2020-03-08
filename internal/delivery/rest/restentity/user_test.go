package restentity

import "testing"

func TestGetUserResponse_SetFullName(t *testing.T) {
	u := GetUserResponse{}
	u.SetFullName("firstName", "lastName")

	if u.FullName != "firstName lastName" {
		t.Error("invalid fullname setter")
	}
}

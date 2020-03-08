package restentity

import (
	"fmt"
)

// GetUserRequest get user request
type GetUserRequest struct {
	ID int `json:"id"`
}

// GetUserResponse get user response
type GetUserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
}

// SetFullName full name setter
func (g *GetUserResponse) SetFullName(firstName string, lastName string) {
	g.FullName = fmt.Sprintf("%s %s", firstName, lastName)
}

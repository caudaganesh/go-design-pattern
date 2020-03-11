package entity

// User holds user properties
type User struct {
	ID        int
	UserName  string
	FirstName string
	LastName  string
	UserType  UserType
}

// UserType holds user type
type UserType struct {
	Type        string
	AccessLevel int
}

// user type
const (
	UserTypeAdmin   = "ADMIN"
	UserTypeUser    = "USER"
	UserTypePartner = "PARTNER"
)

// access level
const (
	AccessLevel1 = 1
	AccessLevel2 = 2
	AccessLevel3 = 3
)

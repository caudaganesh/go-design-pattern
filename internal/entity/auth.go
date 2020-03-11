package entity

// Auth holds auth struct
type Auth struct {
	Token    string
	UserName string
}

// auth type
const (
	AuthTypeGoogle = "GOOGLE"
	AuthTypePw     = "PW"
)

package singleton

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/repository/psql"
	"github.com/caudaganesh/go-design-pattern/internal/usecase"
)

var userUCInstance internal.UserUC

// GetUserUCInstance get singleton instance
func GetUserUCInstance() internal.UserUC {
	if userUCInstance == nil {
		repo := psql.NewUserRepo()
		userUCInstance = usecase.NewUserUC(repo)
	}

	return userUCInstance
}

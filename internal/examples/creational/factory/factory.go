package factory

import (
	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
	"github.com/caudaganesh/go-design-pattern/internal/usecase"
)

// AuthFactory determine the auth method
func AuthFactory(authType string) internal.AuthUC {
	switch authType {
	case entity.AuthTypeGoogle:
		return usecase.NewAuthGoogle()
	case entity.AuthTypePw:
		return usecase.NewAuthPW()
	}

	return nil
}

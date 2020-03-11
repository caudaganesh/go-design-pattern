package factory

import (
	"reflect"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
	"github.com/caudaganesh/go-design-pattern/internal/usecase"
)

func TestAuthFactory(t *testing.T) {
	tests := []struct {
		authType string
		name     string
		want     internal.AuthUC
	}{
		{
			name:     "auth type google, return google auth",
			authType: entity.AuthTypeGoogle,
			want:     usecase.NewAuthGoogle(),
		},
		{
			name:     "auth type pw, return pw auth",
			authType: entity.AuthTypePw,
			want:     usecase.NewAuthPW(),
		},
		{
			name:     "invalid auth type, return nil",
			authType: "something else",
			want:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthFactory(tt.authType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

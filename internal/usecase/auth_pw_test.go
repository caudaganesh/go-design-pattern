package usecase

import (
	"reflect"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

func Test_authPw_Authorize(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		a       *authPw
		args    args
		want    entity.Auth
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAuthPW()
			got, err := a.Authorize(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("authPw.Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authPw.Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}

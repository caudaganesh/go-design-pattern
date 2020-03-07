package mysql

import (
	"reflect"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/entity"
)

func Test_userDBRepo_Find(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    entity.User
		wantErr bool
	}{
		// TODO: update this with the correct implementation checking
		{
			name: "empty implementation",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserRepo()
			got, err := u.Find(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDBRepo.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userDBRepo.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

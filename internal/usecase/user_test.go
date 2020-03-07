package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
	"github.com/caudaganesh/go-design-pattern/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestNewUserUC(t *testing.T) {
	userUC := NewUserUC(nil)
	if userUC == nil {
		t.Error("fail to init new user uc")
	}
}

func Test_userUC_GetUser(t *testing.T) {
	type mockGetUserFromDB struct {
		res entity.User
		err error
	}
	tests := []struct {
		name string
		id   int
		db   internal.UserDBRepo
		mockGetUserFromDB
		want    entity.User
		wantErr bool
	}{
		{
			name: "got error from db, return error",
			mockGetUserFromDB: mockGetUserFromDB{
				err: errors.New("some error"),
			},
			wantErr: true,
		},
		{
			name: "no error from db, return error",
			id:   111,
			mockGetUserFromDB: mockGetUserFromDB{
				res: entity.User{
					ID:       111,
					UserName: "username",
				},
			},
			want: entity.User{
				ID:       111,
				UserName: "username",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUserDB := mocks.NewMockUserDBRepo(ctrl)
			mockUserDB.EXPECT().
				Find(tt.id).
				Return(tt.mockGetUserFromDB.res, tt.mockGetUserFromDB.err)
			u := &userUC{
				db: mockUserDB,
			}
			got, err := u.GetUser(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUC.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUC.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

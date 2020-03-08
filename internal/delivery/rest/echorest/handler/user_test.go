package handler

import (
	"errors"
	"reflect"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/restentity"
	"github.com/caudaganesh/go-design-pattern/internal/entity"
	"github.com/caudaganesh/go-design-pattern/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestNewUserHandler(t *testing.T) {
	userHandler := NewUserHandler(nil)

	if userHandler == nil {
		t.Error("should init user handler")
	}
}

func Test_userHandler_GetUser(t *testing.T) {
	type mockGetUser struct {
		res entity.User
		err error
	}
	tests := []struct {
		name string
		mockGetUser
		request restentity.GetUserRequest
		want    restentity.GetUserResponse
		wantErr bool
	}{
		{
			name: "error get user from uc, return error",
			mockGetUser: mockGetUser{
				err: errors.New("some error"),
			},
			wantErr: true,
		},
		{
			name: "no error get user from uc, return response",
			mockGetUser: mockGetUser{
				res: entity.User{
					FirstName: "firstname",
					LastName:  "lastname",
					ID:        1,
					UserName:  "username",
				},
			},
			want: restentity.GetUserResponse{
				ID:       1,
				FullName: "firstname lastname",
				UserName: "username",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUserUC := mocks.NewMockUserUC(ctrl)
			mockUserUC.EXPECT().
				GetUser(tt.request.ID).
				Return(tt.mockGetUser.res, tt.mockGetUser.err)

			u := &userHandler{
				userUC: mockUserUC,
			}

			got, err := u.GetUser(tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("userHandler.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userHandler.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

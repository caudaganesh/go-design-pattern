package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest/mocks"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/restentity"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestNewUserController(t *testing.T) {
	userCtrl := NewUserController(nil)

	if userCtrl == nil {
		t.Error("should init the new user controller")
	}
}

func Test_userController_GetUser(t *testing.T) {
	type mockUserHandler struct {
		res restentity.GetUserResponse
		err error
	}

	tests := []struct {
		userID string
		mockUserHandler
		mockHelperErr error
		name          string
		wantErr       bool
	}{
		{
			name:    "fail to get user id from context, return error",
			wantErr: true,
		},
		{
			name:   "valid user id, error at handler, return error",
			userID: "1",
			mockUserHandler: mockUserHandler{
				err: errors.New("some error"),
			},
			wantErr: true,
		},
		{
			name:   "success get user, error parse response data, return error",
			userID: "1",
			mockUserHandler: mockUserHandler{
				res: restentity.GetUserResponse{
					ID:       1,
					FullName: "test123",
				},
			},
			mockHelperErr: errors.New("some error"),
			wantErr:       true,
		},
		{
			name:   "success get user, success parse response data, return nil",
			userID: "1",
			mockUserHandler: mockUserHandler{
				res: restentity.GetUserResponse{
					ID:       1,
					FullName: "test123",
				},
			},
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		mockUserHandler := mocks.NewMockUserHandler(ctrl)
		mockUserHandler.EXPECT().
			GetUser(gomock.Any()).
			Return(tt.mockUserHandler.res, tt.mockUserHandler.err)

		e := echo.New()
		r := e.Router()
		r.Add(http.MethodGet, "user/:user_id", func(echo.Context) error { return nil })
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/user/:user_id")
		c.SetParamNames("user_id")
		c.SetParamValues(tt.userID)

		mockRespHelper := mocks.NewMockResponseHelper(ctrl)
		mockRespHelper.EXPECT().
			ResponseData(c, tt.mockUserHandler.res).
			Return(tt.mockHelperErr)

		t.Run(tt.name, func(t *testing.T) {
			u := &userController{
				handler: mockUserHandler,
				helper: echorest.Helpers{
					Response: mockRespHelper,
				},
			}
			if err := u.GetUser(c); (err != nil) != tt.wantErr {
				t.Errorf("userController.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

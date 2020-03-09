package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caudaganesh/go-design-pattern/internal/delivery/deliveryentity"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/mocks"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echoserver"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestNewUserHandler(t *testing.T) {
	userCtrl := NewUserHandler(nil)

	if userCtrl == nil {
		t.Error("should init the new user handler")
	}
}

func Test_UserHandler_GetUser(t *testing.T) {
	type mockUserHandler struct {
		res deliveryentity.GetUserResponse
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
				res: deliveryentity.GetUserResponse{
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
				res: deliveryentity.GetUserResponse{
					ID:       1,
					FullName: "test123",
				},
			},
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		mockUserHandler := mocks.NewMockUserAdapter(ctrl)
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
			u := &userHandler{
				adapter: mockUserHandler,
				helper: echoserver.Helpers{
					Response: mockRespHelper,
				},
			}
			if err := u.GetUser(c); (err != nil) != tt.wantErr {
				t.Errorf("UserHandler.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

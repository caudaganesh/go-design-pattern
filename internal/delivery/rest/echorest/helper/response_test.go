package helper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func Test_responseHelper_ResponseData(t *testing.T) {

	tests := []struct {
		name    string
		args    interface{}
		r       *responseHelper
		wantErr bool
	}{
		{
			name: "error parse json, return error response",
			args: make(chan int),
		},
		{
			name: "success parse json, return error response",
			args: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			helper := NewResponseHelper()
			if err := helper.ResponseData(c, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("responseHelper.ResponseData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

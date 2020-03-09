package helper

import (
	"encoding/json"
	"net/http"

	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echoserver"
	"github.com/labstack/echo"
)

type responseHelper struct{}

// NewResponseHelper init new response helper
func NewResponseHelper() echoserver.ResponseHelper {
	return &responseHelper{}
}

func (r *responseHelper) ResponseData(c echo.Context, args interface{}) error {
	data, err := json.Marshal(args)
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func (r *responseHelper) ErrorResponse(c echo.Context, httpStatus int, err error) error {
	return c.JSON(httpStatus, err)
}

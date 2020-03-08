package routes

import (
	"fmt"

	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/constants"
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest/controller"
	"github.com/labstack/echo"
)

//Register route
func Register(e *echo.Echo, ctrl *controller.Controller) {

	// you can register middleware here
	r := e.Group("v1")
	r.GET(fmt.Sprintf("/user/%s", constants.UserID), ctrl.UserController.GetUser)
}

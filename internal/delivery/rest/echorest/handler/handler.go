package handler

import (
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echorest"
)

// Handler holds all the handlers
type Handler struct {
	UserHandler echorest.UserHandler
}

// NewHandler init new handler
func NewHandler(userHandler echorest.UserHandler) *Handler {
	return &Handler{
		userHandler,
	}
}

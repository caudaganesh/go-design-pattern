package handler

import (
	"github.com/caudaganesh/go-design-pattern/internal/delivery/rest/echoserver"
)

// Handler holds all the handlers
type Handler struct {
	UserHandler echoserver.UserHandler
}

// NewHandler init new handler
func NewHandler(userHandler echoserver.UserHandler) *Handler {
	return &Handler{
		userHandler,
	}
}

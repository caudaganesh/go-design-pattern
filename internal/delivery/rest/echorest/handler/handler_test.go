package handler

import (
	"testing"
)

func TestNewHandler(t *testing.T) {
	ctrl := NewHandler(nil)

	if ctrl == nil {
		t.Error("should init the handler")
	}
}

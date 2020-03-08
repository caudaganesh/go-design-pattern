package controller

import (
	"testing"
)

func TestNewController(t *testing.T) {
	ctrl := NewController(nil)

	if ctrl == nil {
		t.Error("should init the controller")
	}
}

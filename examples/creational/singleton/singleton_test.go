package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	userUC := GetUserUCInstance()

	if userUC == nil {
		t.Error("expected to create new singleton after calling the GetUserUC")
	}
}

func TestGetInstanceAlreadyExists(t *testing.T) {
	currentUserUC := GetUserUCInstance()
	userUC := GetUserUCInstance()

	if currentUserUC != userUC {
		t.Error("get different instance while getting the singleton")
	}
}

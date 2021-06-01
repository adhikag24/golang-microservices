package domain

import (
	"testing"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	if user != nil {
		t.Error("we were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("we were expecting an errror when user id is 0")
	}

	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we were expecting 404 when user is not found")
	// }
}

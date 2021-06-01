package domain

import (
	"net/http"

	"github.com/adhikag24/golang-microservices/mvc/utils"

	"fmt"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Adhika", LastName: "Leon", Email: "adhikag167@gmail.com"},
	}

	UserDao = &userDao{}
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}

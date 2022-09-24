package services

import (
	"fmt"
	"github.com/Hemant1-6/golang-microservices/mvc/domain"
	"github.com/Hemant1-6/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*domain.User{
		1: {
			Id:        1,
			FirstName: "Hemant",
			LastName:  "Ahuja",
			Email:     "hemant.ahuja@gmail.com",
		},
	}
)

func GetUser(userId int64) (user *domain.User, err *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil,
		&utils.ApplicationError{
			Message: fmt.Sprintf("No user found with user id as %v", userId),
			Status:  http.StatusNotFound,
			Code:    "Not Found!",
		}
}

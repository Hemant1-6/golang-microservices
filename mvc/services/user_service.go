package services

import (
	"errors"
	"fmt"
	"github.com/Hemant1-6/golang-microservices/mvc/domain"
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

func GetUser(userId int64) (user *domain.User, err error) {
	if user := users[userId]; users != nil {
		return user, nil
	}
	return nil,
		errors.New(fmt.Sprintf("No user found with user id as %v", userId))
}

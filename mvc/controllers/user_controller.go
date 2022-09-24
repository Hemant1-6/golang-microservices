package controllers

import (
	"github.com/Hemant1-6/golang-microservices/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userid, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		return
	}
	user, err := services.GetUser(userid)
	println(user)
	if err != nil {
		panic(err)
	}

}

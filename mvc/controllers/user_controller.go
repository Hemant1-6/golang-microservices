package controllers

import (
	"encoding/json"
	"github.com/Hemant1-6/golang-microservices/mvc/services"
	"github.com/Hemant1-6/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userid, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.Status)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userid)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.Status)
		resp.Write(jsonValue)
		return
	}
	data, _ := json.Marshal(user)
	resp.Write(data)
	return
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/Hemant1-6/golang-microservices/mvc/services"
	"github.com/Hemant1-6/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userid, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}
	user, apiErr := services.GetUser(userid)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

package app

import (
	"github.com/Hemant1-6/golang-microservices/mvc/controllers"
)

func routes() {
	router.GET("/user/:user_id", controllers.GetUser)
}

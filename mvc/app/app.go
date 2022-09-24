package app

import (
	"github.com/Hemant1-6/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/user", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

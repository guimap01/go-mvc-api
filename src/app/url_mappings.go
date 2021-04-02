package app

import (
	"go-mvc/src/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}

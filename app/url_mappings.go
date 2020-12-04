package app

import "golang/micro/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
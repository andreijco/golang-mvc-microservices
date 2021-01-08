package app

import (
	"golang/micro/src/api/controllers/polo"
	"golang/micro/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
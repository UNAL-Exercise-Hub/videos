package routes

import (
	"github.com/UNWorkout/Videos/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/videos", controller.GetVideos)

	router.POST("/videos", controller.PostVideo)

	router.GET("/videos/:id", controller.GetVideoID)

	router.PUT("/videos/:id", controller.UpdateVideo)

	router.DELETE("/videos/:id", controller.DeleteVideo)

	category := router.Group("/categoria")
	{
		category.GET("/musculos", controller.GetMusculos)
		category.GET("/grupoMuscular", controller.GetGrupoMuscular)
		category.GET("/objetivo", controller.GetObjetivos)
		category.GET("/dificultad", controller.GetDificultad)
		category.GET("/equipamento", controller.GetEquipamento)
		category.GET("/disciplina", controller.GetDisciplina)
	}

	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

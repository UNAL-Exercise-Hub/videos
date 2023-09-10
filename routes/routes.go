package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "get",
		})
	})

	router.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "post",
		})
	})

	router.GET("/videos/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "get id",
			"id":      id,
		})
	})

	router.PUT("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "put",
		})
	})

	router.DELETE("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "del",
		})
	})

	category := router.Group("/categoria")
	{
		category.GET("/musculos", categoryM)
		category.GET("/grupoMuscular", categoryM)
		category.GET("/objetivo", categoryM)
		category.GET("/dificultad", categoryM)
		category.GET("/equipamento", categoryM)
		category.GET("/disciplina", categoryM)
	}

	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func categoryM(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "category",
	})
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func start() {
	router := gin.Default()
	router.GET("/", hello)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

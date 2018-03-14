package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Angular",
		})
	})
	router.Run()
}

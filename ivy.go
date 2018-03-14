package ivy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.New()

	r.GET("v1/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Angular",
		})
	})

	http.Handle("/", r)
}

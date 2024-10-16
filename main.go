package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//endpoint 1
	r.GET("/morning", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Good Morning Everyone!!",
		})
	})
	//endpoint 2
	r.GET("/night", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Good Night Everyone!!",
		})
	}) //endpoint 3
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Everyone!!",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080(for windows "localhost:8080")
}

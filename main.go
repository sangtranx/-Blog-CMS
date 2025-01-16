package main

import (
	"Blog-CMS/component/initialize"
	"github.com/gin-gonic/gin"
)

func main() {

	r := initialize.RunInit()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

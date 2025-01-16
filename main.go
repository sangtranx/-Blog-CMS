package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//dsn := os.Getenv("BLOG_CMS")
	//
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//db = db.Debug()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

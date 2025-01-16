package initialize

import "github.com/gin-gonic/gin"

func run() *gin.Engine {
	LoadConfig()

	InitMysql()
	InitRedis()

	r := InitRouter()

	return r
}

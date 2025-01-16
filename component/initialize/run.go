package initialize

import (
	"github.com/gin-gonic/gin"
)

func RunInit() *gin.Engine {
	LoadConfig()

	InitMysql()
	InitRedis()

	r := InitRouter()

	return r
}

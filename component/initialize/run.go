package initialize

import (
	"Blog-CMS/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RunInit() *gin.Engine {
	LoadConfig()
	InitLogger()
	common.Logger.Info("Config Log ok!!", zap.String("ok", "success"))

	InitMysql()
	InitRedis()

	r := InitRouter()

	return r
}

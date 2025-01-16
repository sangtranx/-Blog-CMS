package initialize

import (
	"Blog-CMS/common"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{

	var r *gin.Engine

	if common.Config.Server.Mode == "dev"{
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	return r
}

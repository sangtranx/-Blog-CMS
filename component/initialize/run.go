package initialize

import (
	"Blog-CMS/component/appctx"
	"github.com/gin-gonic/gin"
)

func RunInit() (*gin.Engine, appctx.AppContext) {
	LoadConfig()

	db := InitMysql()
	logger := InitLogger()
	redis := InitRedis()
	appCtx := appctx.NewAppContext(db, redis, logger)

	r := InitRouter()

	return r, appCtx
}

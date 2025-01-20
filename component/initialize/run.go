package initialize

import (
	"Blog-CMS/component/appctx"
	"github.com/gin-gonic/gin"
	"os"
)

func RunInit() (*gin.Engine, appctx.AppContext) {
	LoadConfig()

	db := InitMysql()
	logger := InitLogger()
	redis := InitRedis()
	secretKey := os.Getenv("SYSTEM_SECRET")

	appCtx := appctx.NewAppContext(db, redis, logger, secretKey)

	r := InitRouter()

	return r, appCtx
}

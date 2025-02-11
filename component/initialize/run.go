package initialize

import (
	"Blog-CMS/component/appctx"
	"Blog-CMS/component/pubsub/pblocal"
	"github.com/gin-gonic/gin"
	"os"
)

func RunInit() (*gin.Engine, appctx.AppContext) {
	LoadConfig()

	db := InitMysql()
	logger := InitLogger()
	redis := InitRedis()
	secretKey := os.Getenv("SYSTEM_SECRET")

	pb := pblocal.NewPubsub()
	appCtx := appctx.NewAppContext(db, redis, logger, pb, secretKey)

	r := InitRouter()

	return r, appCtx
}

package initialize

import (
	"Blog-CMS/common"
	"Blog-CMS/component/package/logger"
)

func InitLogger() {
	common.Logger = logger.NewLogger(common.Config.Logger)
}

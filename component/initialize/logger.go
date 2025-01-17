package initialize

import (
	"Blog-CMS/common"
	"Blog-CMS/component/package/logger"
)

func InitLogger() *logger.LoggerZap {
	return logger.NewLogger(common.Config.Logger)
}

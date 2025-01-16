package common

import (
	"Blog-CMS/component/package/logger"
	"Blog-CMS/component/package/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config *setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
)

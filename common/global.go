package common

import (
	"Blog-CMS/component/package/logger"
	"Blog-CMS/component/package/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *setting.Config
	Logger *logger.LoggerZap
	DB     *gorm.DB
	Rdb    *redis.Client
)

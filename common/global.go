package common

import (
	"Blog-CMS/component/package/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *setting.Config
	DB     *gorm.DB
	Rdb    *redis.Client
)

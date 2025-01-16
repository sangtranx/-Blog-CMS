package common

import (
	"Blog-CMS/component/package/setting"
	"gorm.io/gorm"
)

var (
	Config *setting.Config
	DB     *gorm.DB
)

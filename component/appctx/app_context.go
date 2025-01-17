package appctx

import (
	"Blog-CMS/component/package/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetRedisDBConnection() *redis.Client
	GetLogger() *logger.LoggerZap
}

type appCtx struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *logger.LoggerZap
}

func NewAppContext(db *gorm.DB, rdb *redis.Client, logger *logger.LoggerZap) *appCtx {
	return &appCtx{db: db}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB       { return ctx.db }
func (ctx *appCtx) GetRedisDBConnection() *redis.Client { return ctx.rdb }
func (ctx *appCtx) GetLogger() *logger.LoggerZap        { return ctx.logger }

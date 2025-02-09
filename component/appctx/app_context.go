package appctx

import (
	"Blog-CMS/component/package/logger"
	"Blog-CMS/component/pubsub"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetRedisDBConnection() *redis.Client
	GetLogger() *logger.LoggerZap
	GetPubsub() pubsub.Pubsub
	SecretKey() string
}

type appCtx struct {
	db        *gorm.DB
	rdb       *redis.Client
	logger    *logger.LoggerZap
	pubsub    pubsub.Pubsub
	secretKey string
}

func NewAppContext(db *gorm.DB, rdb *redis.Client, logger *logger.LoggerZap, pubsub pubsub.Pubsub, secretKey string) *appCtx {
	return &appCtx{
		db:        db,
		rdb:       rdb,
		logger:    logger,
		pubsub:    pubsub,
		secretKey: secretKey,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB       { return ctx.db }
func (ctx *appCtx) GetRedisDBConnection() *redis.Client { return ctx.rdb }
func (ctx *appCtx) GetLogger() *logger.LoggerZap        { return ctx.logger }
func (ctx *appCtx) GetPubsub() pubsub.Pubsub            { return ctx.pubsub }
func (ctx *appCtx) SecretKey() string                   { return ctx.secretKey }

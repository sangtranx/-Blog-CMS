package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func AddToBlackList(rdb *redis.Client, token string, expiry int) error {
	return rdb.Set(context.Background(), token, "revoked", time.Duration(expiry)*time.Second).Err()
}

func DeleteFromBlackList(rdb *redis.Client, token string) error {
	return rdb.Del(context.Background(), token).Err()
}

func IsTokenBlackListed(rdb *redis.Client, token string) bool {
	_, err := rdb.Get(context.Background(), token).Result()
	return err == nil
}

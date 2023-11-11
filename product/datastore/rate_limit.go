package datastore

import (
	"context"
	"fmt"
	"product/internal/constants"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RateLimiter *redis_rate.Limiter

func InitializeRateLimiter() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", viper.GetString(constants.REDIS_HOST), viper.GetString(constants.REDIS_PORT)),
	})
	_ = rdb.FlushDB(ctx).Err()

	RateLimiter = redis_rate.NewLimiter(rdb)
}

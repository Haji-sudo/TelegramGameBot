package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func InitRedisdb() (*redis.Client, context.Context) {
	opt, err := redis.ParseURL("redis://default:redispw@localhost:49153/0")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	ctx := context.Background()
	return rdb, ctx
}

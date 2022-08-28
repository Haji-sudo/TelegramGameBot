package DataBase

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InitRedisdb(user, pass, server, port, db string) *redis.Client {
	redisURL := fmt.Sprintf("redis://%v:%v@%v:%v/%v", user, pass, server, port, db)
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	return rdb
}

package DataBase

import (
	"github.com/go-redis/redis/v8"
)

func InitRedisdb() *redis.Client {
	opt, err := redis.ParseURL("redis://default:redispw@localhost:49153/0")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	return rdb
}

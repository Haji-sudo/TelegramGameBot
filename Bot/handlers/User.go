package handlers

import (
	db "dogegambling/config/DB"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
)

type UserRedis struct {
	UserID   int64     `json:"user_id"`
	Lock     bool      `json:"lock"`
	TimeSpam time.Time `json:"time_spam"`
}

var rdb *redis.Client
var ctx context.Context

func UserInit() {
	rdb, ctx = db.InitRedisdb()
}

func (u *UserRedis) CreateUser() bool {
	userdata, _ := json.Marshal(u)

	user, err := rdb.Set(ctx, strconv.FormatInt(u.UserID, 10), userdata, 0).Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)

	return true
}

func GetUser(userid int64) UserRedis {

	user, err := rdb.Get(ctx, strconv.FormatInt(userid, 10)).Result()

	if err != nil {
		log.Println(err)
	}
	if err == redis.Nil {
		return UserRedis{}
	}
	userdata := UserRedis{}
	json.Unmarshal([]byte(user), &userdata)
	return userdata
}

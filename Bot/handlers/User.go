package handlers

import (
	db "dogegambling/config/DB"
	"encoding/json"
	"strconv"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
)

type UserRedis struct {
	UserID   int64 `json:"user_id"`
	IsLock   bool  `json:"lock"`
	TimeSpam int64 `json:"time_spam"`
}

func (u *UserRedis) Bind(userid int64) {
	u.UserID = userid
	u.IsLock = false
	u.TimeSpam = time.Now().Unix()

}

var rdb *redis.Client
var ctx context.Context

func UserInit() {
	rdb, ctx = db.InitRedisdb()
}

func (u *UserRedis) CreateUser() {
	userdata, _ := json.Marshal(u)
	rdb.Set(ctx, strconv.FormatInt(u.UserID, 10), userdata, 0).Result()
}

func GetUser(userid int64) UserRedis {

	user, err := rdb.Get(ctx, strconv.FormatInt(userid, 10)).Result()

	if err == redis.Nil {
		return UserRedis{}
	}
	userdata := UserRedis{}
	json.Unmarshal([]byte(user), &userdata)
	return userdata
}

func (u UserRedis) Exist() bool {

	return u != UserRedis{}
}

func (u *UserRedis) update() {
	userdata, _ := json.Marshal(u)
	rdb.Set(ctx, strconv.FormatInt(u.UserID, 10), userdata, 0).Result()
}

func (u *UserRedis) UpdateTime() {
	u.TimeSpam = time.Now().Unix()

}

func (u *UserRedis) lock() {
	u.IsLock = true
	u.update()
}
func (u *UserRedis) unlock() {
	u.IsLock = false
	u.update()
}

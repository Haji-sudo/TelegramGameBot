package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	rdb *redis.Client
	ctx context.Context
	DB  *gorm.DB
)

func (h Handler) UserInit() {
	rdb, ctx, DB = h.RDB, h.CTX, h.DB
}

func (u UserRedis) CreateUser(userid int64) { //Create User in Both DB
	u.createinRedis(userid)  //Create in Redis
	createinPostgres(userid) //Create in Postgresql
}

// START User in Redis Functions -------->
func (u *UserRedis) createinRedis(userid int64) { //Create User In Redis
	u.UserID = userid
	u.IsLock = false
	u.TimeSpam = time.Now().Unix()
	u.Location = "main"

	userdata, _ := json.Marshal(u)
	rdb.Set(ctx, strconv.FormatInt(u.UserID, 10), userdata, 0).Result()
}

func GetUser(userid int64) UserRedis { //Return User From Redis
	user, err := rdb.Get(ctx, strconv.FormatInt(userid, 10)).Result()
	if err == redis.Nil {
		return UserRedis{}
	}
	userdata := UserRedis{}
	json.Unmarshal([]byte(user), &userdata)
	return userdata
}

func (u UserRedis) Exist() bool { //Check User Exist
	return u != UserRedis{}
}

func (u *UserRedis) UpdateTime() { //Update Spam Time in Redis
	u.TimeSpam = time.Now().UnixMilli()
}

func (u *UserRedis) update() { //Update User in Redis
	userdata, _ := json.Marshal(u)
	rdb.Set(ctx, strconv.FormatInt(u.UserID, 10), userdata, 0).Result()
}

func (u *UserRedis) lock() { //Lock User in Redis
	u.IsLock = true
	u.update()
}

func (u *UserRedis) unlock() { //unLock User in Redis
	u.IsLock = false
	u.update()
}
func (u *UserRedis) ChangeLocation(loc string) {
	u.Location = loc
	u.update()
}

// -----> END User Redis Functions

// Start User Postgresql Functions -------->
func createinPostgres(userid int64) { //Create User in Postgresql
	user := User{UserID: userid}
	if result := DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
}

func GetUserFromDB(userid int64) User {
	user := User{}
	if result := DB.First(&user, userid); result.Error != nil {
		fmt.Println(result.Error)
	}
	return user
}
func (u *User) UpdateWalletAddress(wallet string) {
	u.Wallet = wallet
	DB.Save(&u)
}

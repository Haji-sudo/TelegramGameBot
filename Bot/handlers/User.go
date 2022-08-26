package handlers

import (
	"dogegambling/config"
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

func CreateUserInAllDB(userid int64) { //Create User in Both DB
	CreateInRedis(userid)    //Create in Redis
	CreateInPostgres(userid) //Create in Postgresql
}

// START User in Redis Functions -------->
func CreateInRedis(userid int64) { //Create User In Redis
	user := UserRedis{}
	user.UserID = userid
	user.IsLock = false
	user.Location = "main"
	user.TimeSpam = 1661515274337

	userdata, _ := json.Marshal(user)
	rdb.Set(ctx, strconv.FormatInt(user.UserID, 10), userdata, 0).Result()
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
func UserExist(userid int64) bool { //Check User Exist Without dependency
	_, err := rdb.Get(ctx, strconv.FormatInt(userid, 10)).Result()

	return err != redis.Nil
}
func (u UserRedis) Exist() bool { //Check User Exist
	return u != UserRedis{}
}

func (u *UserRedis) UpdateTime() { //Update Spam Time in Redis
	u.TimeSpam = time.Now().UnixMilli()
	u.update()
}

func (u *UserRedis) NotSpam() bool {
	return time.Now().UnixMilli()-u.TimeSpam > 800 //if spam return True
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
func CreateInPostgres(userid int64) { //Create User in Postgresql
	user := User{UserID: userid}
	if result := DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
}
func UserExistInDB(userid int64) bool {
	return GetUserFromDB(userid) != User{}
}
func GetUserFromDB(userid int64) User {
	user := User{}

	result := DB.First(&user, userid)
	if result.RowsAffected <= 0 {
		return user
	}

	return user
}
func (u *User) UpdateWalletAddress(wallet string) {
	u.Wallet = wallet
	DB.Save(&u)
}
func (u *User) AddReferral() {
	u.Referrals += 1
	u.Balance += float32(config.Gift)
	DB.Save(&u)
}

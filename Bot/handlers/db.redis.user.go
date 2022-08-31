package handlers

import (
	"encoding/json"
	"strconv"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx context.Context
)

func (h Handler) UserInit() {
	rdb, ctx,
		DB = h.RDB, //init GormDB
		h.CTX, h.DB
}

func CreateUserInAllDB(userid int64) { //Create User in Both DB
	CreateInRedis(userid)    //Create in Redis
	CreateInPostgres(userid) //Create in Postgresql
}

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
	return time.Now().UnixMilli()-u.TimeSpam > 500 //if spam return True
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
func (u *UserRedis) ChangeLocation(loc string) { //Change User State
	u.Location = loc
	u.update()
}
func (u *UserRedis) SetBetAmount(amount float32) { //Set Amount for playing game
	u.AmountofBet = amount
	u.update()
}
func (u *UserRedis) SetGuessNumber(guess int, number int) { // set guess for Dice Game
	if guess == 1 {
		u.Dice.Guess1 = number
		u.update()
	} else {
		u.Dice.Guess2 = number
		u.update()
	}
}
func (u *UserRedis) SetWithdrawAmount(amount float32) { //set Amount for Withdraw
	u.AmountofBet = amount
	u.update()
}
func (u *UserRedis) GetWithdrawAmount() float32 { //get amount for withdraw
	return u.AmountofBet
}

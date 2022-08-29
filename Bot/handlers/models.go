package handlers

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Handler struct {
	RDB *redis.Client
	CTX context.Context
	DB  *gorm.DB
}

type UserRedis struct {
	UserID      int64   `json:"user_id"`
	IsLock      bool    `json:"lock"`
	TimeSpam    int64   `json:"time_spam"`
	Location    string  `json:"loc"`
	AmountofBet float32 `json:"bet_amount"`
}

type User struct {
	UserID         int64     `json:"user_id" gorm:"primaryKey"`
	Balance        float32   `json:"balance gorm:check:balance > 0"`
	Wallet         string    `json:"wallet_address"`
	DepositAddress string    `json:"deposit_address"`
	Referrals      uint      `json:"count_referrals"`
	Warn           byte      `json:"warn"`
	Ban            bool      `json:"isban"`
	CreatedAt      time.Time `json:"CreatedAt"`
	UpdatedAt      int64     `json:"last_update" gorm:"autoUpdateTime:milli"`
}
type Payment struct {
	PID       int32   `json:"pyment_id" gorm:"primaryKey"`
	UserRefer int64   `json:"user_id"`
	User      User    `gorm:"foreignKey:UserRefer"`
	TxID      string  `json:"tx_id"`
	Type      bool    `json:"type"` //True : Deposit | False : Withdraw
	Amount    float32 `json:"amount gorm:check:amount > 0"`
	Date      int64   `json:"time"`
}

type Config struct {
	Redis struct {
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		Server string `yaml:"server"`
		Port   string `yaml:"port"`
		DB     string `yaml:"db"`
	} `yaml:"redisdb"`
	Postgresql struct {
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		Server string `yaml:"server"`
		Port   string `yaml:"port"`
		DB     string `yaml:"db"`
	} `yaml:"posgresql"`
	BlockIO struct {
		Token   string `yaml:"token"`
		Pin     string `yaml:"pin"`
		Webhook string `yaml:"webhookurl"`
	} `yaml:"blockio"`
}

package handlers

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type Handler struct {
	RDB *redis.Client
	CTX context.Context
	DB  *gorm.DB
	Bot *telebot.Bot
}
type DiceGame struct {
	Guess1 int `json:"dice_guess_1"`
	Guess2 int `json:"dice_guess_2"`
}
type UserRedis struct {
	UserID      int64   `json:"user_id"`
	IsLock      bool    `json:"lock"`
	TimeSpam    int64   `json:"time_spam"`
	Location    string  `json:"loc"`
	AmountofBet float32 `json:"bet_amount"`
	Dice        DiceGame
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
	PID       int       `json:"pyment_id" gorm:"primaryKey;autoIncrement"`
	UserRefer int64     `json:"user_id"`
	User      User      `gorm:"foreignKey:UserRefer"`
	TxID      string    `json:"tx_id"`
	Type      bool      `json:"type"` //True : Deposit | False : Withdraw
	Amount    float32   `json:"amount gorm:check:amount > 0"`
	Date      time.Time `json:"time"`
	Status    string    `json:"status"`
}
type Bet struct {
	BID       int       `json:"bet_id" gorm:"primaryKey;autoIncrement"`
	UserRefer int64     `json:"user_id"`
	User      User      `gorm:"foreignKey:UserRefer"`
	Type      string    `json:"type"`
	Amount    float32   `json:"amount"`
	Date      time.Time `json:"time"`
	Result    string    `json:"result"`
}

type Config struct {
	Bot struct {
		Token                string `yaml:"token"`
		Username             string `yaml:"username"`
		Gift                 string `yaml:"gift"`
		WithdrawChannelID    string `yaml:"withdrawchannelID"`
		TransactionChannelID string `yaml:"transactionchannelID"`
		GamesChannelID       string `yaml:"gameschannelId"`
		Admins               string `yaml:"admins"`
	}
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

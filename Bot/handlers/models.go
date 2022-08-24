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
	UserID   int64  `json:"user_id"`
	IsLock   bool   `json:"lock"`
	TimeSpam int64  `json:"time_spam"`
	Warn     byte   `json:"warn"`
	Location string `json:"loc"`
}

type User struct {
	UserID        int64   `json:"user_id" gorm:"primaryKey"`
	Balance       float32 `json:"balance"`
	Wallet        string  `json:"wallet_address"`
	DepositAddres string  `json:"deposit_address"`
	Payment       Payment `json:"payments" gorm:"references:UserID"`
}
type Payment struct {
	PID    int32 `json:"pyment_id" gorm:"primaryKey"`
	UserID int64
	TxID   string    `json:"tx_id"`
	Type   string    `json:"type"`
	Amount float32   `json:"amount"`
	Date   time.Time `json:"time"`
}

package handlers

import (
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Handler struct {
	RDB *redis.Client
	CTX context.Context
	DB  *gorm.DB
}

type UserRedis struct {
	UserID   int64 `json:"user_id"`
	IsLock   bool  `json:"lock"`
	TimeSpam int64 `json:"time_spam"`
	Warn     byte  `json:"warn"`
}

type User struct {
	UserID  int64   `json:"user_id" gorm:"primaryKey"`
	Balance float32 `json:"balance"`
	Wallet  string  `json:"wallet_address"`
}

package main

import (
	"context"
	"dogegambling/config"
	p "dogegambling/config/DataBase/Postgresql"
	r "dogegambling/config/DataBase/Redis"
	"dogegambling/handlers"
	"fmt"
)

func NewHandler() handlers.Handler {
	db := handlers.Handler{
		RDB: r.InitRedisdb(),
		CTX: context.Background(),
		DB:  p.InitPostgredb(),
	}
	return db

}
func main() {

	handler := NewHandler()
	handler.Init()

	fmt.Println("Bot Started ...")
	config.Bot.Start()

}

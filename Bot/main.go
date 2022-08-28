package main

import (
	"context"
	p "dogegambling/DataBase/Postgresql"
	r "dogegambling/DataBase/Redis"
	gateway "dogegambling/Gateway"
	"dogegambling/config"
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
	gateway.Init("", "", "")
	fmt.Println("Bot Started ...")
	config.Bot.Start()

}

package main

import (
	"dogegambling/config"
	psdb "dogegambling/config/DataBase/Postgres"
	rddb "dogegambling/config/DataBase/Redis"
	"dogegambling/handlers"
	"fmt"
)

func NewHandler() handlers.Handler {
	DB := psdb.InitPostgredb()
	rdb, ctx := rddb.InitRedisdb()
	return handlers.Handler{RDB: rdb, CTX: ctx, DB: DB}

}
func main() {

	handler := NewHandler()
	handler.Init()

	fmt.Println("Bot Started ...")
	config.Bot.Start()

}

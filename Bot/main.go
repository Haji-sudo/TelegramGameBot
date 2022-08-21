package main

import (
	"dogegambling/config"
	"dogegambling/handlers"
	"fmt"
)

func main() {

	handlers.Init()

	fmt.Println("Bot Started ...")
	config.Bot.Start()

}

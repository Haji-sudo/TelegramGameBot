package handlers

import (
	c "dogegambling/config"

	"gopkg.in/telebot.v3"
)

func Init() {
	MenuInint()

	c.Bot.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send("Hello!", MainMenu)
	})

	c.Bot.Handle(&BtnGames, func(ctx telebot.Context) error {
		return ctx.Send("GameBoard", GameMenu)
	})

	c.Bot.Handle(&BtnHome, func(ctx telebot.Context) error {
		return ctx.Send("Home", MainMenu)
	})
}

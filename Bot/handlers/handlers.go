package handlers

import (
	c "dogegambling/config"
	"strconv"

	b "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func Init() {
	MenuInint()
	Admin := c.Bot.Group()
	Admin.Use(middleware.Whitelist(c.Admins...))

	c.Bot.Handle("/start", func(ctx b.Context) error {
		link := "tg://user?id=" + strconv.FormatInt(ctx.Chat().ID, 10)
		return ctx.Send(START(ctx.Chat().FirstName, link), MainMenu, b.ModeMarkdown)
	})

	c.Bot.Handle(&BtnGames, func(ctx b.Context) error {
		return ctx.Send("GameBoard", GameMenu)
	})

	c.Bot.Handle(&BtnDice, func(ctx b.Context) error {

		return ctx.Send("hi", DiceMenu)
	})

	c.Bot.Handle(&BtnHome, func(ctx b.Context) error {
		return ctx.Send("Home", MainMenu)
	})

}

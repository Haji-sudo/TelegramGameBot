package config

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

var Pref = tele.Settings{
	Token:  "951774571:AAEwmnWr3jXpNRULeBC-oOjIGjosIdgVqK0",
	Poller: &tele.LongPoller{Timeout: 10 * time.Second},
}

var (
	Bot, _      = tele.NewBot(Pref)
	Admins      = []int64{78246181}
	Gift        = 0.5
	BotUsername = "GetInfocryptoBot"
)

package config

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

var Pref = tele.Settings{
	Token:  "1653487808:AAFmjN6wIFHBf1KLkfx0-5uLZgEfUieRVIg",
	Poller: &tele.LongPoller{Timeout: 10 * time.Second},
}

var (
	Bot, _ = tele.NewBot(Pref)
	Admins = []int64{78246181}
	Gift   = 0.5
)

package config

import (
	"strconv"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	Token                string  = ""
	Admins                       = []int64{}
	Gift                 float32 = 0.5
	BotUsername          string  = "GetInfocryptoBot"
	WithdrawChannelID    int64   = -100
	TransactionChannelID int64   = -100
)

func InitBot(token, username, gift, with_ch, tx_ch, admins string) *tele.Bot {
	Token = token
	BotUsername = username
	n, _ := strconv.ParseFloat(gift, 32)
	Gift = float32(n)

	n2, _ := strconv.ParseInt(with_ch, 10, 64)
	WithdrawChannelID = n2

	n2, _ = strconv.ParseInt(tx_ch, 10, 64)
	TransactionChannelID = n2

	ad := strings.Split(admins, ",")
	for _, a := range ad {
		admin, _ := strconv.ParseInt(a, 10, 64)
		Admins = append(Admins, admin)
	}

	var Pref = tele.Settings{
		Token:  Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	Bot, _ := tele.NewBot(Pref)
	return Bot
}

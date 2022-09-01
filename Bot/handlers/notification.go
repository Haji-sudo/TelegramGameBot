package handlers

import b "gopkg.in/telebot.v3"

func SendToUser(Bot *b.Bot, userid int64, text string) {
	us := b.ChatID(userid)
	Bot.Send(us, text, b.ModeMarkdown)
}
func SendToChannel(Bot *b.Bot, channelid int64, text string) {
	ch := b.ChatID(channelid)
	Bot.Send(ch, text, b.ModeMarkdown)
}

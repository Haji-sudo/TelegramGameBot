package handlers

import (
	c "dogegambling/config"
	"strconv"

	b "gopkg.in/telebot.v3"
)

func HandelWithdraw(ctx b.Context, user *UserRedis, Bot *b.Bot) {
	input := ctx.Text()
	if input == BtnHome.Text {
		ctx.Send("Home", MainMenu)
		user.ChangeLocation(Main)
		return
	}
	if user.Location == Withdraw1 {
		amount, err := strconv.ParseFloat(input, 64)
		if err != nil {
			ctx.Send("Amount Is not correct")
			return
		} else if balance := UserBalance(ctx.Chat().ID); balance < float32(amount) {
			ctx.Send(BalanceNotEnough(balance))
			return
		}
		user.SetWithdrawAmount(float32(amount))
		user.ChangeLocation(Withdraw2)
		ctx.Send(WithdrawConfirm(float32(amount), ctx.Chat().ID), ConfirmBetMenu, b.ModeMarkdown)
		return
	} else if user.Location == Withdraw2 {
		if input != BtnConfirm.Text {
			ctx.Send("Please Confirm Or Back To Menu")
			return
		}

		pid := SubmitWithdraw(ctx.Chat().ID, user.GetWithdrawAmount())
		ch := b.ChatID(c.WithdrawChannelID)
		Bot.Send(ch, WithdrawText(ctx.Chat().FirstName, ctx.Chat().ID, user.GetWithdrawAmount()), b.ModeMarkdown, WithdrawButton(pid))
		user.ChangeLocation(Main)
		ctx.Send("Withdrawal request registered \n You can see the withdrawal status in Account 👤 -> 📉 Withdraw History", MainMenu)
	}
}

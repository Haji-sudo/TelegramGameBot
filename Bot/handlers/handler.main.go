package handlers

import (
	c "dogegambling/config"
	"fmt"

	b "gopkg.in/telebot.v3"
)

func HandelMain(ctx b.Context, user *UserRedis) {
	input := ctx.Text()
	UserID := ctx.Chat().ID
	if input == BtnGames.Text {
		user.ChangeLocation(Games)
		ctx.Send(GameBoard(), GameMenu)
		return

	} else if input == BtnReferrals.Text {
		link := fmt.Sprintf("t.me/%v?start=%v", c.BotUsername, UserID)
		ctx.Send(link)
		return
	} else if input == BtnAccount.Text {

		user.ChangeLocation(Account1)
		ctx.Send(ACCOUNT(ctx.Chat().FirstName, UserID), AccountMenu, b.ModeMarkdown)
		return
	} else if input == BtnWithdraw.Text {
		userdata := GetUserFromDB(UserID)
		if userdata.Wallet == "" {
			ctx.Send("Add Your Wallet Address From Account 👤")
		} else {
			user.ChangeLocation(Withdraw1)
			ctx.Send(Balance(userdata.Balance), b.ModeMarkdown)
			ctx.Send("Enter amount You want withdraw")
		}
		return
	}
}

func CopyedString(str string) string {
	return fmt.Sprintf("`%v`", str)
}

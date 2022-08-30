package handlers

import (
	c "dogegambling/config"
	"fmt"
	"strconv"

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
		userdata := GetUserFromDB(UserID)
		link := "tg://user?id=" + strconv.FormatInt(UserID, 10)
		ctx.Send(ACCOUNT(ctx.Chat().FirstName, link, userdata.Balance, userdata.Referrals, userdata.Warn, CopyedString(userdata.Wallet)), b.ModeMarkdown)
		return
	}
}

func CopyedString(str string) string {
	return fmt.Sprintf("`%v`", str)
}

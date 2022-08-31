package handlers

import (
	gateway "dogegambling/Gateway"
	"fmt"
	"strconv"

	b "gopkg.in/telebot.v3"
)

func HandelAccount(ctx b.Context, user *UserRedis) {
	input := ctx.Text()
	if input == BtnHome.Text {
		ctx.Send("Home", MainMenu)
		user.ChangeLocation(Main)
		return
	} else if input == BtnAccount.Text {
		userdata := GetUserFromDB(ctx.Chat().ID)
		link := "tg://user?id=" + strconv.FormatInt(ctx.Chat().ID, 10)
		user.ChangeLocation(Account1)
		ctx.Send(ACCOUNT(ctx.Chat().FirstName, link, userdata.Balance, userdata.Referrals, userdata.Warn, CopyedString(userdata.Wallet)), AccountMenu, b.ModeMarkdown)
	}
	if user.Location == Account1 {
		if input == BtnDepositHistory.Text {
			deposits := GetDepositHistory(ctx.Chat().ID)
			res := "Deposit's :\n 〰️〰️〰️〰️〰️〰️〰️〰️〰️〰️"
			for _, d := range deposits {
				res += fmt.Sprintf(`
				Amount : %v
				TXID : %v
				Date : %v
				〰️〰️〰️〰️〰️〰️〰️〰️〰️〰️`, d.Amount, CopyedString(d.TxID), d.Date.Format("2006-01-02 15:04"))
			}
			ctx.Send(res)
			return
		} else if input == BtnWithdrawHistory.Text {
			withdraws := GetWithdrawHistory(ctx.Chat().ID)
			res := "Withdraw's :\n 〰️〰️〰️〰️〰️〰️〰️〰️〰️〰️"
			for _, d := range withdraws {
				res += fmt.Sprintf(`
				Amount : %v
				TXID : %v
				Date : %v
				〰️〰️〰️〰️〰️〰️〰️〰️〰️〰️`, d.Amount, CopyedString(d.TxID), d.Date.Format("2006-01-02 15:04"))
			}
			ctx.Send(res)
			return
		} else if input == BtnGamesHistory.Text {
			gameshistory := GetGamesHistory(ctx.Chat().ID)
			res := "Games :\n 〰️〰️〰️〰️〰️〰️〰️〰️〰️〰️"
			for _, v := range gameshistory {
				res += fmt.Sprintf(`
				Game : %v
				Amount : %v
				Date : %v
				Result : %v
				〰️〰️〰️〰️〰️〰️〰️〰️〰️〰️`, v.Type, v.Amount, v.Date.Format("2006-01-02 15:04"), v.Result)
			}
			ctx.Send(res)
			return
		} else if input == BtnChangeAddress.Text {
			user.ChangeLocation(Account2)
			ctx.Send("Send your wallet address", AccountMenu2)
			return
		}
	} else if user.Location == Account2 {
		if gateway.ValidateAddress(input) {
			userdata := GetUserFromDB(ctx.Chat().ID)
			if input == userdata.Wallet {
				user.ChangeLocation(Account1)
				ctx.Send("The entered address is similar to the previous address", AccountMenu, b.ModeMarkdown)
				return
			} else if input == userdata.DepositAddress {
				ctx.Send("The entered address is similar to your deposit address, you cannot use this address for withdrawal \n Enter a valid Doge Address")
				return
			}
			userdata.UpdateWalletAddress(input)
			user.ChangeLocation(Account1)
			ctx.Send("Your wallet has been updated", AccountMenu, b.ModeMarkdown)
			return
		} else {
			ctx.Send("The entered address is wrong \n Enter a valid Doge Address")
			return
		}
	}
}

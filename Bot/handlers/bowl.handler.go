package handlers

import (
	"fmt"
	"strconv"
	"time"

	b "gopkg.in/telebot.v3"
)

func HandelBowl(ctx b.Context, user *UserRedis) {
	input := ctx.Text()
	if input == BtnGames.Text {
		ctx.Send("GameBoard", GameMenu, b.ModeMarkdown)
		user.ChangeLocation(Games)
		return
	} else if input == BtnHome.Text {
		ctx.Send("Home", MainMenu)
		user.ChangeLocation(Main)
		return
	} else if user.Location == Bowl1 { // Enter Bet Number State
		bet_amount, err := strconv.ParseFloat(input, 64)
		if err != nil {
			ctx.Send("Amount Is not correct")
			return
		} else if bet_amount < Minbet || bet_amount > Maxbet {
			ctx.Send(fmt.Sprintf("Amount must be between %v and %v", Minbet, Maxbet))
			return
		} else if balance := UserBalance(ctx.Chat().ID); balance < float32(bet_amount) {
			ctx.Send(BalanceNotEnough(balance))
			return
		}
		user.SetBetAmount(float32(bet_amount))
		user.ChangeLocation(Bowl2)
		ctx.Send(BowlText2(float32(bet_amount)), ConfirmBetMenu, b.ModeMarkdown)
		return
	} else if user.Location == Bowl2 {
		if input != BtnConfirm.Text {
			ctx.Send("Please Confirm Or Back To Menu")
			return
		}
		ctx.Send("Throwing the Bowling 🎳", b.RemoveKeyboard)
		bowl, err := b.Bowl.Send(ctx.Bot(), ctx.Recipient(), nil)
		if err != nil {
			ctx.Send("Something Wrong /start Again")
			return
		}
		ConfirmBet(ctx.Chat().ID, user.AmountofBet)
		time.Sleep(time.Second * 3)
		switch bowl.Dice.Value {
		case 6:
			BetWin(ctx.Chat().ID, user.AmountofBet*2)
			ctx.Send(WonText(2))
		case 5:
			BetWin(ctx.Chat().ID, user.AmountofBet*1.3)
			ctx.Send(WonText(1.3))
		case 4:
			BetWin(ctx.Chat().ID, user.AmountofBet*0.9)
			ctx.Send(WonText(0.9))
		case 3:
			BetWin(ctx.Chat().ID, user.AmountofBet*0.6)
			ctx.Send(WonText(0.6))
		case 2:
			BetWin(ctx.Chat().ID, user.AmountofBet*0.3)
			ctx.Send(WonText(0.3))
		default:
			ctx.Send("You Lost")
		}
		user.ChangeLocation(Games)
		ctx.Send("Game Menu", GameMenu)

	}
}

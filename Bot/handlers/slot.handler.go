package handlers

import (
	"fmt"
	"strconv"
	"time"

	b "gopkg.in/telebot.v3"
)

func HandelSlot(ctx b.Context, user *UserRedis) {
	input := ctx.Text()
	if input == BtnGames.Text {
		ctx.Send("GameBoard", GameMenu, b.ModeMarkdown)
		user.ChangeLocation(Games)
		return
	} else if input == BtnHome.Text {
		ctx.Send("Home", MainMenu)
		user.ChangeLocation(Main)
		return
	} else if user.Location == Slot1 { // Enter Bet Number State
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
		user.ChangeLocation(Slot2)
		ctx.Send(SlotText2(float32(bet_amount)), ConfirmBetMenu, b.ModeMarkdown)
		return
	} else if user.Location == Slot2 {
		if input != BtnConfirm.Text {
			ctx.Send("Please Confirm Or Back To Menu")
			return
		}
		ctx.Send("Spin the Slot Machine 🎰", b.RemoveKeyboard)
		slot, err := b.Slot.Send(ctx.Bot(), ctx.Recipient(), nil)
		if err != nil {
			ctx.Send("Something Wrong /start Again")
			return
		}
		ConfirmBet(ctx.Chat().ID, user.AmountofBet)
		time.Sleep(time.Second * 3)

		switch slot.Dice.Value {
		case 1, 22, 43, 64: //Win
			BetWin(ctx.Chat().ID, user.AmountofBet*2)
			ctx.Send(WonText(2))
		case 7, 8, 10, 12, 14, 15, 19, 20, 25, 28, 29, 31, 34, 36, 37, 40, 45, 46, 50, 51, 55, 53, 57, 58: //Lose
			ctx.Send("You Lost")
		default:
			winrnd := GetRandomWinNumber()
			BetWin(ctx.Chat().ID, user.AmountofBet*winrnd)
			ctx.Send(WonText(winrnd))

		}
		user.ChangeLocation(Games)
		ctx.Send("Game Menu", GameMenu)

	}
}

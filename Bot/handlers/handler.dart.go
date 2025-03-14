package handlers

import (
	"fmt"
	"strconv"
	"time"

	b "gopkg.in/telebot.v3"
)

func HandelDart(ctx b.Context, user *UserRedis) {
	input := ctx.Text()
	if input == BtnGames.Text {
		ctx.Send(GameBoard(), GameMenu, b.ModeMarkdown)
		user.ChangeLocation(Games)
		return
	} else if input == BtnHome.Text {
		ctx.Send("Home", MainMenu)
		user.ChangeLocation(Main)
		return
	} else if user.Location == Dart1 {
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
		user.ChangeLocation(Dart2)
		ctx.Send(DartText2(float32(bet_amount)), ConfirmBetMenu, b.ModeMarkdown)
		return
	} else if user.Location == Dart2 {
		if input != BtnConfirm.Text {
			ctx.Send("Please Confirm Or Back To Menu")
			return
		}
		ctx.Send("Throwing the Dart 🎯", b.RemoveKeyboard)
		dart, err := b.Dart.Send(ctx.Bot(), ctx.Recipient(), nil)
		if err != nil {
			ctx.Send("Something Wrong /start Again")
			return
		}
		ConfirmBet(ctx.Chat().ID, user.AmountofBet)
		time.Sleep(time.Second * 3)
		switch dart.Dice.Value {
		case 6:
			BetWin(ctx.Chat().ID, user.AmountofBet*2)
			ctx.Send(WinText(2))
			SaveGameHistroy(ctx.Chat().ID, Dart, user.AmountofBet, `Win 2x`)
		case 5:
			BetWin(ctx.Chat().ID, user.AmountofBet*1.3)
			ctx.Send(WinText(1.3))
			SaveGameHistroy(ctx.Chat().ID, Dart, user.AmountofBet, `Win 1.3x`)
		case 4:
			BetWin(ctx.Chat().ID, user.AmountofBet*0.9)
			ctx.Send(WinText(0.9))
			SaveGameHistroy(ctx.Chat().ID, Dart, user.AmountofBet, `Lose 0.9x`)
		case 3:
			BetWin(ctx.Chat().ID, user.AmountofBet*0.6)
			SaveGameHistroy(ctx.Chat().ID, Dart, user.AmountofBet, `Lose 0.6x`)
			ctx.Send(WinText(0.6))
		case 2:
			BetWin(ctx.Chat().ID, user.AmountofBet*0.3)
			SaveGameHistroy(ctx.Chat().ID, Dart, user.AmountofBet, `Lose 0.3x`)
			ctx.Send(WinText(0.3))
		default:
			ctx.Send("You Lost")
			SaveGameHistroy(ctx.Chat().ID, Dart, user.AmountofBet, `Lose 0x`)
		}
		user.ChangeLocation(Games)
		ctx.Send(GameBoard(), GameMenu)

	}
}

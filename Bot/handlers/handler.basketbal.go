package handlers

import (
	"fmt"
	"strconv"
	"time"

	b "gopkg.in/telebot.v3"
)

func HandelBasket(ctx b.Context, user *UserRedis) {
	input := ctx.Text()
	if input == BtnGames.Text {
		ctx.Send(GameBoard(), GameMenu, b.ModeMarkdown)
		user.ChangeLocation(Games)
		return
	} else if input == BtnHome.Text {
		ctx.Send("Home", MainMenu)
		user.ChangeLocation(Main)
		return
	} else if user.Location == Basket1 { // Enter Bet Number State
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
		user.ChangeLocation(Basket2)
		ctx.Send(BasketText2(float32(bet_amount)), ConfirmBetMenu, b.ModeMarkdown)
		return
	} else if user.Location == Basket2 {
		if input != BtnConfirm.Text {
			ctx.Send("Please Confirm Or Back To Menu")
			return
		}
		ctx.Send("Throwing the Basketball 🏀", b.RemoveKeyboard)
		basket, err := b.Ball.Send(ctx.Bot(), ctx.Recipient(), nil)
		if err != nil {
			ctx.Send("Something Wrong /start Again")
			return
		}
		ConfirmBet(ctx.Chat().ID, user.AmountofBet)
		time.Sleep(time.Second * 3)

		switch basket.Dice.Value {
		case 4, 5: //Win
			SaveGameHistroy(ctx.Chat().ID, Basket, user.AmountofBet, `Win 1.8x`)
			BetWin(ctx.Chat().ID, user.AmountofBet*1.8)
			ctx.Send(WinText(1.8))
			SaveGameHistroy(ctx.Chat().ID, Basket, user.AmountofBet, `Win 1.8x`)
		case 1, 2, 3: //Lose
			ctx.Send("You Lost")
			SaveGameHistroy(ctx.Chat().ID, Basket, user.AmountofBet, `Lose 0x`)
		default:
			ctx.Send("Return")

		}
		user.ChangeLocation(Games)
		ctx.Send(GameBoard(), GameMenu)

	}
}

package handlers

import (
	"fmt"
	"strconv"
	"time"

	b "gopkg.in/telebot.v3"
)

func HandelDice(ctx b.Context, user *UserRedis) {
	input := ctx.Text()
	if input == BtnGames.Text {
		ctx.Send("GameBoard", GameMenu)
		user.ChangeLocation(Games)
		return
	} else if input == BtnHome.Text {
		ctx.Send("Home", MainMenu)
		user.ChangeLocation(Main)
		return
	} else if user.Location == Dice1 { // Enter Bet Number State
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
		user.ChangeLocation(Dice2)
		ctx.Send(Dice2Detaile())
		return
	} else if user.Location == Dice2 {
		guess, err := strconv.Atoi(input)
		if err != nil {
			ctx.Send("Number Is not correct")
			return
		} else if guess <= 0 || guess >= 7 {
			ctx.Send("Number Must be between 1 and 6")
			return
		}
		user.SetGuessNumber(1, guess)
		user.ChangeLocation(Dice3)
		ctx.Send(Dice3Detaile())
		return
	} else if user.Location == Dice3 {
		guess, err := strconv.Atoi(input)
		if err != nil {
			ctx.Send("Number Is not correct")
			return
		} else if guess <= 0 || guess >= 7 {
			ctx.Send("Number Must be between 1 and 6")
			return
		}
		user.SetGuessNumber(2, guess)
		user.ChangeLocation(Dice4)
		ctx.Send(DiceConfirmBet(user.Dice.Guess1, user.Dice.Guess2, user.AmountofBet), ConfirmBetMenu, b.ModeMarkdown)
		return
	} else if user.Location == Dice4 {
		if input != BtnConfirm.Text {
			ctx.Send("Please Confirm Or Back To Menu")
			return
		}
		ctx.Send("Throwing the Dice 🎲", b.RemoveKeyboard)
		cube, err := b.Cube.Send(ctx.Bot(), ctx.Recipient(), nil)
		if err != nil {
			ctx.Send("Something Wrong /start Again")
			return
		}
		ConfirmBet(ctx.Chat().ID, user.AmountofBet)
		time.Sleep(time.Second * 3)
		if user.Dice.Guess1 == user.Dice.Guess2 {
			if cube.Dice.Value == user.Dice.Guess1 {
				ctx.Send("You Win 4x")
				BetWin(ctx.Chat().ID, user.AmountofBet*4)
			} else {
				ctx.Send("You Lost")
			}
		} else {
			if cube.Dice.Value == user.Dice.Guess1 || cube.Dice.Value == user.Dice.Guess2 {
				ctx.Send("You Win 2x")
				BetWin(ctx.Chat().ID, user.AmountofBet*2)
			} else {
				ctx.Send("You Lost")
			}
		}
		user.ChangeLocation(Games)
		ctx.Send("Game Menu", GameMenu)
		return
	}
}

package handlers

import (
	"strconv"
	"strings"

	b "gopkg.in/telebot.v3"
)

func HandelGameBoard(ctx b.Context, user UserRedis) {
	input := ctx.Text()
	UserID := ctx.Chat().ID
	if input == BtnDice.Text {
		user.ChangeLocation("dice1")
		userdb := GetUserFromDB(UserID)

		ctx.Send(DiceDetails(userdb.Balance, 1, 200), DiceMenu, b.ModeMarkdown)
	} else if input == BtnBowling.Text {
		user.ChangeLocation("bowl1")
		ctx.Send("Bowling Menu")
	} else if input == BtnDart.Text {
		user.ChangeLocation("dart1")
		ctx.Send("Dart Menu")
	} else if input == BtnSlot.Text {
		user.ChangeLocation("slot1")
		ctx.Send("Slot Menu")
	} else {
		ctx.Send("Not Found")
	}
}
func HandelDice(ctx b.Context, user UserRedis) {
	input := ctx.Text()
	if input == BtnBack.Text {
		user.ChangeLocation("games")
		ctx.Send("GameBoard", GameMenu)
	}
	bet_amount, err := strconv.ParseFloat(input, 32)
	if err != nil {
		ctx.Send("Amount Is not correct")
	}
	user.SetBetAmount(float32(bet_amount))
	user.ChangeLocation("dice2")
}

func HandelBackBTN(location string) string {
	if strings.Contains(location, "dice") {
		location = strings.ReplaceAll(location, "dice", "")
		number, _ := strconv.Atoi(location)
		switch number {
		case 1:
			return "games"
		case 2:
			return "dice1"
		default:
			return "main"
		}

	}
	return "main"
}

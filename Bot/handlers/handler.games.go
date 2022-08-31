package handlers

import (
	b "gopkg.in/telebot.v3"
)

func HandelGameBoard(ctx b.Context, user *UserRedis) {
	input := ctx.Text()
	if input == BtnDice.Text {
		user.ChangeLocation(Dice1)
		ctx.Send(DiceDetails(ctx.Chat().ID), GameMenu2, b.ModeMarkdown)

	} else if input == BtnBowling.Text {
		user.ChangeLocation(Bowl1)
		ctx.Send(BowlText1(ctx.Chat().ID), GameMenu2, b.ModeMarkdown)

	} else if input == BtnDart.Text {
		user.ChangeLocation(Dart1)
		ctx.Send(DartText1(ctx.Chat().ID), GameMenu2, b.ModeMarkdown)

	} else if input == BtnSlot.Text {
		user.ChangeLocation(Slot1)
		ctx.Send(SlotText1(ctx.Chat().ID), GameMenu2, b.ModeMarkdown)

	} else if input == BtnBalance.Text {
		ctx.Send(Balance(UserBalance(ctx.Chat().ID)), b.ModeMarkdown)

	} else if input == BtnBasketball.Text {
		user.ChangeLocation(Basket1)
		ctx.Send(BasketText1(ctx.Chat().ID), GameMenu2, b.ModeMarkdown)

	} else {
		ctx.Send("Not Found")
	}
}

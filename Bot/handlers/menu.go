package handlers

import (
	"gopkg.in/telebot.v3"
)

// Main Keybords
var (
	MainMenu     = &telebot.ReplyMarkup{ResizeKeyboard: true}
	BtnGames     = MainMenu.Text("Games 🎮")
	BtnDeposit   = MainMenu.Text("Deposit 📥")
	BtnWithdraw  = MainMenu.Text("Withdraw 📤")
	BtnReferrals = MainMenu.Text("Referrals 👥")
	BtnFAQ       = MainMenu.Text("FAQ ❓")
	BtnAccount   = MainMenu.Text("Account 👤")
)

var (
	GameMenu = &telebot.ReplyMarkup{ResizeKeyboard: true}
	BtnDice  = GameMenu.Text("Dice 🎲")
	BtnHome  = GameMenu.Text("Home 🏠")
)
var (
	DiceMenu    = &telebot.ReplyMarkup{ResizeKeyboard: true}
	Btn1        = DiceMenu.Text("1")
	Btn2        = DiceMenu.Text("2")
	Btn3        = DiceMenu.Text("3")
	Btn4        = DiceMenu.Text("4")
	Btn5        = DiceMenu.Text("5")
	Btn6        = DiceMenu.Text("6")
	BtnNoneDice = DiceMenu.Text("🎲")
)

func MenuInint() {

	MainMenu.Reply(
		MainMenu.Row(BtnGames),
		MainMenu.Row(BtnDeposit, BtnWithdraw),
		MainMenu.Row(BtnReferrals, BtnFAQ),
		MainMenu.Row(BtnAccount),
	)

	GameMenu.Reply(
		GameMenu.Row(BtnDice),
		GameMenu.Row(BtnHome),
	)

	DiceMenu.Reply(
		DiceMenu.Row(Btn1, Btn2),
		DiceMenu.Row(Btn3, Btn4),
		DiceMenu.Row(Btn5, Btn6),
		DiceMenu.Row(BtnNoneDice),
	)
}

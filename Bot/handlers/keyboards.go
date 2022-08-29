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
	BtnHome      = MainMenu.Text("Home 🏠")
	BtnBack      = MainMenu.Text("🔙 Back")
)

var (
	GameMenu   = &telebot.ReplyMarkup{ResizeKeyboard: true}
	BtnDice    = GameMenu.Text("Dice 🎲")
	BtnBowling = GameMenu.Text("Bowling 🎳")
	BtnDart    = GameMenu.Text("Dart 🎯")
	BtnSlot    = GameMenu.Text("Slot 🎰")
)
var (
	DiceMenu  = &telebot.ReplyMarkup{ResizeKeyboard: true}
	Btn1      = DiceMenu.Text("1")
	Btn2      = DiceMenu.Text("2")
	Btn3      = DiceMenu.Text("3")
	Btn4      = DiceMenu.Text("4")
	Btn5      = DiceMenu.Text("5")
	Btn6      = DiceMenu.Text("6")
	DiceMenu2 = &telebot.ReplyMarkup{ResizeKeyboard: true}
)

func MenuInint() {

	MainMenu.Reply(
		MainMenu.Row(BtnGames),
		MainMenu.Row(BtnDeposit, BtnWithdraw),
		MainMenu.Row(BtnReferrals, BtnFAQ),
		MainMenu.Row(BtnAccount),
	)

	GameMenu.Reply(
		GameMenu.Row(BtnDice, BtnBowling),
		GameMenu.Row(BtnDart, BtnSlot),
		GameMenu.Row(BtnHome),
	)

	DiceMenu.Reply( //Dice Menu1
		DiceMenu.Row(BtnBack),
		DiceMenu.Row(BtnHome),
	)

}

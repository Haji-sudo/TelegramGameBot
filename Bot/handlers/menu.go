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
}

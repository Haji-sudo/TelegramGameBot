package handlers

import (
	"gopkg.in/telebot.v3"
)

// Main Keybords
var ()

var (
	BtnGames           = MainMenu.Text("Games 🎮")
	BtnDeposit         = MainMenu.Text("Deposit 📥")
	BtnWithdraw        = MainMenu.Text("Withdraw 📤")
	BtnReferrals       = MainMenu.Text("Referrals 👥")
	BtnFAQ             = MainMenu.Text("FAQ ❓")
	BtnAccount         = MainMenu.Text("Account 👤")
	BtnChangeAddress   = MainMenu.Text("🔁 Change Wallet Address")
	BtnDepositHistory  = MainMenu.Text("📈 Deposit History")
	BtnWithdrawHistory = MainMenu.Text("📉 Withdraw History")
	BtnGamesHistory    = MainMenu.Text("🧨 Games History")

	BtnHome       = MainMenu.Text("Home 🏠")
	BtnDice       = GameMenu.Text("Dice 🎲")
	BtnBowling    = GameMenu.Text("Bowling 🎳")
	BtnDart       = GameMenu.Text("Dart 🎯")
	BtnSlot       = GameMenu.Text("Slot 🎰")
	BtnBalance    = GameMenu.Text("💰 Balance")
	BtnConfirm    = GameMenu.Text("✅ Confirm")
	BtnBasketball = GameMenu.Text("Basketball 🏀")
)
var (
	MainMenu       = &telebot.ReplyMarkup{ResizeKeyboard: true}
	GameMenu       = &telebot.ReplyMarkup{ResizeKeyboard: true}
	GameMenu2      = &telebot.ReplyMarkup{ResizeKeyboard: true}
	ConfirmBetMenu = &telebot.ReplyMarkup{ResizeKeyboard: true}
	AccountMenu    = &telebot.ReplyMarkup{ResizeKeyboard: true}
	AccountMenu2   = &telebot.ReplyMarkup{ResizeKeyboard: true}
)

func MenuInint() {

	MainMenu.Reply(
		MainMenu.Row(BtnGames),
		MainMenu.Row(BtnDeposit, BtnWithdraw),
		MainMenu.Row(BtnReferrals, BtnFAQ),
		MainMenu.Row(BtnAccount),
	)
	AccountMenu.Reply(
		AccountMenu.Row(BtnChangeAddress),
		AccountMenu.Row(BtnDepositHistory, BtnWithdrawHistory),
		AccountMenu.Row(BtnGamesHistory),
		AccountMenu.Row(BtnHome),
	)
	GameMenu.Reply(
		GameMenu.Row(BtnDice, BtnBowling, BtnBasketball),
		GameMenu.Row(BtnDart, BtnSlot),
		GameMenu.Row(BtnBalance),
		GameMenu.Row(BtnHome),
	)

	GameMenu2.Reply(
		GameMenu2.Row(BtnHome, BtnGames),
	)
	ConfirmBetMenu.Reply(
		ConfirmBetMenu.Row(BtnConfirm),
		ConfirmBetMenu.Row(BtnHome, BtnGames),
	)
	AccountMenu2.Reply(
		AccountMenu2.Row(BtnHome, BtnAccount),
	)
}

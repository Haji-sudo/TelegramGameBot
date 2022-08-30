package handlers

import (
	"fmt"
)

func START(name string, link string) string {
	return fmt.Sprintf(`** Hello ** [%v](%v)

__Welcome to Doge Finance 🐶__
	
__You Can **PLAY** and **BET** here__
	
__Also You Can Invite Your Friends And Earn **DOGE Coin**__
	
**For More Information GoTo FAQ ❓ Menu**`, name, link)
}

func ACCOUNT(name string, link string, balance float32, referrals uint, warning byte, wallet string) string {
	return fmt.Sprintf(`ℹ️ **INFORMATION**

	👤 __Name__: **[%v](%v)**
	
	💰 __Balance__: **%vÐ**
	
	👥 __Total Referrals__: **%v**
	
	⛔ __Warning__: **%v**
	
	💳 **Wallet Address**: %v`, name, link, balance, referrals, warning, wallet)
}

func DEPOSIT(deposit string) string {
	return fmt.Sprintf(`ℹ️ **Deposit**
	
	💳 **Wallet Address**: %v`, deposit)
}

func FAQ() string {
	return `FAQ ❓ 
	Detailes`
}
func GameBoard() string {
	return "Choose The Game that You Want To Play🕹"
}
func BalanceNotEnough(balance float32) string {
	return fmt.Sprintf("❌ Your balance is not enough \n\n💰 Balance = %v ", balance)
}
func Balance(balance float32) string {
	return fmt.Sprintf("Your 💰 Balance: **%vÐ** ", balance)
}
func DiceDetails(userid int64) string {
	balance := UserBalance(userid)
	return fmt.Sprintf(`		**Dice**🎲		

	__If Choose 2 Numbers:__
	__Right Guess__ :** 2x 😍️ **
	__Wrong Guess__ :** 0x  🥺️ **
	
	__If Choose 1 Number:__
	__Right Guess__ :** 4x 😍️ **
	__Wrong Guess__ :** 0x   🥺️ **
	
	💰 __Balance__: **%vÐ**

	❗ __The Minimum Bet Amount Is__ **%vÐ**
	❗ __The Maximum Bet Amount Is__ **%vÐ**


	💸 __Send The Required Bet Amount__
	`, balance, Minbet, Maxbet)
}
func Dice2Detaile() string {
	return `Enter Guess 1
	Guess Number Must be between 1 and 6`
}
func Dice3Detaile() string {
	return `Enter Guess 2

	If you want choose 1 Number send same number`
}
func DiceConfirmBet(guess1 int, guess2 int, amount_bet float32) string {
	return fmt.Sprintf(`Your Guess 1 : Number (**%v**)
	Your Guess 2 : Number (**%v**)
	
	The amount of the bet : **%vÐ**

	✅Confirm Bet
	`, guess1, guess2, amount_bet)
}
func BowlText1(userid int64) string {
	balance := UserBalance(userid)
	return fmt.Sprintf(`** Bowling 🎳 **

	__Knock Down:__
	__Strike__ : ** 2x 😍️ **
	__5 Pins__ : ** 1.3x 🙂️ **
	__4 Pins__ : ** 0.9x 😟 **
	__3 Pins__ : ** 0.6x 😟 **
	__1 Pins__ : ** 0.3x 😟 **
	__0 Pins__ : ** 0x  🥺️ **
	
	💰 __Balance__: **%vÐ**
	❗ __The Minimum Bet Amount Is__ **%vÐ**
	❗ __The Maximum Bet Amount Is__ **%vÐ**

	💸 __Send The Required Bet Amount__
	`, balance, Minbet, Maxbet)
}
func BowlText2(amount_bet float32) string {
	return fmt.Sprintf(`** Bowling 🎳 **

	💸 __Bet Amount__ : **%vÐ**
	
	__ If Are You **Ready?**__
	
	✅Confirm Bet `, amount_bet)
}
func WonText(won_rate float32) string {
	return fmt.Sprintf(`You Won %vx`, won_rate)
}
func DartText1(userid int64) string {
	balance := UserBalance(userid)
	return fmt.Sprintf(`** Darts 🎯 **

	__Tatget__ : ** 2x 😍️ **
	__2nd Ring__ : ** 1.3x 🙂️ **
	__3nd Ring__ : ** 0.9x 😟 **
	__4nd Ring__ : ** 0.6x 😟 **
	__5nd Ring__ : ** 0.3x 😟 **
	__Out__ : ** 0x  🥺️ **
	
	💰 __Balance__: **%vÐ**

	❗ __The Minimum Bet Amount Is__ **%vÐ**
	❗ __The Maximum Bet Amount Is__ **%vÐ**

	💸 __Send The Required Bet Amount__
	`, balance, Minbet, Maxbet)
}
func DartText2(amount_bet float32) string {
	return fmt.Sprintf(`** Darts 🎯 **

	💸 __Bet Amount__ : **%vÐ**
	
	__ If Are You **Ready?**__
	
	✅Confirm Bet `, amount_bet)
}
func SlotText1(userid int64) string {
	balance := UserBalance(userid)
	return fmt.Sprintf(`** Slot Machine 🎰 **

	🟢 🟢 🟢 : ** 2x 😍️ **
	🟢 🔴 🟢 : ** 1-1.2x 🙂️ **
	🟢 🔴 🟡 : ** 0x  🥺️ **
	
	💰 __Balance__: **%vÐ**

	❗ __The Minimum Bet Amount Is__ **%vÐ**
	❗ __The Maximum Bet Amount Is__ **%vÐ**

	💸 __Send The Required Bet Amount__
	`, balance, Minbet, Maxbet)
}
func SlotText2(amount_bet float32) string {
	return fmt.Sprintf(`** Slot Machine 🎰 **

	💸 __Bet Amount__ : **%vÐ**
	
	__ If Are You **Ready?**__
	
	✅Confirm Bet `, amount_bet)
}
func BasketText1(userid int64) string {
	balance := UserBalance(userid)
	return fmt.Sprintf(`** Basketball 🏀 **

	✅ : ** 1.8x 😍️ **
	❌ : ** 0x  🥺️ **
	
	▫ __Balance__: **%vÐ**
	❗ __The Minimum Bet Amount Is__ **%vÐ**
	❗ __The Maximum Bet Amount Is__ **%vÐ**
	💸 __Send The Required Bet Amount__`, balance, Minbet, Maxbet)
}
func BasketText2(amount_bet float32) string {
	return fmt.Sprintf(`** Basketball 🏀 **

	💸 __Bet Amount__ : **%vÐ**
	
	__ If Are You **Ready?**__
	
	✅Confirm Bet `, amount_bet)
}

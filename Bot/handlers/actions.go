package handlers

import (
	"fmt"
	"strconv"
)

func START(name string, link string) string {
	return fmt.Sprintf(`** Hello ** [%v](%v)

__Welcome to Doge Finance 🐶__
	
__You Can **PLAY** and **BET** here__
	
__Also You Can Invite Your Friends And Earn **DOGE Coin**__
	
**For More Information GoTo FAQ ❓ Menu**`, name, link)
}
func CreateLinkWithUserID(userid int64) string {
	return "tg://user?id=" + strconv.FormatInt(userid, 10)
}
func ACCOUNT(name string, userid int64) string {
	user := GetUserFromDB(userid)

	return fmt.Sprintf(`ℹ️ **INFORMATION**

	👤 __Name__: **[%v](%v)**
	
	💰 __Balance__: **%vÐ**
	
	👥 __Total Referrals__: **%v**
	
	⛔ __Warning__: **%v**
	
	💳 **Wallet Address**: %v`, name, CreateLinkWithUserID(userid), user.Balance, user.Referrals, user.Warn, CopyedString(user.Wallet))
}
func WithdrawText(name string, userid int64, amount float32) string {
	user := GetUserFromDB(userid)

	return fmt.Sprintf(`ℹ️ **INFORMATION**

	👤 __Name__: **[%v](%v)**
	
	💰 __Balance__: **%vÐ**
	
	👥 __Total Referrals__: **%v**
	
	⛔ __Warning__: **%v**
	
	💳 **Wallet Address**: %v
	
	Amount requested : %v`, name, CreateLinkWithUserID(userid), user.Balance, user.Referrals, user.Warn, CopyedString(user.Wallet), amount)
}
func ConfirmWithdrawTextChannel(userid int64, amount float32, txid string) string {
	return fmt.Sprintf(`**    ✅ Confirmed**

	👤**[User](%v)**
	
	Amount requested : %v
	
	TXID : https://blockchair.com/dogecoin/transaction/%v`, CreateLinkWithUserID(userid), amount, txid)
}
func RejectWithdrawTextChannel(userid int64, amount float32, pid int) string {
	return fmt.Sprintf(`**    ❌ Rejected**

	👤**[User](%v)**
	
	Amount requested : %v
	
	Request ID : %v`, CreateLinkWithUserID(userid), amount, pid)
}
func ResponseConfirmWithdraw(txid string) string {
	return fmt.Sprintf(`Your withdrawal has been confirmed and sent
	
	TXID : https://blockchair.com/dogecoin/transaction/%v`, txid)
}
func ResponseRejectWithdraw(pid int) string {
	return fmt.Sprintf(`Your withdrawal request was rejected

	If you think there is a problem, contact support
	
	Request ID : %v`, pid)
}
func ResponseSubmitDepoist(amount float64, txid string) string {
	return fmt.Sprintf(`You received a deposit
	It will be added to your balance after 10 confirmations

	Amount : %v
	
	TXID: %v`, amount, txid)
}
func ResponseConfirmDepoist(amount float64, txid string) string {
	return fmt.Sprintf(`Your deposit has been confirmed
	Amount : %v
	
	TXID: %v`, amount, txid)
}
func DEPOSIT(deposit string) string {
	return fmt.Sprintf(`ℹ️ **Deposit**
	
	💳 **Wallet Address**: %v`, deposit)
}
func WithdrawConfirm(amount float32, userid int64) string {
	userdata := GetUserFromDB(userid)
	return fmt.Sprintf(`** Withdraw 📤 **

	💸 __Amount__ : **%vÐ**
	
	💳 __Wallet Address__ : %v
	
	__ If You Want **Withdraw?**__
	
	✅Confirm Withdraw `, amount, userdata.Wallet)
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

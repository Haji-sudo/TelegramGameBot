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

func DiceDetails(balance float32, minbet int, maxbet int) string {
	return fmt.Sprintf(`** Dice 🎲 **

	__If Choose 2 Numbers:__
	__Right Guess__ :** 2x 😍️ **
	__Wrong Guess__ :** 0x  🥺️ **
	
	__If Choose 1 Number:__
	__Right Guess__ :** 4x 😍️ **
	__Wrong Guess__ :** 0x   🥺️ **
	
	💰 __Balance__: **%vÐ**

	❗ __The Minimum Bet Amount Is__ **%vÐ**
	❗ __The Maximum Bet Amount Is__ **%vÐ**

	💸 __Send The Required Bet Amount__`, balance, minbet, maxbet)
}

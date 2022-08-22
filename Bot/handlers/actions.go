package handlers

import (
	"fmt"
)

func START(name string, link string) string {
	return fmt.Sprintf(`** Hello ** [%v](%v)

__Welcome to Doge Finance 🐶__
	
__You Can **INVEST** or **PLAY** and **BET** here__
	
__Also You Can Invite Your Friends And Earn **DOGE Coin**__
	
**For More Information GoTo FAQ ❓ Menu**`, name, link)
}

func ACCOUNT(name string, balance float64, totaldeposit float64, totalwithdraw float64, referrals int, EarnFromRef int, warning int, wallet string) string {
	return fmt.Sprintf(`ℹ️ **INFORMATION**

	👤 __Name__: **%v**
	
	💰 __Balance__: **%vÐ**
	
	📥 __Total Deposit__: **%vÐ**
	
	📤 __Total Withdraw__: **%vÐ**
	
	👥 __Total Referrals__: **%v**
	
	🎁 __Total Earn From Referrals__: **%vÐ**
	
	⛔ __Warning__: **%v**
	
	💳 **Wallet Address**: %v`, name, balance, totaldeposit, totalwithdraw, referrals, EarnFromRef, warning, wallet)
}

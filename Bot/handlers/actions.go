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

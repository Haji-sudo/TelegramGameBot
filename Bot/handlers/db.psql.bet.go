package handlers

func BetWin(userid int64, amount float32) {
	user := GetUserFromDB(userid)
	user.Balance += amount
	DB.Save(&user)
}
func ConfirmBet(userid int64, amount float32) {
	user := GetUserFromDB(userid)
	user.Balance -= amount
	DB.Save(&user)
}
func GetGamesHistory(userid int64) []Bet {
	bets := []Bet{}
	DB.Model(&Bet{}).Order("date DESC").Where("user_refer = ?", userid).Limit(10).Find(&bets)
	return bets
}

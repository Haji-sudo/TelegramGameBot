package handlers

import (
	"log"
	"time"
)

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

func SaveGameHistroy(userid int64, gametype string, amount float32, result string) {
	bet := Bet{UserRefer: userid, Type: gametype, Amount: amount, Date: time.Now(), Result: result}
	if result := DB.Create(&bet); result.Error != nil {
		log.Println(result.Error)
	}
}

package handlers

import (
	gateway "dogegambling/Gateway"
	"dogegambling/config"
	"fmt"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func CreateInPostgres(userid int64) {
	user := User{UserID: userid}
	if result := DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
}
func UserExistInDB(userid int64) bool {
	user := User{}
	result := DB.Select(&user, userid)
	return result.RowsAffected < 0
}
func GetUserFromDB(userid int64) User {
	user := User{}

	result := DB.First(&user, userid)
	if result.RowsAffected < 0 {
		return user
	}

	return user
}
func (u *User) UpdateWalletAddress(wallet string) {
	u.Wallet = wallet
	DB.Save(&u)
}
func (u *User) AddReferral() {
	u.Referrals += 1
	u.Balance += float32(config.Gift)
	DB.Save(&u)
}
func (u *User) CreateDepositAddress() {
	u.DepositAddress = gateway.GenerateAddress()
	DB.Save(&u)
}

func UserBalance(userid int64) float32 {
	user := GetUserFromDB(int64(userid))
	return user.Balance
}

func SubmitWithdawInUser(userid int64, amount float32) {
	user := GetUserFromDB(userid)
	user.Balance -= amount
	DB.Save(&user)
}
func GetWalletAddress(userid int64) string {
	user := GetUserFromDB(userid)
	return user.Wallet
}
func GetUserByDepositAddress(address string) int64 {
	user := User{}
	DB.Model(&User{}).Where("deposit_address = ?", address).Find(&user)
	return user.UserID
}
func ConfirmDepositInUser(userid int64, amount float32) {
	user := GetUserFromDB(userid)
	user.Balance += amount
	DB.Save(&user)
}

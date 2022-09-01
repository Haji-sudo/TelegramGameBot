package handlers

import (
	"log"
	"time"
)

func GetDepositHistory(userid int64) []Payment {
	payments := []Payment{}
	DB.Model(&Payment{}).Order("date DESC").Where("user_refer = ?", userid).Where("type = ?", true).Limit(10).Find(&payments)
	return payments
}
func GetWithdrawHistory(userid int64) []Payment {
	payments := []Payment{}
	DB.Model(&Payment{}).Order("date DESC").Where("user_refer = ?", userid).Where("type = ?", false).Limit(10).Find(&payments)
	return payments
}
func SubmitWithdraw(userid int64, amount float32) int {
	SubmitWithdawInUser(userid, amount)
	payment := Payment{UserRefer: userid, Date: time.Now(), Amount: amount, Type: false, Status: "Pending"}
	if result := DB.Create(&payment); result.Error != nil {
		log.Println(result.Error)
	}
	return payment.PID
}
func GetPaymentByID(pid int) Payment {
	payment := Payment{}
	DB.Model(&Payment{}).Where("p_id = ?", pid).Find(&payment)
	return payment
}
func GetPaymentByTXID(txid string) Payment {
	payment := Payment{}
	DB.Model(&Payment{}).Where("tx_id = ?", txid).Find(&payment)
	return payment
}
func ConfirmWithdraw(pid int, txid string) {
	payment := GetPaymentByID(pid)
	payment.TxID = txid
	payment.Status = "Done"
	DB.Save(&payment)
}
func RejectWithdraw(pid int) {
	payment := GetPaymentByID(pid)
	payment.Status = "Reject"
	DB.Save(&payment)
}
func SubmitDeposit(address string, amount float32, txid string) int64 {
	userid := GetUserByDepositAddress(address)
	payment := Payment{UserRefer: userid, Date: time.Now(), Amount: amount, Type: true, TxID: txid, Status: "Pending"}
	if result := DB.Create(&payment); result.Error != nil {
		log.Println(result.Error)
	}
	return userid
}
func ConfirmDeposit(txid string) int64 {
	payment := GetPaymentByTXID(txid)
	ConfirmDepositInUser(payment.UserRefer, payment.Amount)
	payment.Status = "Done"
	DB.Save(&payment)
	return payment.UserRefer
}

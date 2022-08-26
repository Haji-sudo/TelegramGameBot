package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type GenerateAddressSchema struct {
	Status string `json:"status"`
	Data   struct {
		Network string `json:"network"`
		UserID  int    `json:"user_id"`
		Address string `json:"address"`
		Label   string `json:"label"`
	} `json:"data"`
}

type GetBalanceSchema struct {
	Status string `json:"status"`
	Data   struct {
		Network                string `json:"network"`
		AvailableBalance       string `json:"available_balance"`
		PendingReceivedBalance string `json:"pending_received_balance"`
	} `json:"data"`
}

var Token = ""
var WebHookURL = ""

func Init(token string, webhook_url string) {
	Token = token
	WebHookURL = webhook_url
}

func GetBalance() float64 {

	url := fmt.Sprintf("https://block.io/api/v2/get_balance/?api_key=%v", Token)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error in GetBalance http Request : %v", err)
	}
	defer resp.Body.Close()
	jsonDataFromHttp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error in Read Body GetBalance : %v", err)
	}
	result := GetBalanceSchema{}
	err = json.Unmarshal([]byte(jsonDataFromHttp), &result)
	if err != nil {
		log.Printf("Error in Parse response to GetBalance object : %v", err)
	}
	AvailableBalance, _ := strconv.ParseFloat(result.Data.AvailableBalance, 64)

	return AvailableBalance
}

func GenerateAddress() string {
	url := fmt.Sprintf("https://block.io/api/v2/get_new_address/?api_key=%v", Token)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error in GetAddress http Request : %v", err)
	}
	defer resp.Body.Close()
	jsonDataFromHttp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error in Read Body GenerateAddress : %v", err)
	}
	result := GenerateAddressSchema{}
	err = json.Unmarshal([]byte(jsonDataFromHttp), &result)
	if err != nil {
		log.Printf("Error in Parse response to GetAddress object : %v", err)
	}
	return result.Data.Address
}

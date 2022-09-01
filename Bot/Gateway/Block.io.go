package gateway

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"runtime"
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
type WithdrawSchema struct {
	Status string `json:"status"`
	Data   struct {
		Network string `json:"network"`
		Txid    string `json:"txid"`
	} `json:"data"`
}
type ValidateAddressSchema struct {
	Status string `json:"status"`
	Data   struct {
		Network string `json:"network"`
		Address string `json:"address"`
		IsValid bool   `json:"is_valid"`
	} `json:"data"`
}

var (
	token string
	pin   string
	//webHookURL string
)

func Init(Token string, Pin string) {
	token = Token
	//webHookURL = webhook
	pin = Pin
}

func GetBalance() float64 {

	url := fmt.Sprintf("https://block.io/api/v2/get_balance/?api_key=%v", token)
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
	url := fmt.Sprintf("https://block.io/api/v2/get_new_address/?api_key=%v", token)
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
func Withdraw(amount float64, to_address string) (WithdrawSchema, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("python", "withdraw.py", token, pin, fmt.Sprintf("%f", amount), to_address)
	} else {
		cmd = exec.Command("python3", "withdraw.py", token, pin, fmt.Sprintf("%f", amount), to_address)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Error in cmd.StdoutPipe() : %v", err)
		return WithdrawSchema{}, err
	}
	err = cmd.Start()
	if err != nil {
		log.Printf("Error in cmd.Start() : %v", err)
		return WithdrawSchema{}, err
	}
	scanner := bufio.NewScanner(stdout)
	result := ""
	for scanner.Scan() {
		result += scanner.Text()
	}

	data := WithdrawSchema{}
	json.Unmarshal([]byte(result), &data)
	return data, nil
}
func ValidateAddress(address string) bool {
	url := fmt.Sprintf("https://block.io/api/v2/is_valid_address/?api_key=%v&address=%v", token, address)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error in ValidateAddress http Request : %v", err)
	}
	defer resp.Body.Close()
	jsonDataFromHttp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error in Read Body ValidateAddress : %v", err)
	}
	result := ValidateAddressSchema{}
	err = json.Unmarshal([]byte(jsonDataFromHttp), &result)
	if err != nil {
		log.Printf("Error in Parse response to ValidateAddress object : %v", err)
	}

	return result.Data.IsValid
}

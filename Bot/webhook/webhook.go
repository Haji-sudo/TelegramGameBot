package webhook

import (
	h "dogegambling/handlers"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/telebot.v3"
)

type NotifResponse struct {
	NotificationID  string `json:"notification_id"`
	DeliveryAttempt int    `json:"delivery_attempt"`
	CreatedAt       int    `json:"created_at"`
	Type            string `json:"type"`
	Data            struct {
		Network        string `json:"network"`
		Address        string `json:"address"`
		BalanceChange  string `json:"balance_change"`
		AmountSent     string `json:"amount_sent"`
		AmountReceived string `json:"amount_received"`
		Txid           string `json:"txid"`
		Confirmations  string `json:"confirmations"`
		IsGreen        bool   `json:"is_green"`
	} `json:"data"`
}

func Serve(Bot *telebot.Bot) {
	log.Println("Serve WebHook")
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		result := NotifResponse{}
		json.Unmarshal([]byte(b), &result)
		confirms, _ := strconv.Atoi(result.Data.Confirmations)
		if confirms == 0 {
			amount, _ := strconv.ParseFloat(result.Data.BalanceChange, 64)
			userid := h.SubmitDeposit(result.Data.Address, float32(amount), result.Data.Txid)
			h.SendToUser(Bot, userid, h.ResponseSubmitDepoist(amount, result.Data.Txid))
		} else if confirms == 10 {
			amount, _ := strconv.ParseFloat(result.Data.BalanceChange, 64)
			userid := h.ConfirmDeposit(result.Data.Txid)
			h.SendToUser(Bot, userid, h.ResponseSubmitDepoist(amount, result.Data.Txid))
		}
	})
	log.Fatal(http.ListenAndServe(":8585", nil))
}

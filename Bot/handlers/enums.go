package handlers

import "math/rand"

const (
	Main      = "main"
	Withdraw1 = "withdraw1"
	Account1  = "acc1"
	Account2  = "acc2"
	Games     = "games"
	Dice1     = "dice1"
	Dice2     = "dice2"
	Dice3     = "dice3"
	Dice4     = "dice4"
	Bowl1     = "bowl1"
	Bowl2     = "bowl2"
	Dart1     = "dart1"
	Dart2     = "dart2"
	Slot1     = "slot1"
	Slot2     = "slot2"
	Basket1   = "basket1"
	Basket2   = "basket2"

	Minbet float64 = 1
	Maxbet float64 = 200
)

func GetRandomWinNumber() float32 {
	var rnd = []float32{1.01, 1.02, 1.03, 1.04, 1.05, 1.06, 1.07, 1.08, 1.09,
		1.1, 1.11, 1.12, 1.13, 1.14, 1.15, 1.16, 1.17, 1.18, 1.19, 1.2}
	return rnd[rand.Intn(len(rnd))]
}

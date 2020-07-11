package main

import (
	"log"

	"github.com/ramezanius/crypex/exchange/binance"
)

var Binance *binance.Binance

func main() {
	var err error

	Binance, err = binance.New("YOUR_BINANCE_PUBLIC_KEY", "YOUR_BINANCE_SECRET_KEY")
	if err != nil {
		log.Panic(err)
	}

	GetSymbols()

	GetBalances()

	NewOrder()
	CancelOrder()
}

func GetSymbols() {
	symbols, err := Binance.GetSymbols()
	if err != nil {
		log.Panic(err)
	}

	log.Println(symbols)
}

func GetBalances() {
	balances, err := Binance.GetBalances()
	if err != nil {
		log.Panic(err)
	}

	log.Println(balances)
}

func NewOrder() {
	order, err := Binance.NewOrder(binance.NewOrder{
		Price:    10,
		Quantity: 1,

		Side:        binance.Buy,
		Type:        binance.Limit,
		Symbol:      binance.BNB + binance.USD,
		TimeInForce: binance.GoodTillCancel,
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}

func CancelOrder() {
	order, err := Binance.CancelOrder("FAKE_ORDER_ID", binance.BNB+binance.USD)
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}

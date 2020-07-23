package main

import (
	"log"

	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/binance/converter"
	"github.com/ramezanius/crypex/exchange/tests"
)

var Binance *binance.Binance

func main() {
	Binance = binance.New()
	Binance.PublicKey = "YOUR_BINANCE_PUBLIC_KEY"
	Binance.SecretKey = "YOUR_BINANCE_SECRET_KEY"

	ToUSD()

	GetSymbols()
	GetBalances()

	NewOrder()
	CancelOrder()
}

func ToUSD() {
	cache := &tests.Repository{}
	value, err := converter.ToUSD(cache, binance.BTC, 10, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(value)
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

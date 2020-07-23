package main

import (
	"log"

	"github.com/ramezanius/crypex/exchange/hitbtc"
	"github.com/ramezanius/crypex/exchange/hitbtc/converter"
	"github.com/ramezanius/crypex/exchange/tests"
)

var HitBTC *hitbtc.HitBTC

func main() {
	HitBTC = hitbtc.New()
	HitBTC.PublicKey = "YOUR_HITBTC_PUBLIC_KEY"
	HitBTC.SecretKey = "YOUR_HITBTC_SECRET_KEY"

	ToUSD()

	GetSymbols()
	GetBalances()

	NewOrder()
	CancelOrder()
}

func ToUSD() {
	cache := &tests.Repository{}
	value, err := converter.ToUSD(cache, hitbtc.BTC, 10, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(value)
}

func GetSymbols() {
	symbols, err := HitBTC.GetSymbols()
	if err != nil {
		log.Panic(err)
	}

	log.Println(symbols)
}

func GetBalances() {
	balances, err := HitBTC.GetBalances()
	if err != nil {
		log.Panic(err)
	}

	log.Println(balances)
}

func NewOrder() {
	order, err := HitBTC.NewOrder(hitbtc.NewOrder{
		Price:    2000,
		Quantity: 0.00002,

		Side:   hitbtc.Buy,
		Type:   hitbtc.Limit,
		Symbol: hitbtc.BTC + hitbtc.USD,
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}

func CancelOrder() {
	order, err := HitBTC.CancelOrder("FAKE_ORDER_ID")
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}

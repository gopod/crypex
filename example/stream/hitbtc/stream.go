package main

import (
	"log"
	"sync"

	"github.com/ramezanius/crypex/exchange/hitbtc"
)

var HitBTC *hitbtc.HitBTC

func main() {
	var err error

	HitBTC, err = hitbtc.New("YOUR_HITBTC_PUBLIC_KEY", "YOUR_HITBTC_SECRET_KEY")
	if err != nil {
		log.Panic(err)
	}

	GetSymbols()

	GetBalances()

	NewOrder()
	CancelOrder()
	ReplaceOrder()

	SubscribeReports()
	SubscribeCandles()
	UnsubscribeCandles()
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
		Symbol: hitbtc.Demo,
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

func ReplaceOrder() {
	order, err := HitBTC.ReplaceOrder(hitbtc.ReplaceOrder{
		Price:    1000,
		Quantity: 0.00001,
		OrderID:  "FAKE_ORDER_ID",
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println(order)
}

func SubscribeReports() {
	reports, err := HitBTC.SubscribeReports()
	if err != nil {
		log.Panic(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	defer wg.Done()

	for {
		log.Println(<-reports)
	}
}

func SubscribeCandles() {
	candles, err := HitBTC.SubscribeCandles(
		hitbtc.CandlesParams{
			Limit:  100,
			Symbol: hitbtc.Demo,
			Period: hitbtc.Period1Minute,
		})
	if err != nil {
		log.Panic(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	defer wg.Done()

	for {
		log.Println(<-candles)
	}
}

func UnsubscribeCandles() {
	err := HitBTC.UnsubscribeCandles(
		hitbtc.CandlesParams{
			Symbol: hitbtc.Demo,
		})
	if err != nil {
		log.Panic(err)
	}
}

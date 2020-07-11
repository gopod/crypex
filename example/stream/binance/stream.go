package main

import (
	"log"
	"sync"

	"github.com/ramezanius/crypex/exchange/binance"
)

var Binance *binance.Binance

func main() {
	var err error

	Binance, err = binance.New("YOUR_BINANCE_PUBLIC_KEY", "YOUR_BINANCE_SECRET_KEY")
	if err != nil {
		log.Panic(err)
	}

	SubscribeReports()
	SubscribeCandles()
	UnsubscribeCandles()
}

func SubscribeReports() {
	candles, err := Binance.SubscribeReports()
	if err != nil {
		log.Panic(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	defer wg.Done()

	for {
		log.Println(<-candles)
	}
}

func SubscribeCandles() {
	candles, err := Binance.SubscribeCandles(
		binance.CandlesParams{
			Symbol: binance.Demo,
			Period: binance.Period1Minute,
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
	err := Binance.UnsubscribeCandles(
		binance.CandlesParams{
			Symbol: binance.Demo,
		})
	if err != nil {
		log.Panic(err)
	}
}

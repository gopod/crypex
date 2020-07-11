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

	SubscribeReports()
	SubscribeCandles()
	UnsubscribeCandles()
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

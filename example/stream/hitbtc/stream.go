package main

import (
	"log"
	"sync"

	"github.com/gopod/crypex/exchange/hitbtc"
)

var HitBTC *hitbtc.HitBTC

func main() {
	HitBTC = hitbtc.New()

	HitBTC.PublicKey = "YOUR_HITBTC_PUBLIC_KEY"
	HitBTC.SecretKey = "YOUR_HITBTC_SECRET_KEY"

	wg := sync.WaitGroup{}
	wg.Add(1)

	go receiveStreams()

	err := HitBTC.SubscribeReports()
	if err != nil {
		log.Panic(err)
	}

	err = HitBTC.SubscribeCandles(
		hitbtc.CandlesParams{
			Symbol: hitbtc.BTC + hitbtc.USD,
			Period: hitbtc.Period1Minute,
		})
	if err != nil {
		log.Panic(err)
	}

	wg.Wait()
}

func receiveStreams() {
	go func() {
		for candles := range HitBTC.Feeds.Candles {
			log.Println(candles)
		}
	}()

	go func() {
		for report := range HitBTC.Feeds.Reports {
			log.Println(report)
		}
	}()
}

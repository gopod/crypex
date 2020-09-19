package main

import (
	"log"
	"sync"

	"github.com/gopod/crypex/exchange/binance"
)

var Binance *binance.Binance

func main() {
	Binance = binance.New()

	Binance.PublicKey = "YOUR_BINANCE_PUBLIC_KEY"
	Binance.SecretKey = "YOUR_BINANCE_SECRET_KEY"

	wg := sync.WaitGroup{}
	wg.Add(1)

	go receiveStreams()

	err := Binance.SubscribeReports()
	if err != nil {
		log.Panic(err)
	}

	err = Binance.SubscribeCandles(
		binance.CandlesParams{
			Symbol: binance.BTC + binance.USD,
			Period: binance.Period1Minute,
		})
	if err != nil {
		log.Panic(err)
	}

	wg.Wait()
}

func receiveStreams() {
	go func() {
		for candles := range Binance.Feeds.Candles {
			log.Println(candles)
		}
	}()

	go func() {
		for report := range Binance.Feeds.Reports {
			log.Println(report)
		}
	}()
}

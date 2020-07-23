package main

import (
	"log"

	"github.com/ramezanius/crypex/exchange/hitbtc"
)

var HitBTC *hitbtc.HitBTC

func main() {
	HitBTC = hitbtc.New()
	HitBTC.PublicKey = "YOUR_HITBTC_PUBLIC_KEY"
	HitBTC.SecretKey = "YOUR_HITBTC_SECRET_KEY"

	SubscribeReports()
	SubscribeCandles(hitbtc.CandlesParams{
		Symbol: hitbtc.BTC + hitbtc.USD,
		Period: hitbtc.Period1Minute,
	})
}

func SubscribeReports() {
	defer func() {
		err := HitBTC.UnsubscribeReports()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := HitBTC.SubscribeReports(
		func(response interface{}) {
			log.Println(response)
		})
	if err != nil {
		log.Panic(err)
	}
}

func SubscribeCandles(params hitbtc.CandlesParams) {
	defer func() {
		err := HitBTC.UnsubscribeCandles(params)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := HitBTC.SubscribeCandles(
		params, func(response interface{}) {
			log.Println(response)
		})
	if err != nil {
		log.Panic(err)
	}
}

package main

import (
	"log"

	"github.com/ramezanius/crypex/exchange/binance"
)

var Binance *binance.Binance

func main() {
	Binance = binance.New()
	Binance.PublicKey = "YOUR_BINANCE_PUBLIC_KEY"
	Binance.SecretKey = "YOUR_BINANCE_SECRET_KEY"

	SubscribeReports()
	SubscribeCandles(binance.CandlesParams{
		Symbol: binance.BNB + binance.USD,
		Period: binance.Period1Minute,
	})
}

func SubscribeReports() {
	defer func() {
		err := Binance.UnsubscribeReports()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := Binance.SubscribeReports(
		func(response interface{}) {
			log.Println(response)
		})
	if err != nil {
		log.Panic(err)
	}
}

func SubscribeCandles(params binance.CandlesParams) {
	defer func() {
		err := Binance.UnsubscribeCandles(params)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := Binance.SubscribeCandles(
		params, func(response interface{}) {
			log.Println(response)
		})
	if err != nil {
		log.Panic(err)
	}
}

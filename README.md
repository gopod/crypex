Crypex
======

[![Build](https://img.shields.io/github/workflow/status/gopod/crypex/Continuous%20Integration?label=build)](https://github.com/gopod/crypex/actions)
[![Coverage](https://img.shields.io/codacy/coverage/006240f7b0a0451b9479b907e24640e2?label=coverage)](https://app.codacy.com/manual/gopod/crypex/dashboard?bid=18899044#coverageData)
[![Issues](https://img.shields.io/github/issues/gopod/crypex?label=issues)](https://github.com/gopod/crypex/issues)
[![PullRequests](https://img.shields.io/github/issues-pr/gopod/crypex?label=pull%20requests)](https://github.com/gopod/crypex/pulls)
[![CodeSize](https://img.shields.io/github/languages/code-size/gopod/crypex?label=code%20size)](https://github.com/gopod/crypex)
[![CodeQuality](https://img.shields.io/codacy/grade/006240f7b0a0451b9479b907e24640e2?label=code%20quality)](https://app.codacy.com/manual/gopod/crypex/dashboard?bid=18899044#issuesData)
[![Licence](https://img.shields.io/github/license/gopod/crypex?label=licence)](https://github.com/gopod/crypex/blob/master/LICENCE)

Crypex is a Go package for trading and communicating with [various](#supported-exchanges) exchange API for cryptocurrency assets!

Features include:

  * [Websocket streams](#websocket-streams)
  * [Asset converters](#assets-converters)

Get started:

  * Install crypex with [one line of code](#installation), or [update it with another](#staying-up-to-date)
  * Check out the [Go Doc](https://godoc.org/github.com/gopod/crypex) API Documentation
  * A little about [Crypto currency](https://en.wikipedia.org/wiki/Cryptocurrency)


Websocket streams
-----------------

Crypex provides some helpful methods that allow you to consume data in the concurrency.

See it in action:

```go
import (
	"log"

	"github.com/gopod/crypex/exchange/binance"
	"github.com/gopod/crypex/exchange/hitbtc"
)

var HitBTC *hitbtc.HitBTC
var Binance *binance.Binance

func main() {
	HitBTC = hitbtc.New()
	HitBTC.PublicKey, HitBTC.SecretKey = "YOUR_HITBTC_PUBLIC_KEY", "YOUR_HITBTC_SECRET_KEY"
	Binance = binance.New()
	Binance.PublicKey, Binance.SecretKey = "YOUR_BINANCE_PUBLIC_KEY", "YOUR_BINANCE_SECRET_KEY"

	SubscribeReports()
	SubscribeCandles()
}

func SubscribeReports() {
	defer func() {
		err := HitBTC.UnsubscribeReports()
		if err != nil {
			log.Fatal(err)
		}

		err = Binance.UnsubscribeReports()
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

	err = Binance.SubscribeReports(
		func(response interface{}) {
			log.Println(response)
		})
	if err != nil {
		log.Panic(err)
	}
}

func SubscribeCandles() {
	hitbtcParams := hitbtc.CandlesParams{
		Symbol: hitbtc.BTC + hitbtc.USD,
		Period: hitbtc.Period1Minute,
	}
	binanceParams := binance.CandlesParams{
		Symbol: binance.BNB + binance.USD,
		Period: binance.Period1Minute,
	}
	defer func() {
		err := HitBTC.UnsubscribeCandles(hitbtcParams)
		if err != nil {
			log.Fatal(err)
		}

		err = Binance.UnsubscribeCandles(binanceParams)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := HitBTC.SubscribeCandles(
		hitbtcParams, func(response interface{}) {
			log.Println(response)
		})
	if err != nil {
		log.Panic(err)
	}

	err = Binance.SubscribeCandles(
		binanceParams, func(response interface{}) {
			log.Println(response)
		})
	if err != nil {
		log.Panic(err)
	}
}

```

------

Assets converters
-----------------

With it, you can convert a symbol or currency value to another symbol value.

An example convert is shown below:

```go
import (
	"log"

	"github.com/gopod/crypex/exchange/binance"
	"github.com/gopod/crypex/exchange/hitbtc"
	"github.com/gopod/crypex/exchange/tests"

	BinanceConverter "github.com/gopod/crypex/exchange/binance/converter"
	HitBTCConverter "github.com/gopod/crypex/exchange/hitbtc/converter"
)

func main() {
	// You should define your own repository.
	cache := &tests.Repository{}

	hitbtcValue, err := HitBTCConverter.ToUSD(cache, hitbtc.BTC, 56, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(hitbtcValue)

	binanceValue, err := BinanceConverter.ToUSD(cache, binance.BTC, 24, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(binanceValue)
}

```

  * You must implement a repository that have GetPrice, GetSymbol methods.

Installation
============

To install Crypex, use `go get`:

    go get github.com/gopod/crypex

This will then make the following packages available to you:

    github.com/gopod/crypex/exchange/hitbtc
    github.com/gopod/crypex/exchange/hitbtc/converter
    github.com/gopod/crypex/exchange/binance
    github.com/gopod/crypex/exchange/binance/converter

------

Staying up to date
==================

To update Crypex to the latest version, use `go get -u github.com/gopod/crypex`.

------

Supported exchanges
=====================

Exchange | Stream methods | API methods
:-:|:-:|:-:
HitBTC | `Candles`, `Reports` | `NewOrder`, `CancelOrder`, `GetSymbols`, `GetCandles`, `GetBalances`
Binance | `Klines`, `Reports` | `NewOrder`, `CancelOrder`, `GetSymbols`, `GetCandles`,`GetBalances`

------

License
=======

This project licensed under the terms of the [MIT](LICENCE) license.

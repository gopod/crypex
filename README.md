Crypex
======

[![Build](https://img.shields.io/github/workflow/status/ramezanius/crypex/Continuous%20Integration?label=build)](https://google.com)
[![Coverage](https://img.shields.io/codacy/coverage/6996e8a7fdb845eea86f02740f57e94b?label=coverage)](https://app.codacy.com/manual/ramezanius/crypex/dashboard?bid=18899044#coverageData)
[![Issues](https://img.shields.io/github/issues/ramezanius/crypex?label=issues)](https://github.com/ramezanius/crypex/issues)
[![PullRequests](https://img.shields.io/github/issues-pr/ramezanius/crypex?label=pull%20requests)](https://github.com/ramezanius/crypex/pulls)
[![CodeSize](https://img.shields.io/github/languages/code-size/ramezanius/crypex?label=code%20size)](https://github/com/ramezanius/crypex)
[![CodeQuality](https://img.shields.io/codacy/grade/6996e8a7fdb845eea86f02740f57e94b?label=code%20quality)](https://app.codacy.com/manual/ramezanius/crypex/dashboard?bid=18899044#issuesData)
[![Licence](https://img.shields.io/github/license/ramezanius/crypex?label=licence)](https://github.com/ramezanius/crypex/blob/master/LICENCE)

Crypex is a Go package for trading and communicating with [various](#supported-exchanges) exchange API for cryptocurrency assets!

Features include:

  * [Websocket streams](#websocket-streams)
  * [Asset converters](#assets-converters)

Get started:

  * Install crypex with [one line of code](#installation), or [update it with another](#staying-up-to-date)
  * Check out the [Go Doc](https://godoc.org/github.com/ramezanius/crypex) API Documentation
  * A little about [Crypto currency](https://en.wikipedia.org/wiki/Cryptocurrency)


Websocket streams
-----------------

Crypex provides some helpful methods that allow you to consume data in the concurrency.

See it in action:

```go
import (
	"log"

	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/hitbtc"
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

	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/hitbtc"
	"github.com/ramezanius/crypex/exchange/tests"

	BinanceConverter "github.com/ramezanius/crypex/exchange/binance/converter"
	HitBTCConverter "github.com/ramezanius/crypex/exchange/hitbtc/converter"
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

    go get github.com/ramezanius/crypex

This will then make the following packages available to you:

    github.com/ramezanius/crypex/exchange/hitbtc
    github.com/ramezanius/crypex/exchange/hitbtc/converter
    github.com/ramezanius/crypex/exchange/binance

------

Staying up to date
==================

To update Crypex to the latest version, use `go get -u github.com/ramezanius/crypex`.

------

Supported exchanges
=====================

Exchange | Stream methods | API methods
:-:|:-:|:-:
HitBTC | `Candles`, `Reports` | `NewOrder`, `ReplaceOrder`, `CancelOrder`, <br/> `GetSymbol`, `GetSymbols`, `GetBalances`
Binance | `Klines`, `Reports` | `NewOrder`, `CancelOrder`, `GetSymbols`, `GetBalances`

------

License
=======

This project is licensed under the terms of the [MIT](LICENCE) license.

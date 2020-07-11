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
	"sync"

	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/hitbtc"
)

var HitBTC *hitbtc.HitBTC
var Binance *binance.Binance

func setupExchange() {
	var err error

	HitBTC, err = hitbtc.New("YOUR_HITBTC_PUBLIC_KEY", "YOUR_HITBTC_SECRET_KEY")
	if err != nil {
		log.Panic(err)
	}

	Binance, err = binance.New("YOUR_BINANCE_PUBLIC_KEY", "YOUR_BINANCE_SECRET_KEY")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	setupExchange()

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

```

------

Assets converters
-----------------

With it, you can convert a symbol or currency value to another symbol value.

An example convert is shown below:

```go
import (
	"log"

	"github.com/ramezanius/crypex/exchange/hitbtc"
	"github.com/ramezanius/crypex/exchange/hitbtc/converter"
)

func main() {
	ToUSD()
}

const price, quantity = 10000.0, 10.0

type repository struct{}

// GetPrice returns fake price (BTC/USD)
func (r *repository) GetPrice(_, _ string) float64 {
	return price
}

// GetSymbol returns fake symbol detail (BTC/USD)[Demo]
func (r *repository) GetSymbol(_, _ string) interface{} {
	return &hitbtc.Symbol{
		Base:  hitbtc.BTC,
		Quote: hitbtc.USD,
		ID:    hitbtc.Demo,
	}
}

func ToUSD() {
	cache := &repository{}
	value := converter.ToUSD(cache, hitbtc.BTC, quantity, false)

	log.Println(value)
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

To update Crypex to the latest version, use `go get github.com/ramezanius/crypex`.

------

Supported exchanges
=====================

Exchange | Stream methods | API methods
:-:|:-:|:-:
HitBTC | `Candles`, `Reports` | `NewOrder`, `ReplaceOrder`, `CancelOrder`
Binance | `Klines` | 

------

License
=======

This project is licensed under the terms of the [MIT](LICENCE) license.

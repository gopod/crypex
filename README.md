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

Get started:

  * Install crypex with [one line of code](#installation), or [update it with another](#staying-up-to-date)
  * Check out the [Go Doc](https://godoc.org/github.com/ramezanius/crypex) API Documentation
  * A little about [Crypto currency](https://en.wikipedia.org/wiki/Cryptocurrency)


### Exchanges
The following cryptocurrency exchanges are supported:  

Exchanges | Supported Methods
:-:|:-:
HitBTC | `Orderbook`, `Candles`, `Reports`

### Quick examples
```go
package main

import (
	"fmt"

	"github.com/ramezanius/crypex/hitbtc"
)

const (
	Public = "YOUR_PUBLIC_KEY"
	Secret = "YOUR_SECRET_KEY"
)

func main() {
	// New hitbtc client
	client, err := hitbtc.New()
	if err != nil {
		panic(err)
	}

	// Authenticate
	client.PublicKey, client.SecretKey = Public, Secret
	err = client.Authenticate()
	if err != nil {
		panic(err)
	}

Staying up to date
==================

To update Crypex to the latest version, use `go get github.com/ramezanius/crypex`.

------

Supported exchanges
=====================

Exchange | Stream methods | API methods
:-:|:-:|:-:
HitBTC | `Candles`, `Reports` | `NewOrder`, `ReplaceOrder`, `CancelOrder`
Binance | `Klines`, `Reports`, `Balances` | `NewOrder`, `ReplaceOrder`, `CancelOrder`

------

License
=======

This project is licensed under the terms of the [MIT](LICENCE) license.

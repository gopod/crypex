<img src=".github/crypex.png" alt="Crypex Logo" />

![Build](https://img.shields.io/github/workflow/status/ramezanius/crypex/Crypex?label=build)
![Coverage](https://img.shields.io/codacy/coverage/6996e8a7fdb845eea86f02740f57e94b?label=coverage)
![Issues](https://img.shields.io/github/issues/ramezanius/crypex?label=issues)
![PullRequests](https://img.shields.io/github/issues-pr/ramezanius/crypex?label=pull%20requests)
![CodeSize](https://img.shields.io/github/languages/code-size/ramezanius/crypex?label=code%20size)
![CodeQuality](https://img.shields.io/codacy/grade/6996e8a7fdb845eea86f02740f57e94b?label=code%20quality)
![Licence](https://img.shields.io/github/license/ramezanius/crypex?label=licence)

## Overview
Crypex is a Go package for trading and communicating with [various](#Exchanges) exchange API for cryptocurrency assets!

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

	// Get balances
	balances, err := client.GetBalances()
	fmt.Println(err, balances)

	// Subscribe and consume data
	candlesFeed, err := client.SubscribeCandles("BTCUSD", "M1", 100)

	// If you subscribed many symbols, So you should use goroutines =)
	for {
		candles := <-candlesFeed
		fmt.Println(snapshot)
	}
}
```

## Licence
[MIT](LICENCE)

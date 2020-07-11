package main

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

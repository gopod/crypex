package convert

import (
	"github.com/forestgiant/sliceutil"

	"github.com/ramezanius/crypex/hitbtc"
)

// Caching repository interface
type Repository interface {
	GetPrice(symbol string, exchange string) float64

	GetSymbol(symbol string, exchange string) interface{}
}

// Convert an currency to available pair
func ToSymbol(cache Repository, currency string) (symbol *hitbtc.Symbol) {
	baseCurrencies := []string{hitbtc.USD, hitbtc.BTC, hitbtc.ETH}

	if sliceutil.Contains(baseCurrencies, currency) {
		symbol = &hitbtc.Symbol{
			Base:  currency,
			Quote: hitbtc.USD,
		}

		if symbol.Base == hitbtc.USD {
			symbol.Base = hitbtc.BTC
		}

		return
	}

	for _, base := range baseCurrencies {
		symbol = cache.GetSymbol(currency+base, hitbtc.Exchange).(*hitbtc.Symbol)
		if symbol.ID != "" {
			break
		}
	}

	return
}

// Convert any value of symbol(name) to USD
func ToUSD(cache Repository, name string, value float64, pure bool) float64 {
	switch {
	case value == 0 || name == hitbtc.USD:
		return value
	case name == hitbtc.BTC || name == hitbtc.ETH:
		BaseUsd := cache.GetPrice(name+hitbtc.USD, hitbtc.Exchange)

		return value * BaseUsd
	}

	symbol := ToSymbol(cache, name)

	switch symbol.Quote {
	case hitbtc.USD:
		if !pure {
			BaseUsd := cache.GetPrice(symbol.Base+hitbtc.USD, hitbtc.Exchange)

			return value * BaseUsd
		}

		return value
	case hitbtc.BTC:
		BaseBtc := cache.GetPrice(symbol.Base+hitbtc.BTC, hitbtc.Exchange)
		BtcUsd := cache.GetPrice(hitbtc.BTC+hitbtc.USD, hitbtc.Exchange)

		if !pure {
			return value * BtcUsd * BaseBtc
		}

		return value * BtcUsd
	case hitbtc.ETH:
		BaseEth := cache.GetPrice(symbol.Base+hitbtc.ETH, hitbtc.Exchange)
		EthUsd := cache.GetPrice(hitbtc.ETH+hitbtc.USD, hitbtc.Exchange)

		if !pure {
			return value * EthUsd * BaseEth
		}

		return value * EthUsd
	}

	return value
}

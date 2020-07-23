package converter

import (
	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/hitbtc"
	"github.com/ramezanius/crypex/exchange/util"
)

var hitbtcCurrencies = []string{hitbtc.USD, hitbtc.BTC, hitbtc.ETH}

// Repository caching exchange repository
type Repository interface {
	GetPrice(symbol string, exchange string) float64

	GetSymbol(symbol string, exchange string) interface{}
}

// ToSymbol convert an currency to available pair.
func ToSymbol(cache Repository, currency string) (symbol *hitbtc.Symbol, err error) {
	if len(currency) >= 6 {
		symbol = cache.GetSymbol(currency, hitbtc.Exchange).(*hitbtc.Symbol)
		if symbol.ID == "" {
			return nil, binance.ErrSymbolNotFound
		}

		return
	}

	if util.Contains(hitbtcCurrencies, currency) {
		symbol = &hitbtc.Symbol{
			Base:  currency,
			Quote: hitbtc.USD,
		}

		if symbol.Base == hitbtc.USD {
			symbol.Base = hitbtc.BTC
		}

		symbol.ID = symbol.Base + symbol.Quote

		return
	}

	for _, base := range hitbtcCurrencies {
		symbol = cache.GetSymbol(currency+base, hitbtc.Exchange).(*hitbtc.Symbol)
		if symbol.ID != "" {
			return
		}
	}

	return nil, hitbtc.ErrCurrencyNotFound
}

// ToUSD convert any value of symbol(name) to hitbtc.USD.
func ToUSD(cache Repository, name string, value float64, pure bool) (result float64, err error) {
	switch {
	case value == 0 || name == hitbtc.USD:
		return value, nil
	case name == hitbtc.BTC || name == hitbtc.ETH:
		BaseUsd := cache.GetPrice(name+hitbtc.USD, hitbtc.Exchange)

		return value * BaseUsd, nil
	}

	symbol, err := ToSymbol(cache, name)
	if err != nil {
		return 0, err
	}

	switch symbol.Quote {
	case hitbtc.USD:
		if !pure {
			BaseUsd := cache.GetPrice(symbol.Base+hitbtc.USD, hitbtc.Exchange)

			return value * BaseUsd, err
		}

		return value, err
	case hitbtc.BTC:
		BaseBtc := cache.GetPrice(symbol.Base+hitbtc.BTC, hitbtc.Exchange)
		BtcUsd := cache.GetPrice(hitbtc.BTC+hitbtc.USD, hitbtc.Exchange)

		if !pure {
			return value * BtcUsd * BaseBtc, err
		}

		return value * BtcUsd, err
	case hitbtc.ETH:
		BaseEth := cache.GetPrice(symbol.Base+hitbtc.ETH, hitbtc.Exchange)
		EthUsd := cache.GetPrice(hitbtc.ETH+hitbtc.USD, hitbtc.Exchange)

		if !pure {
			return value * EthUsd * BaseEth, err
		}

		return value * EthUsd, err
	default:
		return value, nil
	}
}

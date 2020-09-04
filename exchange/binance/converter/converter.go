package converter

import (
	"github.com/gopod/crypex/exchange/binance"
	"github.com/gopod/crypex/exchange/util"
)

var binanceCurrencies = []string{binance.USD, binance.BTC, binance.ETH, binance.BNB}

// Repository caching exchange repository
type Repository interface {
	GetPrice(symbol string, exchange string) float64

	GetSymbol(symbol string, exchange string) interface{}
}

// ToSymbol convert an currency to available pair.
func ToSymbol(cache Repository, currency string) (symbol *binance.Symbol, err error) {
	if len(currency) >= 6 {
		symbol = cache.GetSymbol(currency, binance.Exchange).(*binance.Symbol)
		if symbol.ID == "" {
			return nil, binance.ErrSymbolNotFound
		}

		return
	}

	if util.Contains(binanceCurrencies, currency) {
		symbol = &binance.Symbol{
			Base:  currency,
			Quote: binance.USD,
		}

		if symbol.Base == binance.USD {
			symbol.Base = binance.BTC
		}

		symbol.ID = symbol.Base + symbol.Quote

		return
	}

	for _, base := range binanceCurrencies {
		symbol = cache.GetSymbol(currency+base, binance.Exchange).(*binance.Symbol)
		if symbol.ID != "" {
			return
		}
	}

	return nil, binance.ErrCurrencyNotFound
}

// ToUSD convert any value of symbol(name) to binance.USD.
func ToUSD(cache Repository, name string, value float64, pure bool) (result float64, err error) {
	switch {
	case value == 0 || name == binance.USD:
		return value, err
	case name == binance.BTC || name == binance.ETH:
		BaseUsd := cache.GetPrice(name+binance.USD, binance.Exchange)

		return value * BaseUsd, err
	}

	symbol, err := ToSymbol(cache, name)
	if err != nil {
		return 0, err
	}

	switch symbol.Quote {
	case binance.USD:
		if !pure {
			BaseUsd := cache.GetPrice(symbol.Base+binance.USD, binance.Exchange)

			return value * BaseUsd, err
		}

		return value, err
	case binance.BTC:
		BaseBtc := cache.GetPrice(symbol.Base+binance.BTC, binance.Exchange)
		BtcUsd := cache.GetPrice(binance.BTC+binance.USD, binance.Exchange)

		if !pure {
			return value * BtcUsd * BaseBtc, err
		}

		return value * BtcUsd, err
	case binance.ETH:
		BaseEth := cache.GetPrice(symbol.Base+binance.ETH, binance.Exchange)
		EthUsd := cache.GetPrice(binance.ETH+binance.USD, binance.Exchange)

		if !pure {
			return value * EthUsd * BaseEth, err
		}

		return value * EthUsd, err
	case binance.BNB:
		BaseBnb := cache.GetPrice(symbol.Base+binance.BNB, binance.Exchange)
		BnbUsd := cache.GetPrice(binance.BNB+binance.USD, binance.Exchange)

		if !pure {
			return value * BnbUsd * BaseBnb, err
		}

		return value * BnbUsd, err
	default:
		return value, nil
	}
}

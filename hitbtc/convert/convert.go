package convert

import (
	"github.com/forestgiant/sliceutil"

	"github.com/ramezanius/crypex/hitbtc"
)

type Repository interface {
	GetPrice(symbol string) float64

	GetSymbol(symbol string) *hitbtc.Symbol
}

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
		symbol = cache.GetSymbol(currency + base)
		if symbol.ID != "" {
			break
		}
	}

	return
}

func ToUSD(cache Repository, name string, value float64, pure bool) float64 {
	switch {
	case value == 0 || name == hitbtc.USD:
		return value
	case name == hitbtc.BTC || name == hitbtc.ETH:
		BaseUsd := cache.GetPrice(name + hitbtc.USD)

		return value * BaseUsd
	}

	symbol := ToSymbol(cache, name)

	switch {
	case symbol.Quote == hitbtc.USD:
		if !pure {
			BaseUsd := cache.GetPrice(symbol.Base + hitbtc.USD)

			return value * BaseUsd
		}

		return value
	case symbol.Quote == hitbtc.BTC:
		BaseBtc := cache.GetPrice(symbol.Base + hitbtc.BTC)
		BtcUsd := cache.GetPrice(hitbtc.BTC + hitbtc.USD)

		if !pure {
			return value * BtcUsd * BaseBtc
		}

		return value * BtcUsd
	case symbol.Quote == hitbtc.ETH:
		BaseEth := cache.GetPrice(symbol.Base + hitbtc.ETH)
		EthUsd := cache.GetPrice(hitbtc.ETH + hitbtc.USD)

		if !pure {
			return value * EthUsd * BaseEth
		}

		return value * EthUsd
	}

	return value
}

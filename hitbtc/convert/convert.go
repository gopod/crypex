package convert

import (
	"github.com/forestgiant/sliceutil"

	"github.com/ramezanius/crypex/hitbtc"
)

type Repository interface {
	GetPrice(symbol string) (float64, error)

	GetSymbol(symbol string) (*hitbtc.Symbol, error)
}

func ToSymbol(repo Repository, currency string) (symbol *hitbtc.Symbol) {
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

	var err error

	for _, base := range baseCurrencies {
		if symbol, err = repo.GetSymbol(currency + base); err == nil {
			break
		}
	}

	return
}

func ToUSD(repo Repository, name string, value float64, pure bool) (float64, error) {
	if value == 0 {
		return 0, nil
	}

	if name == hitbtc.USD {
		return value, nil
	} else if name != "" {
		if name == hitbtc.BTC || name == hitbtc.ETH {
			if Usd, err := repo.GetPrice(name + hitbtc.USD); err == nil {
				return value * Usd, nil
			}
		}
	}

	symbol := ToSymbol(repo, name)

	switch {
	case symbol.Quote == hitbtc.USD:
		if !pure {
			if BaseEth, err := repo.GetPrice(symbol.Base + hitbtc.USD); err == nil {
				return value * BaseEth, nil
			}
		}

		return value, nil
	case symbol.Quote == hitbtc.BTC:
		if BaseBtc, err := repo.GetPrice(symbol.Base + hitbtc.BTC); err == nil {
			BtcUsd, _ := repo.GetPrice(hitbtc.BTC + hitbtc.USD)

			if !pure {
				return value * BtcUsd * BaseBtc, nil
			}

			return value * BtcUsd, nil
		}
	case symbol.Quote == hitbtc.ETH:
		if BaseEth, err := repo.GetPrice(symbol.Base + hitbtc.ETH); err == nil {
			EthUsd, _ := repo.GetPrice(hitbtc.ETH + hitbtc.USD)

			if !pure {
				return value * EthUsd * BaseEth, nil
			}

			return value * EthUsd, nil
		}
	}

	return 0, nil
}

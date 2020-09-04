package tests

import (
	"time"

	"github.com/gopod/crypex/exchange/binance"
	"github.com/gopod/crypex/exchange/hitbtc"
)

// Wait waits 10 second to receive data from readers.
// the wait group cannot be used because the number of deltas is unknown.
func Wait() {
	time.Sleep(time.Second * 5)
}

// Repository struct
type Repository struct{}

// GetPrice returns price.
func (r *Repository) GetPrice(symbol, _ string) float64 {
	switch symbol {
	case hitbtc.BTC + hitbtc.USD:
		return 9494
	case hitbtc.ETH + hitbtc.USD:
		return 262
	case hitbtc.XRP + hitbtc.USD:
		return 0.20
	case hitbtc.XRP + hitbtc.BTC:
		return 0.00002
	case hitbtc.XRP + hitbtc.ETH:
		return 0.00076

	case binance.BTC + binance.USD:
		return 9122
	case binance.ETH + binance.USD:
		return 232
	case binance.BNB + binance.USD:
		return 17
	case binance.XRP + binance.USD:
		return 0.19
	case binance.XRP + binance.BTC:
		return 0.00002
	case binance.XRP + binance.ETH:
		return 0.00084
	case binance.XRP + binance.BNB:
		return 0.01114
	default:
		return 0
	}
}

// GetSymbol returns symbol detail.
func (r *Repository) GetSymbol(symbol, exchange string) interface{} {
	switch symbol {
	case hitbtc.BTC + hitbtc.USD:
		return &hitbtc.Symbol{
			Base:  hitbtc.BTC,
			Quote: hitbtc.USD,
			ID:    hitbtc.BTC + hitbtc.USD,
		}
	case hitbtc.ETH + hitbtc.USD:
		return &hitbtc.Symbol{
			Base:  hitbtc.ETH,
			Quote: hitbtc.USD,
			ID:    hitbtc.ETH + hitbtc.USD,
		}
	case hitbtc.XRP + hitbtc.USD:
		return &hitbtc.Symbol{
			Base:  hitbtc.XRP,
			Quote: hitbtc.USD,
			ID:    hitbtc.XRP + hitbtc.USD,
		}
	case hitbtc.XRP + hitbtc.BTC:
		return &hitbtc.Symbol{
			Base:  hitbtc.XRP,
			Quote: hitbtc.BTC,
			ID:    hitbtc.XRP + hitbtc.BTC,
		}
	case hitbtc.XRP + hitbtc.ETH:
		return &hitbtc.Symbol{
			Base:  hitbtc.XRP,
			Quote: hitbtc.ETH,
			ID:    hitbtc.XRP + hitbtc.ETH,
		}
	case binance.BTC + binance.USD:
		return &binance.Symbol{
			Base:  binance.BTC,
			Quote: binance.USD,
			ID:    binance.BTC + binance.USD,
		}
	case binance.ETH + binance.USD:
		return &binance.Symbol{
			Base:  binance.ETH,
			Quote: binance.USD,
			ID:    binance.ETH + binance.USD,
		}
	case binance.BNB + binance.USD:
		return &binance.Symbol{
			Base:  binance.BNB,
			Quote: binance.USD,
			ID:    binance.BNB + binance.USD,
		}
	case binance.XRP + binance.USD:
		return &binance.Symbol{
			Base:  binance.XRP,
			Quote: binance.USD,
			ID:    binance.XRP + binance.USD,
		}
	case binance.XRP + binance.BTC:
		return &binance.Symbol{
			Base:  binance.XRP,
			Quote: binance.BTC,
			ID:    binance.XRP + binance.BTC,
		}
	case binance.XRP + binance.ETH:
		return &binance.Symbol{
			Base:  binance.XRP,
			Quote: binance.ETH,
			ID:    binance.XRP + binance.ETH,
		}
	case binance.XRP + binance.BNB:
		return &binance.Symbol{
			Base:  binance.XRP,
			Quote: binance.BNB,
			ID:    binance.XRP + binance.BNB,
		}
	default:
		if exchange == binance.Exchange {
			return &binance.Symbol{}
		} else if exchange == hitbtc.Exchange {
			return &hitbtc.Symbol{}
		}

		return nil
	}
}

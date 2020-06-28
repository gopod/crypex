package convert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc"
	"github.com/ramezanius/crypex/hitbtc/convert"
	"github.com/ramezanius/crypex/hitbtc/tests"
)

type repository struct{}

func (r *repository) GetPrice(symbol, _ string) float64 {
	switch symbol {
	case tests.Demo:
		return 10000
	default:
		return 0
	}
}

func (r *repository) GetSymbol(symbol, _ string) interface{} {
	switch symbol {
	case hitbtc.BTC:
		return &hitbtc.Symbol{
			Base:  "BTC",
			Quote: "USD",
			ID:    "BTC" + "USD",
		}
	default:
		return &hitbtc.Symbol{}
	}
}

func TestConverters(t *testing.T) {
	cache := &repository{}

	t.Run("ToUSD", func(t *testing.T) {
		value := convert.ToUSD(cache, hitbtc.BTC, 10, false)

		assert.Equal(t, value, 10.0*10000.0)
	})
}

package convert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc"
	"github.com/ramezanius/crypex/hitbtc/convert"
)

type repository struct{}

func (r repository) GetPrice(symbol string) (float64, error) {
	switch symbol {
	case hitbtc.DemoSymbol:
		return 10000, nil
	default:
		return 0, nil
	}
}

func (r repository) GetSymbol(symbol string) (*hitbtc.Symbol, error) {
	switch symbol {
	case hitbtc.BTC:
		return &hitbtc.Symbol{
			Base:  "BTC",
			Quote: "USD",
			ID:    "BTC" + "USD",
		}, nil
	default:
		return &hitbtc.Symbol{}, nil
	}
}

func TestConverters(t *testing.T) {
	var repo repository

	t.Run("ToUSD", func(t *testing.T) {
		value, err := convert.ToUSD(repo, hitbtc.BTC, 10, false)

		assert.NoError(t, err)
		assert.Equal(t, value, 10.0*10000.0)
	})
}

package converter_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ramezanius/crypex/exchange/hitbtc"
	"github.com/ramezanius/crypex/exchange/hitbtc/converter"
)

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

type hitbtcConverterSuite struct {
	suite.Suite
}

func TestConverter(t *testing.T) {
	suite.Run(t, new(hitbtcConverterSuite))
}

func (suite *hitbtcConverterSuite) TestToUSD() {
	cache := &repository{}
	value, err := converter.ToUSD(cache, hitbtc.BTC, quantity, false)

	suite.NoError(err)
	suite.Equal(value, quantity*price)
}

package converter_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/binance/converter"
)

const price, quantity = 10000.0, 10.0

type repository struct{}

// GetPrice returns fake price (BTC/USD)
func (r *repository) GetPrice(_, _ string) float64 {
	return price
}

// GetSymbol returns fake symbol detail (BTC/USD)[Demo]
func (r *repository) GetSymbol(_, _ string) interface{} {
	return &binance.Symbol{
		Base:  binance.BTC,
		Quote: binance.USD,
		ID:    binance.Demo,
	}
}

type binanceConverterSuite struct {
	suite.Suite
}

func TestConverter(t *testing.T) {
	suite.Run(t, new(binanceConverterSuite))
}

func (suite *binanceConverterSuite) TestToUSD() {
	cache := &repository{}
	value, err := converter.ToUSD(cache, binance.BTC, quantity, false)

	suite.NoError(err)
	suite.Equal(value, quantity*price)
}

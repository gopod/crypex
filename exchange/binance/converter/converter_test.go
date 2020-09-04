package converter_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gopod/crypex/exchange/binance"
	"github.com/gopod/crypex/exchange/binance/converter"
	Tests "github.com/gopod/crypex/exchange/tests"
)

type binanceConverterSuite struct {
	suite.Suite
	repository *Tests.Repository
}

func TestConverter(t *testing.T) {
	suite.Run(t, new(binanceConverterSuite))
}

func (suite *binanceConverterSuite) TestToSymbol() {
	tests := []struct {
		base     string
		quote    string
		currency string
	}{
		{binance.XRP, binance.USD, binance.XRP},
		{binance.BTC, binance.USD, binance.USD},
		{binance.BTC, binance.USD, binance.BTC},
		{binance.ETH, binance.USD, binance.ETH},
		{binance.BNB, binance.USD, binance.BNB},
		{binance.BTC, binance.USD, binance.BTC + binance.USD},
		{binance.ETH, binance.USD, binance.ETH + binance.USD},
		{binance.BNB, binance.USD, binance.BNB + binance.USD},
	}

	for _, test := range tests {
		suite.Run(strings.ToUpper(test.currency), func() {
			symbol, err := converter.ToSymbol(
				suite.repository, test.currency)

			suite.NoError(err)
			suite.Equal(symbol, &binance.Symbol{
				Base:  test.base,
				Quote: test.quote,
				ID:    test.base + test.quote,
			})
		})
	}

	failTests := []struct {
		currency string
	}{
		{"trx"},
		{"bch"},
		{"eth" + "xxx"},
		{"btc" + "xxx"},
	}

	for _, test := range failTests {
		suite.Run(strings.ToUpper(test.currency), func() {
			_, err := converter.ToSymbol(
				suite.repository, test.currency,
			)

			suite.Error(err)
		})
	}
}

func (suite *binanceConverterSuite) TestToUSD() {
	tests := []struct {
		pure      bool
		name      string
		value     float64
		converted float64
	}{
		{false, binance.USD, 1, 1 * 1},
		{false, binance.BTC, 5, 5 * 9122},
		{false, binance.ETH, 2, 2 * 232},
		{false, binance.BNB, 10, 10 * 17},
		{true, binance.XRP + binance.USD, 4, 4 * 1},
		{true, binance.XRP + binance.BTC, 4, 4 * 9122},
		{true, binance.XRP + binance.ETH, 3, 3 * 232},
		{true, binance.XRP + binance.BNB, 2, 2 * 17},
		{false, binance.XRP + binance.USD, 4, 4 * 0.19},
		{false, binance.XRP + binance.BTC, 3, 3 * 0.00002 * 9122},
		{false, binance.XRP + binance.ETH, 2, 2 * 0.00084 * 232},
		{false, binance.XRP + binance.BNB, 3, 3 * 0.01114 * 17},
	}

	for _, test := range tests {
		name := strings.ToUpper(test.name)
		if test.pure {
			name += "_Pure"
		}

		suite.Run(name, func() {

			value, err := converter.ToUSD(
				suite.repository, test.name, test.value, test.pure,
			)

			suite.NoError(err)
			suite.Equal(value, test.converted)
		})
	}

	failTests := []struct {
		pure bool
		name string
	}{
		{false, "trx"},
		{false, "bch"},
		{false, "eth" + "xxx"},
		{false, "btc" + "xxx"},
	}

	for _, test := range failTests {
		_, err := converter.ToUSD(
			suite.repository, test.name, 1, test.pure,
		)

		suite.Error(err)
	}
}

package converter_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gopod/crypex/exchange/hitbtc"
	"github.com/gopod/crypex/exchange/hitbtc/converter"
	Tests "github.com/gopod/crypex/exchange/tests"
)

type hitbtcConverterSuite struct {
	suite.Suite
	repository *Tests.Repository
}

func TestConverter(t *testing.T) {
	suite.Run(t, new(hitbtcConverterSuite))
}

func (suite *hitbtcConverterSuite) TestToSymbol() {
	tests := []struct {
		base     string
		quote    string
		currency string
	}{
		{hitbtc.XRP, hitbtc.USD, hitbtc.XRP},
		{hitbtc.BTC, hitbtc.USD, hitbtc.USD},
		{hitbtc.BTC, hitbtc.USD, hitbtc.BTC},
		{hitbtc.ETH, hitbtc.USD, hitbtc.ETH},
		{hitbtc.BTC, hitbtc.USD, hitbtc.BTC + hitbtc.USD},
		{hitbtc.ETH, hitbtc.USD, hitbtc.ETH + hitbtc.USD},
	}

	for _, test := range tests {
		suite.Run(strings.ToUpper(test.currency), func() {
			symbol, err := converter.ToSymbol(
				suite.repository, test.currency)

			suite.NoError(err)
			suite.Equal(symbol, &hitbtc.Symbol{
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

func (suite *hitbtcConverterSuite) TestToUSD() {
	tests := []struct {
		pure      bool
		name      string
		value     float64
		converted float64
	}{
		{false, hitbtc.USD, 1, 1 * 1},
		{false, hitbtc.BTC, 5, 5 * 9494},
		{false, hitbtc.ETH, 2, 2 * 262},
		{true, hitbtc.XRP + hitbtc.USD, 4, 4 * 1},
		{true, hitbtc.XRP + hitbtc.BTC, 4, 4 * 9494},
		{true, hitbtc.XRP + hitbtc.ETH, 5, 5 * 262},
		{false, hitbtc.XRP + hitbtc.USD, 4, 4 * 0.20},
		{false, hitbtc.XRP + hitbtc.BTC, 3, 3 * 0.00002 * 9494},
		{false, hitbtc.XRP + hitbtc.ETH, 7, 7 * 0.00076 * 262},
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

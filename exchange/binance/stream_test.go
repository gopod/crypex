package binance_test

import (
	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/tests"
)

func (suite *binanceSuite) TestSubscribeReports() {
	suite.NoError(
		suite.exchange.SubscribeReports(func(interface{}) {}),
	)

	suite.TestOrders()

	tests.Wait()
	suite.NoError(suite.exchange.UnsubscribeReports())
}

func (suite *binanceSuite) TestSubscribeCandles() {
	params := binance.CandlesParams{
		Snapshot: true,
		Period:   binance.Period1Minute,
		Symbol:   binance.BNB + binance.BTC,
	}

	suite.NoError(
		suite.exchange.SubscribeCandles(params, func(interface{}) {}),
	)

	tests.Wait()
	suite.NoError(suite.exchange.UnsubscribeCandles(params))
}

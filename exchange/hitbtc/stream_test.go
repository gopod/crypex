package hitbtc_test

import (
	"github.com/ramezanius/crypex/exchange/hitbtc"
	"github.com/ramezanius/crypex/exchange/tests"
)

func (suite *hitbtcSuite) TestSubscribeReports() {
	suite.NoError(
		suite.exchange.SubscribeReports(func(interface{}) {}),
	)

	suite.TestOrders()

	tests.Wait()
	suite.NoError(suite.exchange.UnsubscribeReports())
}

func (suite *hitbtcSuite) TestSubscribeCandles() {
	params := hitbtc.CandlesParams{
		Snapshot: true,
		Period:   hitbtc.Period1Minute,
		Symbol:   hitbtc.BTC + hitbtc.USD,
	}

	suite.NoError(
		suite.exchange.SubscribeCandles(params, func(interface{}) {}),
	)

	tests.Wait()
	suite.NoError(suite.exchange.UnsubscribeCandles(params))
}

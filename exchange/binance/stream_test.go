package binance_test

import (
	"github.com/gopod/crypex/exchange/binance"
	"github.com/gopod/crypex/exchange/tests"
)

func (suite *binanceSuite) TestSubscribeReports() {
	suite.NoError(suite.exchange.SubscribeReports())

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

	suite.NoError(suite.exchange.SubscribeCandles(params))

	tests.Wait()
	suite.NoError(suite.exchange.UnsubscribeCandles(params))

	suite.Run("Fail", func() {
		params := binance.CandlesParams{
			Period: binance.Period1Minute,
			Symbol: binance.USD + binance.BTC,
		}

		suite.NoError(suite.exchange.SubscribeCandles(params))
	})
}

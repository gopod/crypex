package binance_test

import (
	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/tests"
)

func (suite *binanceSuite) TestSubscribeReports() {
	reports, err := suite.Exchange.SubscribeReports()
	suite.NoError(err)

	suite.Run("ReadReports", func() {
		tests.ReceiveWithTimeout(suite.T(), reports, &tests.DefaultTimeout)
	})
}

func (suite *binanceSuite) TestSubscribeCandles() {
	candles, err :=
		suite.Exchange.SubscribeCandles(
			binance.CandlesParams{
				Symbol: binance.Demo,
				Period: binance.Period1Hour,
			})
	suite.NoError(err)

	suite.Run("ReadCandles", func() {
		tests.ReceiveWithTimeout(suite.T(), candles, &tests.DefaultTimeout)
	})
}

func (suite *binanceSuite) TestUnsubscribeCandles() {
	err :=
		suite.Exchange.UnsubscribeCandles(
			binance.CandlesParams{
				Symbol: binance.Demo,
			})
	suite.NoError(err)
}

package hitbtc_test

import (
	"github.com/ramezanius/crypex/exchange/hitbtc"
	"github.com/ramezanius/crypex/exchange/tests"
)

func (suite *hitbtcSuite) TestSubscribeReports() {
	reports, err := suite.Exchange.SubscribeReports()
	suite.NoError(err)

	suite.Run("ReadReports", func() {
		tests.ReceiveWithTimeout(suite.T(), reports, &tests.DefaultTimeout)
	})
}

func (suite *hitbtcSuite) TestSubscribeCandles() {
	candles, err :=
		suite.Exchange.SubscribeCandles(
			hitbtc.CandlesParams{
				Limit:  100,
				Symbol: hitbtc.Demo,
				Period: hitbtc.Period1Minute,
			})
	suite.NoError(err)

	suite.Run("ReadCandles", func() {
		tests.ReceiveWithTimeout(suite.T(), candles, &tests.DefaultTimeout)
	})
}

func (suite *hitbtcSuite) TestUnsubscribeCandles() {
	err :=
		suite.Exchange.UnsubscribeCandles(
			hitbtc.CandlesParams{
				Symbol: hitbtc.Demo,
			})
	suite.NoError(err)
}

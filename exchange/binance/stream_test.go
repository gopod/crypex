package binance_test

import (
	"sync"

	"github.com/ramezanius/crypex/exchange/binance"
)

func (suite *binanceSuite) TestSubscribeReports() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	suite.NoError(
		suite.exchange.SubscribeReports(func(response interface{}) {
			wg.Done()
		}),
	)

	suite.TestOrders()

	wg.Wait()
	suite.NoError(suite.exchange.UnsubscribeReports())
}

func (suite *binanceSuite) TestSubscribeCandles() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	params := binance.CandlesParams{
		Period: binance.Period1Minute,
		Symbol: binance.BNB + binance.BTC,
	}

	suite.NoError(
		suite.exchange.SubscribeCandles(params, func(response interface{}) {
			wg.Done()
		}),
	)

	wg.Wait()
	suite.NoError(suite.exchange.UnsubscribeCandles(params))
}

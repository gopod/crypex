package hitbtc_test

import (
	"sync"

	"github.com/ramezanius/crypex/exchange/hitbtc"
)

func (suite *hitbtcSuite) TestSubscribeReports() {
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

func (suite *hitbtcSuite) TestSubscribeCandles() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	params := hitbtc.CandlesParams{
		Period: hitbtc.Period1Minute,
		Symbol: hitbtc.BTC + hitbtc.USD,
	}

	suite.NoError(
		suite.exchange.SubscribeCandles(params, func(response interface{}) {
			wg.Done()
		}),
	)

	wg.Wait()
	suite.NoError(suite.exchange.UnsubscribeCandles(params))
}

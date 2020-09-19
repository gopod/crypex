package hitbtc_test

import (
	"sync"

	"github.com/gopod/crypex/exchange/hitbtc"
)

func (suite *hitbtcSuite) TestSubscribeReports() {
	needed, used := 1, 0
	wg := sync.WaitGroup{}

	suite.NoError(
		suite.exchange.SubscribeReports(),
	)

	wg.Add(needed)

	go func() {
		for report := range suite.exchange.Feeds.Reports {
			if used == needed {
				break
			}
			used++

			suite.NotEmpty(report)
			suite.Equal(hitbtc.BTC+hitbtc.USD, report.Symbol)

			wg.Done()
		}
	}()

	suite.TestOrders()

	wg.Wait()
	suite.NoError(suite.exchange.UnsubscribeReports())
}

func (suite *hitbtcSuite) TestSubscribeCandles() {
	needed, used := 2, 0
	wg := sync.WaitGroup{}

	suite.Run("Error", func() {
		suite.NoError(suite.exchange.SubscribeCandles(
			hitbtc.CandlesParams{
				Symbol: "XXX",
				Period: hitbtc.Period1Minute,
			}),
		)
	})

	params := hitbtc.CandlesParams{
		Snapshot: true,
		Period:   hitbtc.Period1Minute,
		Symbol:   hitbtc.BTC + hitbtc.USD,
	}

	wg.Add(needed)

	go func() {
		for candles := range suite.exchange.Feeds.Candles {
			if used == needed {
				break
			}
			used++

			suite.T().Log(candles)
			suite.NotEmpty(candles)
			suite.NotEmpty(candles.Candles)
			suite.Equal(hitbtc.BTC+hitbtc.USD, candles.Symbol)

			wg.Done()
		}
	}()

	suite.NoError(suite.exchange.SubscribeCandles(params))
	wg.Wait()
	suite.NoError(suite.exchange.UnsubscribeCandles(params))
}

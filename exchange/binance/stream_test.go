package binance_test

import (
	"strings"
	"sync"

	"github.com/gopod/crypex/exchange/binance"
)

func (suite *binanceSuite) TestSubscribeReports() {
	needed, used := 2, 0
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
			suite.Equal(binance.BNB+binance.USD, strings.ToLower(report.Symbol))

			wg.Done()
		}
	}()

	suite.TestOrders()

	wg.Wait()
	suite.NoError(suite.exchange.UnsubscribeReports())
}

func (suite *binanceSuite) TestSubscribeCandles() {
	needed, used := 2, 0
	wg := sync.WaitGroup{}

	suite.Run("Error", func() {
		suite.NoError(suite.exchange.SubscribeCandles(
			binance.CandlesParams{
				Symbol: "XXX",
				Period: binance.Period1Minute,
			}),
		)
	})

	params := binance.CandlesParams{
		Snapshot: true,
		Period:   binance.Period1Minute,
		Symbol:   binance.BNB + binance.USD,
	}

	wg.Add(needed)

	go func() {
		for candles := range suite.exchange.Feeds.Candles {
			if used == needed {
				break
			}
			used++

			suite.NotEmpty(candles)
			suite.NotEmpty(candles.Candles)
			suite.Equal(binance.BNB+binance.USD, strings.ToLower(candles.Symbol))

			wg.Done()
		}
	}()

	suite.NoError(suite.exchange.SubscribeCandles(params))
	wg.Wait()
	suite.NoError(suite.exchange.UnsubscribeCandles(params))
}

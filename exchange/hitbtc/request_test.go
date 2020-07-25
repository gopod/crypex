package hitbtc_test

import (
	"github.com/ramezanius/crypex/exchange/hitbtc"
)

func (suite *hitbtcSuite) TestGetSymbols() {
	symbols, err := suite.exchange.GetSymbols()

	suite.NoError(err)
	suite.NotEmpty(symbols)
}

func (suite *hitbtcSuite) TestGetCandles() {
	candles, err := suite.exchange.GetCandles(hitbtc.CandlesParams{
		Symbol: hitbtc.ETH + hitbtc.USD,
		Period: hitbtc.Period1Hour,
	})

	suite.NoError(err)
	suite.NotEmpty(candles)
}

func (suite *hitbtcSuite) TestGetBalances() {
	balances, err := suite.exchange.GetBalances()

	suite.NoError(err)
	suite.NotEmpty(balances)
}

func (suite *hitbtcSuite) TestOrders() {
	newRequest := hitbtc.NewOrder{
		Price:    2000,
		Quantity: 0.00002,

		Side:   hitbtc.Buy,
		Type:   hitbtc.Limit,
		Symbol: hitbtc.BTC + hitbtc.USD,
	}

	suite.Run("NewOrder", func() {
		suite := suite
		order, err := suite.exchange.NewOrder(newRequest)

		suite.NoError(err)
		suite.NotEmpty(order)

		suite.Equal(order.Side, newRequest.Side)
		suite.Equal(order.Price, newRequest.Price)
		suite.Equal(order.Quantity, newRequest.Quantity)

		newRequest.OrderID = order.OrderID
	})
	suite.Run("CancelOrder", func() {
		suite := suite
		order, err := suite.exchange.CancelOrder(newRequest.OrderID)

		suite.NoError(err)
		suite.NotEmpty(order)
	})
	suite.Run("FailNewOrder", func() {
		suite := suite
		newRequest.Symbol = ""
		_, err := suite.exchange.NewOrder(newRequest)

		suite.Error(err)
		suite.NotEmpty(err.Error())
	})
	suite.Run("FailCancelOrder", func() {
		suite := suite
		_, err := suite.exchange.CancelOrder("")

		suite.Error(err)
		suite.NotEmpty(err.Error())
	})
}

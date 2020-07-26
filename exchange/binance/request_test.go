package binance_test

import (
	"github.com/ramezanius/crypex/exchange/binance"
)

func (suite *binanceSuite) TestGetSymbols() {
	symbols, err := suite.exchange.GetSymbols()

	suite.NoError(err)
	suite.NotEmpty(symbols)
}

func (suite *binanceSuite) TestGetCandles() {
	candles, err := suite.exchange.GetCandles(binance.CandlesParams{
		Symbol: binance.ETH + binance.BTC,
		Period: binance.Period1Hour,
		Limit:  1000,
	})

	suite.NoError(err)
	suite.NotEmpty(candles)
}

func (suite *binanceSuite) TestGetBalances() {
	balances, err := suite.exchange.GetBalances()

	suite.NoError(err)
	suite.NotEmpty(balances)
}

func (suite *binanceSuite) TestOrders() {
	newRequest := binance.NewOrder{
		Price:    10,
		Quantity: 1,

		Side:        binance.Buy,
		Type:        binance.Limit,
		TimeInForce: binance.GoodTillCancel,
		Symbol:      binance.BNB + binance.USD,
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
		order, err := suite.exchange.CancelOrder(newRequest.OrderID, newRequest.Symbol)

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
		_, err := suite.exchange.CancelOrder("", "")

		suite.Error(err)
		suite.NotEmpty(err.Error())
	})
}

package binance_test

import "github.com/ramezanius/crypex/exchange/binance"

func (suite *binanceSuite) TestGetSymbols() {
	symbols, err := suite.Exchange.GetSymbols()

	suite.NoError(err)
	suite.NotEmpty(symbols)
}

func (suite *binanceSuite) TestGetBalances() {
	balances, err := suite.Exchange.GetBalances()

	suite.NoError(err)
	suite.NotEmpty(balances)
}

func (suite *binanceSuite) TestOrders() {
	newRequest := binance.NewOrder{
		Price:    10,
		Quantity: 1,

		Side:        binance.Buy,
		Type:        binance.Limit,
		Symbol:      binance.BNB + binance.USD,
		TimeInForce: binance.GoodTillCancel,
	}

	suite.Run("NewOrder", func() {
		order, err := suite.Exchange.NewOrder(newRequest)

		suite.NoError(err)
		suite.NotEmpty(order)

		suite.Equal(order.Side, newRequest.Side)
		suite.Equal(order.Price, newRequest.Price)
		suite.Equal(order.Quantity, newRequest.Quantity)

		newRequest.OrderID = order.OrderID
	})
	suite.Run("CancelOrder", func() {
		order, err := suite.Exchange.CancelOrder(
			newRequest.OrderID, newRequest.Symbol)

		suite.NoError(err)
		suite.NotEmpty(order)
	})
}

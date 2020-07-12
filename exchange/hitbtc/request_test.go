package hitbtc_test

import "github.com/ramezanius/crypex/exchange/hitbtc"

func (suite *hitbtcSuite) TestGetSymbol() {
	symbol, err := suite.Exchange.GetSymbol(hitbtc.Demo)

	suite.NoError(err)
	suite.NotEmpty(symbol)
}

func (suite *hitbtcSuite) TestGetSymbols() {
	symbols, err := suite.Exchange.GetSymbols()

	suite.NoError(err)
	suite.NotEmpty(symbols)
}

func (suite *hitbtcSuite) TestGetBalances() {
	balances, err := suite.Exchange.GetBalances()

	suite.NoError(err)
	suite.NotEmpty(balances)
}

func (suite *hitbtcSuite) TestOrders() {
	newRequest := hitbtc.NewOrder{
		Price:    2000,
		Quantity: 0.00002,

		Side:   hitbtc.Buy,
		Type:   hitbtc.Limit,
		Symbol: hitbtc.Demo,
	}
	replaceRequest := hitbtc.ReplaceOrder{
		Price:    1000,
		Quantity: 0.00001,
	}

	suite.Run("NewOrder", func() {
		order, err := suite.Exchange.NewOrder(newRequest)

		suite.NoError(err)
		suite.NotEmpty(order)

		suite.Equal(order.Side, newRequest.Side)
		suite.Equal(order.Price, newRequest.Price)
		suite.Equal(order.Quantity, newRequest.Quantity)

		replaceRequest.OrderID = order.OrderID
	})
	suite.Run("ReplaceOrder", func() {
		order, err := suite.Exchange.ReplaceOrder(replaceRequest)

		suite.NoError(err)
		suite.NotEmpty(order)

		suite.Equal(order.Price, replaceRequest.Price)
		suite.Equal(order.Quantity, replaceRequest.Quantity)

		replaceRequest.RequestOrderID = order.OrderID
	})
	suite.Run("CancelOrder", func() {
		order, err := suite.Exchange.CancelOrder(replaceRequest.RequestOrderID)

		suite.NoError(err)
		suite.NotEmpty(order)
	})
}

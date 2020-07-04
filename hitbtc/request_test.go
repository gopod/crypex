package hitbtc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc"
	"github.com/ramezanius/crypex/hitbtc/tests"
)

func TestRequests(t *testing.T) {
	instance := tests.SetupHitBTC(t)

	defer func() {
		assert.NoError(t, instance.Shutdown())
	}()

	t.Run("GetSymbol", func(t *testing.T) {
		_, err := instance.GetSymbol("BTC" + "USD")
		assert.NoError(t, err)
	})
	t.Run("GetSymbols", func(t *testing.T) {
		_, err := instance.GetSymbols()
		assert.NoError(t, err)
	})
	t.Run("GetBalances", func(t *testing.T) {
		_, err := instance.GetBalances()
		assert.NoError(t, err)
	})
	t.Run("OrderRequests", func(t *testing.T) {
		newRequest := &hitbtc.NewOrder{
			Price:    2000,
			Quantity: 0.00002,

			Side:   hitbtc.Buy,
			Type:   hitbtc.Limit,
			Symbol: tests.Demo,
		}

		t.Run("NewOrder", func(t *testing.T) {
			order, err := instance.NewOrder(newRequest)

			assert.NoError(t, err)
			assert.Equal(t, order.Side, newRequest.Side)
			assert.Equal(t, order.Price, newRequest.Price)
			assert.Equal(t, order.Quantity, newRequest.Quantity)
		})

		replaceRequest := &hitbtc.ReplaceOrder{
			Price:    1000,
			Quantity: 0.00001,
			OrderID:  newRequest.OrderID,
		}

		t.Run("ReplaceOrder", func(t *testing.T) {
			order, err := instance.ReplaceOrder(replaceRequest)

			assert.NoError(t, err)
			assert.Equal(t, order.Price, replaceRequest.Price)
			assert.Equal(t, order.Quantity, replaceRequest.Quantity)
		})
		t.Run("CancelOrder", func(t *testing.T) {
			_, err := instance.CancelOrder(replaceRequest.RequestOrderID)

			assert.NoError(t, err)
		})
	})
}

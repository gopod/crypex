package hitbtc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const DemoSymbol = "BTC" + "USD"

func TestSubscriptions(t *testing.T) {
	var hitbtc = SetupHitBTC(t)

	t.Run("SubscribeReports", func(t *testing.T) {
		_, snapshot, err := hitbtc.SubscribeReports()
		assert.NoError(t, err)

		t.Run("ReportsSnapshot", func(t *testing.T) {
			_ = <-snapshot
		})
	})

	t.Run("SubscribeCandles", func(t *testing.T) {
		_, snapshot, err := hitbtc.SubscribeCandles(DemoSymbol, Period1Minute, 100)
		assert.NoError(t, err)

		t.Run("CandlesSnapshot", func(t *testing.T) {
			_ = <-snapshot
		})
	})

	t.Run("SubscribeOrderbook", func(t *testing.T) {
		_, snapshot, err := hitbtc.SubscribeOrderbook(DemoSymbol)
		assert.NoError(t, err)

		t.Run("OrderbookSnapshot", func(t *testing.T) {
			_ = <-snapshot
		})
	})
}

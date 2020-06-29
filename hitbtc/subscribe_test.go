package hitbtc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc"
	"github.com/ramezanius/crypex/hitbtc/tests"
)

func TestSubscriptions(t *testing.T) {
	instance := tests.SetupHitBTC(t)

	t.Run("SubscribeReports", func(t *testing.T) {
		_, snapshot, err := instance.SubscribeReports()
		assert.NoError(t, err)

		t.Run("ReportsSnapshot", func(t *testing.T) {
			<-snapshot
		})
	})
	t.Run("SubscribeCandles", func(t *testing.T) {
		_, snapshot, err := instance.SubscribeCandles(tests.Demo, hitbtc.Period1Minute, 100)
		assert.NoError(t, err)

		t.Run("CandlesSnapshot", func(t *testing.T) {
			<-snapshot
		})
	})
	t.Run("UnsubscribeCandles", func(t *testing.T) {
		err := instance.UnsubscribeCandles(tests.Demo)
		assert.NoError(t, err)
	})
	t.Run("SubscribeOrderbook", func(t *testing.T) {
		_, snapshot, err := instance.SubscribeOrderbook(tests.Demo)
		assert.NoError(t, err)

		t.Run("OrderbookSnapshot", func(t *testing.T) {
			<-snapshot
		})
	})
	t.Run("UnsubscribeOrderbook", func(t *testing.T) {
		err := instance.UnsubscribeOrderbook(tests.Demo)
		assert.NoError(t, err)
	})
}

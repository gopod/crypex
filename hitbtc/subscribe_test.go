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
		_, snapshot, err := instance.SubscribeCandles(hitbtc.DemoSymbol, hitbtc.Period1Minute, 100)
		assert.NoError(t, err)

		t.Run("CandlesSnapshot", func(t *testing.T) {
			<-snapshot
		})
	})
	t.Run("SubscribeOrderbook", func(t *testing.T) {
		_, snapshot, err := instance.SubscribeOrderbook(hitbtc.DemoSymbol)
		assert.NoError(t, err)

		t.Run("OrderbookSnapshot", func(t *testing.T) {
			<-snapshot
		})
	})
}

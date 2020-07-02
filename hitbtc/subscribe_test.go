package hitbtc_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc"
	"github.com/ramezanius/crypex/hitbtc/tests"
)

func TestSubscriptions(t *testing.T) {
	instance := tests.SetupHitBTC(t)

	t.Run("SubscribeReports", func(t *testing.T) {
		reports, err := instance.SubscribeReports()
		assert.NoError(t, err)

		t.Run("Reports", func(t *testing.T) {
			for stay, timeout := true, time.After(hitbtc.Timeout); stay; {
				select {
				case <-timeout:
					stay = false
				case _, ok := <-reports:
					assert.True(t, ok)
				}
			}
		})
	})
	t.Run("SubscribeCandles", func(t *testing.T) {
		candles, err :=
			instance.SubscribeCandles(tests.Demo, hitbtc.Period1Minute, 100)
		assert.NoError(t, err)

		t.Run("Candles", func(t *testing.T) {
			for stay, timeout := true, time.After(hitbtc.Timeout); stay; {
				select {
				case <-timeout:
					stay = false
				case _, ok := <-candles:
					assert.True(t, ok)
				}
			}
		})
	})
	t.Run("UnsubscribeCandles", func(t *testing.T) {
		err := instance.UnsubscribeCandles(tests.Demo)
		assert.NoError(t, err)
	})
	t.Run("SubscribeOrderbook", func(t *testing.T) {
		orderbook, err := instance.SubscribeOrderbook(tests.Demo)
		assert.NoError(t, err)

		t.Run("Orderbook", func(t *testing.T) {
			for stay, timeout := true, time.After(hitbtc.Timeout); stay; {
				select {
				case <-timeout:
					stay = false
				case _, ok := <-orderbook:
					assert.True(t, ok)
				}
			}
		})
	})
	t.Run("UnsubscribeOrderbook", func(t *testing.T) {
		err := instance.UnsubscribeOrderbook(tests.Demo)
		assert.NoError(t, err)
	})
}

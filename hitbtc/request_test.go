package hitbtc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequests(t *testing.T) {
	var (
		err error

		hitbtc = SetupHitBTC(t)
	)

	t.Run("GetSymbol", func(t *testing.T) {
		_, err = hitbtc.GetSymbol("BTC" + "USD")
		assert.NoError(t, err)
	})

	t.Run("GetSymbols", func(t *testing.T) {
		_, err = hitbtc.GetSymbols()
		assert.NoError(t, err)

	})

	t.Run("GetBalances", func(t *testing.T) {
		_, err = hitbtc.GetBalances()
		assert.NoError(t, err)
	})
}

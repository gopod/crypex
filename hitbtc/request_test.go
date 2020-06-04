package hitbtc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc/tests"
)

func TestRequests(t *testing.T) {
	var (
		err error

		instance = tests.SetupHitBTC(t)
	)

	t.Run("GetSymbol", func(t *testing.T) {
		_, err = instance.GetSymbol("BTC" + "USD")
		assert.NoError(t, err)
	})
	t.Run("GetSymbols", func(t *testing.T) {
		_, err = instance.GetSymbols()
		assert.NoError(t, err)
	})
	t.Run("GetBalances", func(t *testing.T) {
		_, err = instance.GetBalances()
		assert.NoError(t, err)
	})
}

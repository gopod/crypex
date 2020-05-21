package hitbtc

import (
	"testing"
)

func TestRequests(t *testing.T) {
	hitbtc, err := SetupHitBTC()
	if err != nil {
		t.Error(err)
		os.Exit(0)
	}

	t.Run("GetSymbol", func(t *testing.T) {
		_, err = hitbtc.GetSymbol("BTC" + "USD")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("GetSymbols", func(t *testing.T) {
		_, err = hitbtc.GetSymbols()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("GetBalances", func(t *testing.T) {
		_, err = hitbtc.GetBalances()
		if err != nil {
			t.Error(err)
		}
	})
}

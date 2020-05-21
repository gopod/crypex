package hitbtc

import (
	"os"
	"testing"
)

var hitbtc *HitBTC

func TestRequests(t *testing.T) {
	var err error

	hitbtc, err = New()
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
}

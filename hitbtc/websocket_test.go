package hitbtc

import (
	"os"
	"testing"
)

func TestHitBTC(t *testing.T) {
	var err error

	hitbtc, err = New()
	if err != nil {
		t.Error(err)
		os.Exit(0)
	}

	t.Run("Authenticate", func(t *testing.T) {
		err = hitbtc.Authenticate("pubKey", "secKey")
		if err != nil {
			return
		}
	})
}

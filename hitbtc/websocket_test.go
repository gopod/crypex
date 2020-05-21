package hitbtc

import (
	"os"
	"testing"
)

func SetupHitBTC() (instance *HitBTC, err error) {
	var (
		publicKey = os.Getenv("HITBTC_PUBLIC_KEY")
		secretKey = os.Getenv("HITBTC_SECRET_KEY")
	)

	instance, err = New()
	if err != nil {
		return nil, err
	}

	if publicKey != "" && secretKey != "" {
		err = instance.Authenticate(publicKey, secretKey)
		if err != nil {
			return nil, err
		}
	}

	return
}

func TestHitBTC(t *testing.T) {
	var err error

	hitbtc, err := SetupHitBTC()
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

package hitbtc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetupHitBTC(t *testing.T) (instance *HitBTC) {
	var (
		err error

		publicKey = os.Getenv("HITBTC_PUBLIC_KEY")
		secretKey = os.Getenv("HITBTC_SECRET_KEY")
	)

	instance, err = New()
	assert.NoError(t, err)

	if publicKey != "" && secretKey != "" {
		err = instance.Authenticate(publicKey, secretKey)
		assert.NoError(t, err)
	}

	return
}

func TestHitBTC(t *testing.T) {
	var err error

	hitbtc := SetupHitBTC(t)

	t.Run("Authenticate", func(t *testing.T) {
		err = hitbtc.Authenticate("pubKey", "secKey")
		assert.Error(t, err)
	})
}

package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc"
)

func SetupHitBTC(t *testing.T) (instance *hitbtc.HitBTC) {
	var (
		err error

		publicKey = os.Getenv("HITBTC_PUBLIC_KEY")
		secretKey = os.Getenv("HITBTC_SECRET_KEY")
	)

	instance, err = hitbtc.New()
	assert.NoError(t, err)

	if publicKey != "" && secretKey != "" {
		err = instance.Authenticate(publicKey, secretKey)
		assert.NoError(t, err)
	}

	return
}

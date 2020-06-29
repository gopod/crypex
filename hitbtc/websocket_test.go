package hitbtc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc/tests"
)

func TestHitBTC(t *testing.T) {
	var err error

	instance := tests.SetupHitBTC(t)

	t.Run("Authenticate", func(t *testing.T) {
		err = instance.Authenticate()
		assert.NoError(t, err)
	})
}

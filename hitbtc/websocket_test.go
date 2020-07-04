package hitbtc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ramezanius/crypex/hitbtc/tests"
)

func TestHitBTC(t *testing.T) {
	instance := tests.SetupHitBTC(t)

	defer func() {
		assert.NoError(t, instance.Shutdown())
	}()

	t.Run("Authenticate", func(t *testing.T) {
		assert.NoError(t, instance.Authenticate())
	})
}

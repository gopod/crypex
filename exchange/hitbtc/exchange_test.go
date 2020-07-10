package hitbtc_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ramezanius/crypex/exchange/hitbtc"
)

type hitbtcSuite struct {
	suite.Suite
	Exchange *hitbtc.HitBTC
}

func (suite *hitbtcSuite) SetupSuite() {
	exchange, err := hitbtc.New(
		os.Getenv("HITBTC_PUBLIC_KEY"),
		os.Getenv("HITBTC_SECRET_KEY"),
	)
	suite.NoError(err)

	suite.Exchange = exchange
}

func (suite *hitbtcSuite) TearDownSuite() {
	assert.NoError(suite.T(), suite.Exchange.Shutdown())
}

func TestHitBTC(t *testing.T) {
	suite.Run(t, new(hitbtcSuite))
}

package hitbtc_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/gopod/crypex/exchange/hitbtc"
)

type hitbtcSuite struct {
	suite.Suite
	exchange *hitbtc.HitBTC
}

func (suite *hitbtcSuite) SetupSuite() {
	suite.exchange = hitbtc.New()

	suite.exchange.OnErr = func(err error) {
		log.Println(err)
	}

	suite.exchange.PublicKey = os.Getenv("HITBTC_PUBLIC_KEY")
	suite.exchange.SecretKey = os.Getenv("HITBTC_SECRET_KEY")
}

func (suite *hitbtcSuite) TearDownSuite() {
	assert.NoError(suite.T(), suite.exchange.Shutdown())
}

func TestHitBTC(t *testing.T) {
	suite.Run(t, new(hitbtcSuite))
}

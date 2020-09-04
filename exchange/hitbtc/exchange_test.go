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

	suite.exchange.SetStreams(func(res interface{}) {
		switch response := res.(type) {
		case *hitbtc.CandlesStream:
			if len(response.Candles) == 1 {
				log.Println("[HitBTC][Candles] Update Received")
			} else {
				log.Println("[HitBTC][Candles] Snapshot Received")
			}
		case *hitbtc.APIError:
			log.Printf("[HitBTC][Candles] Error: %v\n", response)
		}
	}, func(res interface{}) {
		switch response := res.(type) {
		case *hitbtc.ReportsStream:
			log.Println("[HitBTC][Reports] Report Received")
		case *hitbtc.APIError:
			log.Printf("[HitBTC][Reports] Error: %v\n", response)
		}
	})

	suite.exchange.PublicKey = os.Getenv("HITBTC_PUBLIC_KEY")
	suite.exchange.SecretKey = os.Getenv("HITBTC_SECRET_KEY")
}

func (suite *hitbtcSuite) TearDownSuite() {
	assert.NoError(suite.T(), suite.exchange.Shutdown())
}

func TestHitBTC(t *testing.T) {
	suite.Run(t, new(hitbtcSuite))
}

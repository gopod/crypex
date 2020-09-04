package binance_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/gopod/crypex/exchange/binance"
)

type binanceSuite struct {
	suite.Suite

	exchange *binance.Binance
}

func (suite *binanceSuite) SetupSuite() {
	suite.exchange = binance.New()

	suite.exchange.SetStreams(func(res interface{}) {
		switch response := res.(type) {
		case *binance.CandlesStream:
			if len(response.Candles) == 1 {
				log.Println("[Binance][Candles] Update Received")
			} else {
				log.Println("[Binance][Candles] Snapshot Received")
			}
		case *binance.APIError:
			log.Printf("[Binance][Candles] Error: %v\n", response)
		}
	}, func(res interface{}) {
		switch response := res.(type) {
		case *binance.ReportsStream:
			log.Println("[Binance][Reports] Report Received")
		case *binance.APIError:
			log.Printf("[Binance][Reports] Error: %v\n", response)
		}
	})

	suite.exchange.PublicKey = os.Getenv("BINANCE_PUBLIC_KEY")
	suite.exchange.SecretKey = os.Getenv("BINANCE_SECRET_KEY")
}

func (suite *binanceSuite) TearDownSuite() {
	assert.NoError(suite.T(), suite.exchange.Shutdown())
}

func TestBinance(t *testing.T) {
	suite.Run(t, new(binanceSuite))
}

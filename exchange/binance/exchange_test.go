package binance_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ramezanius/crypex/exchange/binance"
)

type binanceSuite struct {
	suite.Suite

	exchange *binance.Binance
}

func (suite *binanceSuite) SetupSuite() {
	suite.exchange = binance.New()

	suite.exchange.PublicKey = os.Getenv("BINANCE_PUBLIC_KEY")
	suite.exchange.SecretKey = os.Getenv("BINANCE_SECRET_KEY")
}

func (suite *binanceSuite) TearDownSuite() {
	assert.NoError(suite.T(), suite.exchange.Shutdown())
}

func TestBinance(t *testing.T) {
	suite.Run(t, new(binanceSuite))
}

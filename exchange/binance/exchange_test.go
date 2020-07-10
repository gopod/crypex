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
	Exchange *binance.Binance
}

func (suite *binanceSuite) SetupSuite() {
	exchange, err := binance.New(
		os.Getenv("BINANCE_PUBLIC_KEY"),
		os.Getenv("BINANCE_SECRET_KEY"),
	)
	suite.NoError(err)

	suite.Exchange = exchange
}

func (suite *binanceSuite) TearDownSuite() {
	assert.NoError(suite.T(), suite.Exchange.Shutdown())
}

func TestBinance(t *testing.T) {
	suite.Run(t, new(binanceSuite))
}

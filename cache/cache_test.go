package cache_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ramezanius/crypex/cache"
	"github.com/ramezanius/crypex/exchange/binance"
	"github.com/ramezanius/crypex/exchange/hitbtc"
)

const testLimit = 100

type cacheSuite struct {
	suite.Suite
	cache *cache.Cache
}

func (suite *cacheSuite) SetupSuite() {
	var err error
	suite.cache, err = cache.NewCache(testLimit)
	suite.NoError(err)
}

func TestCache(t *testing.T) {
	suite.Run(t, new(cacheSuite))
}

func (suite *cacheSuite) TestGetPrefix() {
	tests := []struct {
		want   string
		opts   []string
		prefix cache.Prefix
	}{
		{"HITBTC:BTCUSD:CANDLES", []string{"HITBTC", "BTCUSD"}, cache.Candles},
		{"BINANCE:ETHUSD:SYMBOL", []string{"BINANCE", "ETHUSD"}, cache.Symbol},
		{"COINBASE:TRXUSD:PRICE", []string{"COINBASE", "TRXUSD"}, cache.Price},
	}

	for _, test := range tests {
		suite.Run(test.want, func() {
			suite.Equal(test.want, cache.GetPrefix(test.prefix, test.opts...))
		})
	}
}

func (suite *cacheSuite) TestCacheStore() {
	tests := []struct {
		key   string
		entry []byte
	}{
		{cache.GetPrefix(cache.Candles, "BTCUSD"), []byte(`{"S": "BTCUSD", "O": 1, "C": 2}`)},
		{cache.GetPrefix(cache.Candles, "ETHUSD"), []byte(`{"S": "ETHUSD", "O": 0.4, "C": 0.45}`)},
	}

	for _, test := range tests {
		suite.Run(test.key, func() {
			suite.NoError(suite.cache.Store(test.key, test.entry))
		})
	}
}

func (suite *cacheSuite) TestCacheLoad() {
	tests := []struct {
		key   string
		entry interface{}
	}{
		{cache.GetPrefix(cache.Candles, "BTCUSD"), hitbtc.Candles{}},
		{cache.GetPrefix(cache.Candles, "ETHUSD"), binance.Candles{}},
	}

	suite.TestCacheStore()

	for _, test := range tests {
		suite.Run(test.key, func() {
			suite.NoError(suite.cache.Load(test.key, test.entry))
		})
	}
}

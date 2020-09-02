package cache

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/allegro/bigcache"
)

type Cache struct {
	cache *bigcache.BigCache
}

func NewCache(maxLimit int) (*Cache, error) {
	cache, err := bigcache.NewBigCache(
		bigcache.Config{
			Shards:             1024,
			LifeWindow:         time.Hour * 24,
			MaxEntrySize:       maxLimit * 400,
			MaxEntriesInWindow: 1000 * 10 * 60,
		})
	if err != nil {
		return nil, err
	}

	return &Cache{cache}, nil
}

type Prefix int

const (
	Candles Prefix = iota
	Symbol
	Price
)

func GetPrefix(prefix Prefix, opts ...string) string {
	switch prefix {
	case Candles:
		opts = append(opts, "CANDLES")
	case Symbol:
		opts = append(opts, "SYMBOL")
	case Price:
		opts = append(opts, "PRICE")
	}

	return strings.Join(opts, ":")
}

func (c *Cache) Store(key string, entry []byte) (err error) {
	err = c.cache.Set(key, entry)

	return
}

func (c *Cache) Load(key string, v interface{}) (err error) {
	entry, err := c.cache.Get(key)
	if err != nil {
		return
	}

	err = json.Unmarshal(entry, &v)

	return
}

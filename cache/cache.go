package cache

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/allegro/bigcache"
)

// Cache cache storage for data loader
type Cache struct {
	cache *bigcache.BigCache
}

// NewCache returns a new cache object.
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

// Prefix type
type Prefix int

const (
	// Data prefixes
	Candles Prefix = iota
	Symbol
	Price
)

// GetPrefix returns the prefix with arguments.
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

// Store stores a entry to the specific key on cache object.
func (c *Cache) Store(key string, entry []byte) (err error) {
	err = c.cache.Set(key, entry)

	return
}

// Load loads entry from cache an bind it to the v.
func (c *Cache) Load(key string, v interface{}) (err error) {
	entry, err := c.cache.Get(key)
	if err != nil {
		return
	}

	err = json.Unmarshal(entry, &v)

	return
}

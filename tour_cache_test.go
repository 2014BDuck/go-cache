// @Author: 2014BDuck
// @Date: 2020/10/17

package cache_test

import (
	"github.com/2014BDuck/cache"
	"github.com/2014BDuck/cache/lru"
	"github.com/matryer/is"
	"log"
	"sync"
	"testing"
)

func TestTourCacheGet(t *testing.T) {
	db := map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
		"key4": "val4",
	}

	getter := cache.GetFunc(func(key string) interface{} {
		log.Println("[From DB] find key", key)

		if val, ok := db[key]; ok {
			return val
		}
		return nil
	})

	tourCache := cache.NewTourCache(getter, lru.New(0, nil))

	i := is.New(t)

	var wg sync.WaitGroup

	for k, v := range db {
		wg.Add(1)
		go func(k, v string) {
			defer wg.Done()
			i.Equal(tourCache.Get(k), v)
			i.Equal(tourCache.Get(k), v)
		}(k, v)
	}
	wg.Wait()

	i.Equal(tourCache.Get("unknown"), nil)
	i.Equal(tourCache.Get("unknown"), nil)

	i.Equal(tourCache.Stat().NGet, 10)
	i.Equal(tourCache.Stat().NHit, 4)
}

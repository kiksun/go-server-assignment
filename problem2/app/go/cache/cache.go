package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	cacheInstance *cache.Cache
)

func init() {
	cacheInstance = cache.New(5*time.Minute, 10*time.Minute)
}

func SetCacheValue(k string, v interface{}) {
	cacheInstance.Set(k, v, cache.DefaultExpiration)
}

func GetCacheValue(k string) (interface{}, bool) {
	return cacheInstance.Get(k)
}

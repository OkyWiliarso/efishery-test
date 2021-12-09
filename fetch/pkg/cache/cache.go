package cachelib

import (
	"time"

	"github.com/ReneKroon/ttlcache/v2"
)

var cache ttlcache.SimpleCache = ttlcache.NewCache()

func SetCache(key string, value float64) {
	cache.SetTTL(time.Duration(5 * time.Second))

	cache.Set(key, value)
}

func GetCache(key string) interface{} {
	val, _ := cache.Get(key)

	return val
}

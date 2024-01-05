package utils

import (
	"time"
	gc "github.com/patrickmn/go-cache"
)

type CacheEntry struct {
	Body        []byte
	StatusCode  int
	ContentType string
}

var cache *gc.Cache

func GetCache() *gc.Cache {
	if cache == nil {
		cache = gc.New(30 * time.Minute, time.Minute)
	}
	return cache
}

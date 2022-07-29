package service

import (
	"github.com/ReneKroon/ttlcache"
	"itaMus90/rateLimit/src/entity"
	"sync"
)

var InputData entity.InputArguments
var (
	mu        sync.Mutex
	protectMe int
)


func IsLimit(hashUrl string, cache *ttlcache.Cache, ch chan int) {
	cacheValue, ok := cache.Get(hashUrl)

	if !ok {
		cache.Set(hashUrl, 1)
		ch <- 0
		return
	}

	protectMe := cacheValue.(int)

	if protectMe >= InputData.Threshold {
		ch <- 1
		return
	}
	
	mu.Lock()
	protectMe++
	cache.Set(hashUrl, protectMe)
	ch <- 0
	mu.Unlock()
	return
}

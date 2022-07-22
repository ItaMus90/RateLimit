package service

import (
	"github.com/ReneKroon/ttlcache"
	"itaMus90/rateLimit/src/entity"
)

var InputData entity.InputArguments

func IsLimit(hashUrl string, cache *ttlcache.Cache) bool {
	cacheValue, ok := cache.Get(hashUrl)

	if !ok {
		cache.Set(hashUrl, 1)
		return false
	}

	counter := cacheValue.(int)

	if counter >= InputData.Threshold {
		return true
	}
	counter++
	cache.Set(hashUrl, counter)
	return false
}

//func Init() {
//	skipTtlExtensionOnHit(true)
//}
//
//func skipTtlExtensionOnHit(value bool) {
//	cache.SkipTtlExtensionOnHit(value)
//}

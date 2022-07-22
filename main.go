package main

import (
	"fmt"
	"github.com/ReneKroon/ttlcache"
	"itaMus90/rateLimit/src/router"
	"itaMus90/rateLimit/src/service"
	"time"
)

func main() {
	isInputOk, ttl := service.CheckInputArguments()

	if !isInputOk {
		fmt.Println("Input arguments must to be number")
		return
	}
	ttlFixed := time.Duration(ttl * 1000000000)
	var cache = ttlcache.NewCache()
	cache.SkipTtlExtensionOnHit(true)
	cache.SetTTL(ttlFixed)

	router := router.SetupRouter(cache)
	router.Run("localhost:8080")

}

package router

import (
	"github.com/ReneKroon/ttlcache"
	"github.com/gin-gonic/gin"
	"itaMus90/rateLimit/src/controller"
)

func SetupRouter(cache *ttlcache.Cache) *gin.Engine {
	router := gin.Default()
	router.POST("/report", func(context *gin.Context) {
		controller.IsLimitURL(context, cache)
	})

	return router
}

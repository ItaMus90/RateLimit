package controller

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/ReneKroon/ttlcache"
	"github.com/gin-gonic/gin"
	"itaMus90/rateLimit/src/service"
	"net/http"
)

type URLRequestBody struct {
	Url string `binding:"required"`
}

func IsLimitURL(context *gin.Context, cache *ttlcache.Cache) {
	var requestBody URLRequestBody

	if err := context.BindJSON(&requestBody); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing url param"})
		return
	}

	if len(requestBody.Url) == 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "URL can't be empty"})
		return
	}
	hashUrl := md5.New()
	hashUrl.Write([]byte(requestBody.Url))
	isBlock := service.IsLimit(hex.EncodeToString(hashUrl.Sum(nil)), cache)

	context.IndentedJSON(http.StatusOK, gin.H{"block": isBlock})
}

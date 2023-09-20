package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/v1/data/shorten", CreateShortUrl)
}

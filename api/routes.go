package api

import (
	"url-shortener-service/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/v1/data/shorten", handlers.CreateShortUrl)
	router.GET(":shortUrl", handlers.RedirectToLongUrl)

	router.POST("/api/v1/auth/register", handlers.RegisterHandler)

	//router.POST("/api/v1/auth/login", h)
}

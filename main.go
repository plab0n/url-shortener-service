package main

import (
	"url-shortener-service/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.SetupRoutes(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

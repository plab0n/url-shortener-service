package main

import (
	"fmt"
	"github.com/spf13/viper"
	"url-shortener-service/api"

	"github.com/gin-gonic/gin"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	r := gin.Default()
	api.SetupRoutes(r)
	// Listen and Server in 0.0.0.0:8080
	port := viper.GetString("httpPort")
	r.Run(port)
}

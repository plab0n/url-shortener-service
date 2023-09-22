package main

import (
	"fmt"
	"os"
	"url-shortener-service/api"
	"url-shortener-service/db"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func main() {
	viper.AddConfigPath("./configs")
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	fmt.Println("Environment : ", env)
	configName := env + "-config"
	viper.SetConfigName(configName)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	db.NewGormDb()
	r := gin.Default()
	api.SetupRoutes(r)
	// Listen and Server in 0.0.0.0:8080
	port := viper.GetString("httpPort")
	r.Run(port)
}

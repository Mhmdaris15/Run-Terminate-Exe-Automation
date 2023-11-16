package main

import (
	"exe-handler/dependencies"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Authorization", "X-Requested-With", "Accept", "Accept-Encoding", "Accept-Language", "Connection", "Host", "Origin", "Referer", "User-Agent", "Username"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	dependencies.SetupRoutes(router)

	err := router.Run(":" + dependencies.EnvPort())

	if err != nil {
		log.Fatal(err)
	}
}

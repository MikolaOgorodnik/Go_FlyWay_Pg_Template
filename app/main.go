package main

import (
	"Check_V8/config"
	"Check_V8/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	entryHost := config.AppCfg.EntryHost
	appPort := config.AppCfg.AppPort
	address := entryHost + ":" + appPort

	router := gin.Default()

	router.GET("/health", service.HealthCheck)
	router.GET("/", service.Index)

	err := router.Run(address)
	if err != nil {
		return
	}
}

package main

import (
	"Pg_FW_Template/config"
	"Pg_FW_Template/service"
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
	router.GET("/api/template/:id", service.GetTemplate)

	err := router.Run(address)
	if err != nil {
		return
	}
}

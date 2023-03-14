package app

import (
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/config"
	"github.com/gin-gonic/gin"
)

func RunApp() error {
	dependencies := GetDependencies()
	engine := gin.New()
	engine.GET("/items/:itemID", dependencies.ItemsHdl().GetItem)
	engine.POST("/items/search", dependencies.ItemsHdl().SearchItems)
	return engine.Run(config.Port())
}

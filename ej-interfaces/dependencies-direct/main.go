package main

import (
	"github.com/emikohmann/ucc-arqsoft-1/ej-interfaces/dependencies-direct/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.GET("/:key", controllers.Get)
	engine.POST("/:key/:value", controllers.Save)
	engine.Run(":8080")
}

package app

import (
	"github.com/emikohmann/dependencies-injection/controllers"
	"github.com/emikohmann/dependencies-injection/services"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	var serviceImplementation services.ServiceImplementation

	var controllerImplementation controllers.ControllerImplementation
	controllerImplementation.Service = serviceImplementation

	engine := gin.New()
	engine.GET("/:key", controllerImplementation.Get)
	engine.POST("/:key/:value", controllerImplementation.Save)
	engine.Run(":8080")
}

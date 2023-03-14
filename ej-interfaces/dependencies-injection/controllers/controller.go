package controllers

import (
	"github.com/emikohmann/dependencies-injection/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	Get(c *gin.Context)
	Save(c *gin.Context)
}

type ControllerImplementation struct {
	Service services.Service
}

func (controller ControllerImplementation) Get(c *gin.Context) {
	key := c.Param("key")
	value := controller.Service.Get(key)
	c.String(http.StatusOK, value)
}

func (controller ControllerImplementation) Save(c *gin.Context) {
	key := c.Param("key")
	value := c.Param("value")
	controller.Service.Save(key, value)
	c.Status(http.StatusOK)
}

func (controller ControllerImplementation) Util() {
	controller.Service.Get("util")
}

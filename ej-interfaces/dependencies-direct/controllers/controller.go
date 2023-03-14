package controllers

import (
	"github.com/emikohmann/ucc-arqsoft-1/ej-interfaces/dependencies-direct/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	key := c.Param("key")
	value := service.Get(key)
	c.String(http.StatusOK, value)
}

func Save(c *gin.Context) {
	key := c.Param("key")
	value := c.Param("value")
	service.Save(key, value)
	c.Status(http.StatusOK)
}

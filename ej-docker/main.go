package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.New()
	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	engine.Run(":8081")
}

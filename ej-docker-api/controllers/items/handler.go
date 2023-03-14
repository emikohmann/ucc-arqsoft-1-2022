package items

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetItem(c *gin.Context)
	SearchItems(c *gin.Context)
}

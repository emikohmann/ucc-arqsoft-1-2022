package items

import (
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/domain/search"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/services/items"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/utils/errs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const (
	errInvalidParam   = "invalid param"
	errInvalidPayload = "invalid payload"
)

type HandlerImpl struct {
	itemSvc items.Service
}

func NewHandlerImpl(itemsSvc items.Service) HandlerImpl {
	return HandlerImpl{
		itemSvc: itemsSvc,
	}
}

func (handler HandlerImpl) GetItem(c *gin.Context) {
	id := c.Param("itemID")

	itemID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiErr := errs.NewStatusBadRequestError(errInvalidParam)
		c.JSON(apiErr.StatusCode(), apiErr)
		return
	}

	item, apiErr := handler.itemSvc.GetItem(c.Request.Context(), itemID)
	if apiErr != nil {
		log.Println("error getting item", apiErr)
		c.JSON(apiErr.StatusCode(), apiErr)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (handler HandlerImpl) SearchItems(c *gin.Context) {
	var req search.Query

	if err := c.ShouldBindJSON(&req); err != nil {
		apiErr := errs.NewStatusBadRequestError(errInvalidPayload)
		c.JSON(apiErr.StatusCode(), apiErr)
		return
	}

	result, apiErr := handler.itemSvc.SearchItems(c.Request.Context(), req)
	if apiErr != nil {
		log.Println("error searching items", apiErr)
		c.JSON(apiErr.StatusCode(), apiErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

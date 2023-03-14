package items

import (
	"github.com/gin-gonic/gin"
)

type HandlerMock struct{}

func NewHandlerMock() HandlerMock {
	return HandlerMock{}
}

func (HandlerMock) GetItem(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (HandlerMock) SearchItems(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

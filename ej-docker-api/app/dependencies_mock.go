package app

import (
	itemsHandler "github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/controllers/items"
	itemsService "github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/services/items"
)

type DependenciesMock struct {
	itemsSvc itemsService.Service
	itemsHdl itemsHandler.Handler
}

func NewDependenciesMock() DependenciesMock {
	itemsSvc := itemsService.NewServiceMock()
	itemsHdl := itemsHandler.NewHandlerMock()

	return DependenciesMock{
		itemsSvc: itemsSvc,
		itemsHdl: itemsHdl,
	}
}

func (dep DependenciesMock) ItemsSvc() itemsService.Service {
	return dep.itemsSvc
}

func (dep DependenciesMock) ItemsHdl() itemsHandler.Handler {
	return dep.itemsHdl
}
